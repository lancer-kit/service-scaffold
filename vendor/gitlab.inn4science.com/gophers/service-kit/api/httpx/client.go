package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/gophers/service-kit/crypto"
)

const (
	HeaderBodyHash  = "X-Auth-BHash"
	HeaderSignature = "X-Auth-Signature"
	HeaderSigner    = "X-Auth-Signer"
	HeaderService   = "X-Auth-Service"
	HeaderJWTParsed = "jwt"
)

type XClient struct {
	http.Client

	auth    bool
	kp      crypto.KP
	service string
}

func NewXClient() *XClient {
	return &XClient{Client: http.Client{Timeout: defaultTimeout}}
}

// Auth returns current state of authentication flag.
func (client *XClient) Auth() bool {
	return client.auth
}

func (client *XClient) OffAuth() Client {
	client.auth = false
	return client
}

// OnAuth enables request authentication.
func (client *XClient) OnAuth() Client {
	client.auth = true
	return client
}

// Service returns auth service name.
func (client *XClient) Service() string {
	return client.service
}

// PublicKey returns client public key.
func (client *XClient) PublicKey() crypto.Key {
	return client.kp.Public
}

// SetAuth sets the auth credentials.
func (client *XClient) SetAuth(service string, kp crypto.KP) Client {
	if client == nil {
		client = &XClient{Client: http.Client{Timeout: defaultTimeout}}
	}

	client.kp = kp
	client.auth = true
	client.service = service

	return client
}

// PostJSON, sets passed `headers` and `body` and executes RequestJSON with POST method.
func (client *XClient) PostJSON(url string, body interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPost, url, body, headers)
}

// PutJSON, sets passed `headers` and `body` and executes RequestJSON with PUT method.
func (client *XClient) PutJSON(url string, body interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPut, url, body, headers)
}

// PatchJSON, sets passed `headers` and `body` and executes RequestJSON with PATCH method.
func (client *XClient) PatchJSON(url string, body interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPatch, url, body, headers)
}

// GetJSON, sets passed `headers` and executes RequestJSON with GET method.
func (client *XClient) GetJSON(url string, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodGet, url, nil, headers)
}

// DeleteJSON, sets passed `headers` and executes RequestJSON with DELETE method.
func (client *XClient) DeleteJSON(url string, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodDelete, url, nil, headers)
}

// RequestJSON creates and executes new request with JSON content type.
func (client *XClient) RequestJSON(method string, url string, data interface{}, headers Headers) (*http.Response, error) {
	var body io.Reader = nil
	var err error
	var rawData []byte

	if data != nil {
		rawData, err = json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "unable to marshal body")
		}
		body = bytes.NewBuffer(rawData)
	}

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	if client.auth {
		req, err = client.SignRequest(req, rawData)
		if err != nil {
			return nil, errors.Wrap(err, "unable to sign request")
		}
	}
	return client.Do(req)
}

// ParseJSONBody decodes `json` body from the `http.Request`.
// > `dest` must be a pointer value.
func (client *XClient) ParseJSONBody(r *http.Request, dest interface{}) error {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal request body")
	}
	return nil
}

// ParseJSONResult decodes `json` body from the `http.Response` body into `dest`
// > `dest` must be a pointer value.
func (client *XClient) ParseJSONResult(httpResp *http.Response, dest interface{}) error {
	defer httpResp.Body.Close()
	err := json.NewDecoder(httpResp.Body).Decode(dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal response body")
	}
	return nil
}

// SignRequest takes body hash, some headers and full URL path,
// sings this request details using the `client.privateKey` and adds the auth headers.
func (client *XClient) SignRequest(req *http.Request, body []byte) (*http.Request, error) {
	bodyHash, err := crypto.HashData(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash body")
	}

	fullPath := req.URL.Path + req.URL.RawQuery
	msg := messageForSigning(client.service, req.Method, fullPath,
		bodyHash, req.Header.Get(HeaderJWTParsed))

	sign, err := client.kp.Sign([]byte(msg))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	req.Header.Set(HeaderBodyHash, bodyHash)
	req.Header.Set(HeaderSignature, sign)
	req.Header.Set(HeaderService, client.service)
	req.Header.Set(HeaderSigner, client.kp.Public.String())
	return req, nil
}

// VerifyBody checks the request body match with it hash.
func (client *XClient) VerifyBody(r *http.Request, body []byte) (bool, error) {
	bodyHash, err := crypto.HashData(body)
	if err != nil {
		return false, errors.Wrap(err, "failed to hash body")
	}
	return bodyHash == r.Header.Get(HeaderBodyHash), nil
}

// VerifyRequest checks the request auth headers.
func (client *XClient) VerifyRequest(r *http.Request, publicKey string) (bool, error) {
	if publicKey != r.Header.Get(HeaderSigner) {
		return false, errors.New("signer mismatch with passed public key")
	}

	bodyHash := r.Header.Get(HeaderBodyHash)
	service := r.Header.Get(HeaderService)
	sign := r.Header.Get(HeaderSignature)
	authHeader := r.Header.Get(HeaderJWTParsed)

	fullPath := r.URL.Path + r.URL.RawQuery
	msg := messageForSigning(service, r.Method, fullPath, bodyHash, authHeader)

	return crypto.VerifySignature(publicKey, msg, sign)
}

// messageForSigning concatenates passed request data in a fixed format.
func messageForSigning(service, method, url, body, authHeaders string) string {
	return fmt.Sprintf("service:%s;method:%s;path:%s;authHeaders:%s;body:%s;",
		service, method, url, authHeaders, body)
}
