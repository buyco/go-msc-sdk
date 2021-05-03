package msc

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"strconv"
	"strings"
	"time"
)

import (
	jwt "github.com/form3tech-oss/jwt-go"
	http "github.com/imroc/req"
)

type AuthClient struct {
	Url             string
	CertFingerprint string
	PrivateKey      string
	AppId           string
	ClientId        string
	TenantId        string
}

func x5t(fingerprint string) string {
	f := strings.Split(fingerprint, ":")
	var z int64
	var e = make([]byte, len(f))
	for i, v := range f {
		z, _ = strconv.ParseInt(v, 16, 32)
		e[i] = byte(z)
	}

	s := base64.URLEncoding.EncodeToString(e)
	return s
}

type CustomClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

func (c CustomClaims) Valid() error {
	return nil
}

func (c *AuthClient) Token() (string, time.Duration, error) {

	block, _ := pem.Decode([]byte(c.PrivateKey))

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", 0, err
	}

	now := time.Now().Unix()
	claims := CustomClaims{
		Subject:   c.ClientId,
		Issuer:    c.ClientId,
		Id:        "myid",
		NotBefore: now,
		Audience:  c.Url + "/" + c.TenantId + "/oauth2/token",
		ExpiresAt: now + 1200,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["x5t"] = x5t(c.CertFingerprint)

	signed, err := token.SignedString(key)

	if err != nil {
		return "", 0, err
	}
	form := http.Param{
		"tenant":                c.TenantId,
		"client_id":             c.ClientId,
		"scope":                 c.AppId + "/.default",
		"client_assertion_type": "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
		"client_assertion":      signed,
		"grant_type":            "client_credentials",
	}

	headers := http.Header{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	r, err := http.Post(c.Url+"/"+c.TenantId+"/oauth2/v2.0/token", headers, form)
	if err != nil {
		return "", 0, nil
	}
	var d map[string]string

	err = r.ToJSON(&d)

	expires_in, err := time.ParseDuration(d["expires_in"])
	return d["access_token"], expires_in, nil

}
