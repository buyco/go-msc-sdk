package client

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

import (
	jwt "github.com/form3tech-oss/jwt-go"
	http "github.com/imroc/req"
)

const (
	//authTokenPath = "/oauth2/token"
	authTokenPath = "/oauth2/v2.0/token"
	clientAssertionType = "urn:ietf:params:oauth:client-assertion-type:jwt-bearer"
	grantType = "client_credentials"
)

type Auth struct {
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

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
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

func (a *Auth) Token(expiresAt int64) (string, time.Duration, error) {
	block, _ := pem.Decode([]byte(a.PrivateKey))
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", 0, err
	}

	randID, err := randomHex(12)
	if err != nil {
		return "", 0, err
	}

	now := time.Now().Unix()
	claims := CustomClaims{
		Subject:   a.ClientId,
		Issuer:    a.ClientId,
		Id:        randID,
		NotBefore: now,
		Audience:  a.Url + "/" + a.TenantId + authTokenPath,
		ExpiresAt: now + expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["x5t"] = x5t(a.CertFingerprint)

	signed, err := token.SignedString(key)

	if err != nil {
		return "", 0, err
	}
	form := http.Param{
		"tenant":                a.TenantId,
		"client_id":             a.ClientId,
		"scope":                 a.AppId + "/.default",
		"client_assertion_type": clientAssertionType,
		"client_assertion":      signed,
		"grant_type":            grantType,
	}

	headers := http.Header{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	r, err := http.Post(a.Url + "/" + a.TenantId + authTokenPath, headers, form)
	if err != nil {
		return "", 0, nil
	}
	var d map[string]string

	err = r.ToJSON(&d)

	expiresIn, err := time.ParseDuration(d["expires_in"])
	return d["access_token"], expiresIn, nil

}
