package auth

import (
	"context"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"strconv"
	"strings"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/pkg/errors"
)

const (
	authTokenPath       = "/oauth2/v2.0/token"
	clientAssertionType = "urn:ietf:params:oauth:client-assertion-type:jwt-bearer"
	grantType           = "client_credentials"
)

type customClaims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

func (c customClaims) Valid() error {
	return nil
}

type Client struct {
	httpClient      HTTPClient
	url             string
	certFingerprint string
	privateKey      string
	appId           string
	clientId        string
	tenantId        string
}

func NewClient(httpClient HTTPClient, url, certFingerprint, privateKey, appId, clientId, tenantId string) *Client {
	return &Client{
		httpClient:      httpClient,
		url:             url,
		certFingerprint: certFingerprint,
		privateKey:      privateKey,
		appId:           appId,
		clientId:        clientId,
		tenantId:        tenantId,
	}
}

func (a Client) Token(ctx context.Context) (string, time.Duration, error) {
	form, err := a.buildParams()
	if err != nil {
		return "", 0, err
	}

	r, err := a.httpClient.NewRequest().
		SetHeaders(a.buildHeaders()).
		SetContext(ctx).
		SetFormData(form).
		Post(a.url + "/" + a.tenantId + authTokenPath)
	if err != nil {
		return "", 0, err
	}

	var content map[string]interface{}
	err = r.UnmarshalJson(&content)
	if err != nil {
		return "", 0, err
	}
	if errorDesc, valid := content["error_description"].(string); valid {
		return "", 0, errors.New(errorDesc)
	}
	if _, valid := content["access_token"].(string); !valid {
		return "", 0, errors.New("access token not set, error not understood")
	}

	return content["access_token"].(string), time.Duration(content["expires_in"].(float64)) * time.Second, nil

}

func (a Client) x5t(fingerprint string) string {
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

func (a Client) randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (a Client) buildParams() (map[string]string, error) {
	signed, err := a.buildClientAssertion()
	if err != nil {
		return nil, err
	}

	return map[string]string{
			"tenant":                a.tenantId,
			"client_id":             a.clientId,
			"scope":                 a.appId + "/.default",
			"client_assertion_type": clientAssertionType,
			"client_assertion":      signed,
			"grant_type":            grantType,
		},
		nil
}

func (a Client) buildHeaders() map[string]string {
	return map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}
}

func (a Client) buildClientAssertion() (string, error) {
	block, _ := pem.Decode([]byte(a.privateKey))
	if block == nil {
		return "", errors.New("nothing to decode private key is invalid")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	randID, err := a.randomHex(12)
	if err != nil {
		return "", err
	}

	now := time.Now().Unix()
	claims := customClaims{
		Subject:   a.clientId,
		Issuer:    a.clientId,
		Id:        randID,
		NotBefore: now,
		Audience:  a.url + "/" + a.tenantId + authTokenPath,
		ExpiresAt: now + 3600,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["x5t"] = a.x5t(a.certFingerprint)

	signed, err := token.SignedString(key)

	if err != nil {
		return "", err
	}
	return signed, nil
}
