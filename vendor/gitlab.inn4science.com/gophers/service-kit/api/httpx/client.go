package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"gitlab.inn4science.com/gophers/service-kit/log"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/gophers/service-kit/crypto"
)

const (
	HeaderBodyHash    = "X-Auth-BHash"
	HeaderSignature   = "X-Auth-Signature"
	HeaderSigner      = "X-Auth-Signer"
	HeaderService     = "X-Auth-Service"
	HeaderJWTParsed   = "jwt"
	HeaderPassHeaders = "X-Custom-Headers"
)

type XClient struct {
	http.Client

	auth    bool
	kp      crypto.KP
	service string

	defaultHeaders Headers
	cookies        []*http.Cookie

	logger log.Entry
}

func NewXClient() *XClient {
	return &XClient{
		Client:         http.Client{Timeout: defaultTimeout},
		defaultHeaders: map[string]string{},
		cookies:        []*http.Cookie{},
	}
}

// SetLogger - Set logger to enable log requests
func (client *XClient) SetLogger(l log.Entry) {
	client.logger = l
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
		client = NewXClient()
	}

	client.kp = kp
	client.auth = true
	client.service = service

	return client
}

// DefaultCookies returns a client's default cookies.
func (client *XClient) DefaultCookies() []*http.Cookie {
	if client == nil {
		client = NewXClient()
	}
	return client.cookies
}

// SetCookies sets a default cookies to the client.
func (client *XClient) SetDefaultCookies(cookies []*http.Cookie) Client {
	if client == nil {
		client = NewXClient()
	}
	client.cookies = append(client.cookies, cookies...)
	return client
}

// RemoveDefaultCookies removes a default client's cookies.
func (client *XClient) RemoveDefaultCookies() Client {
	if client == nil {
		client = NewXClient()
	}
	client.cookies = nil
	return client
}

// WithCookies append cookies to the client and return new instance.
func (client *XClient) WithCookies(cookies []*http.Cookie) Client {
	if client == nil {
		client = NewXClient()

	}
	newClient := *client
	newClient.cookies = append(client.cookies, cookies...)
	return &newClient
}

// DefaultHeaders returns a client's default headers.
func (client *XClient) DefaultHeaders() Headers {
	if client == nil {
		client = NewXClient()
	}
	return client.defaultHeaders
}

// SetDefaultHeaders sets a default headers to the client.
func (client *XClient) SetDefaultHeaders(headers Headers) Client {
	if client == nil {
		client = NewXClient()
	}
	if client.defaultHeaders == nil {

	}
	for key := range headers {
		client.defaultHeaders[key] = headers[key]
	}
	return client
}

// SetHeader sets new default header to the client.
func (client *XClient) SetHeader(key, val string) Client {
	client.defaultHeaders[key] = val
	return client
}

// RemoveDefaultHeaders removes a default client's headers.
func (client *XClient) RemoveDefaultHeaders() Client {
	if client == nil {
		client = NewXClient()
	}
	client.defaultHeaders = map[string]string{}
	return client
}

// WithHeaders append headers to the client and return new instance.
func (client *XClient) WithHeaders(headers Headers) Client {
	if client == nil {
		client = NewXClient()

	}
	newClient := *client
	for key := range headers {
		newClient.defaultHeaders[key] = headers[key]
	}
	return &newClient
}

// PostJSON, sets passed `headers` and `body` and executes RequestJSON with POST method.
func (client *XClient) PostJSON(url string, bodyStruct interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPost, url, bodyStruct, headers)
}

// PutJSON, sets passed `headers` and `body` and executes RequestJSON with PUT method.
func (client *XClient) PutJSON(url string, bodyStruct interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPut, url, bodyStruct, headers)
}

// PatchJSON, sets passed `headers` and `body` and executes RequestJSON with PATCH method.
func (client *XClient) PatchJSON(url string, bodyStruct interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPatch, url, bodyStruct, headers)
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
func (client *XClient) RequestJSON(method string, url string, bodyStruct interface{}, headers Headers) (*http.Response, error) {
	var body io.Reader = nil
	var err error
	var rawData []byte

	switch bodyStruct.(type) {
	case []byte:
		rawData = bodyStruct.([]byte)
		body = bytes.NewBuffer(rawData)
	default:
		if bodyStruct != nil {
			rawData, err = json.Marshal(bodyStruct)
			if err != nil {
				return nil, errors.Wrap(err, "unable to marshal body")
			}
			body = bytes.NewBuffer(rawData)
		}
	}

	if client.logger != nil {
		client.logger.
			WithField("method", method).
			WithField("url", method).
			WithField("headers", headers).
			WithField("body", string(rawData)).Debug()
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if len(client.cookies) != 0 {
		req = addCookies(req, client.cookies)
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		client.defaultHeaders[key] = value
	}
	for key, value := range client.defaultHeaders {
		req.Header.Set(key, value)
	}

	if client.auth {
		req, err = client.SignRequest(req, rawData, headers)
		if err != nil {
			return nil, errors.Wrap(err, "unable to sign request")
		}
	}
	return client.Do(req)
}

// ParseJSONBody decodes `json` body from the `http.Request`.
// !> `dest` must be a pointer value.
func (client *XClient) ParseJSONBody(r *http.Request, dest interface{}) error {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal request body")
	}
	if client.logger != nil {
		client.logger.WithField("url", r.URL.String()).
			WithField("method", r.Method).
			WithField("auth", r.Header.Get("Authorization")).
			WithField("body", string(b)).Debug()
	}
	err = json.Unmarshal(b, dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal request body")
	}
	return nil
}

// ParseJSONResult decodes `json` body from the `http.Response` body into `dest`
// > `dest` must be a pointer value.
func (client *XClient) ParseJSONResult(httpResp *http.Response, dest interface{}) error {
	defer httpResp.Body.Close()
	b, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal request body")
	}
	if client.logger != nil {
		client.logger.WithField("url", httpResp.Request.URL.String()).
			WithField("method", httpResp.Request.Method).
			WithField("auth", httpResp.Header.Get("Authorization")).
			WithField("body", string(b)).Debug()

	}
	err = json.Unmarshal(b, dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal response body")
	}
	return nil
}

// SignRequest takes body hash, some headers and full URL path,
// sings this request details using the `client.privateKey` and adds the auth headers.
func (client *XClient) SignRequest(req *http.Request, body []byte, headers map[string]string) (*http.Request, error) {
	bodyHash, err := crypto.HashData(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash body")
	}

	fullPath := req.URL.Path + req.URL.RawQuery
	signHeadrs := headersForSigning(headers)
	msg := messageForSigning(client.service, req.Method, fullPath,
		bodyHash, signHeadrs)

	sign, err := client.kp.Sign([]byte(msg))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	req.Header.Set(HeaderPassHeaders, signHeadrs)
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
	headers := r.Header.Get(HeaderPassHeaders)

	fullPath := r.URL.Path + r.URL.RawQuery
	msg := messageForSigning(service, r.Method, fullPath, bodyHash, headers)

	return crypto.VerifySignature(publicKey, msg, sign)
}

// PostSignedWithHeaders create new POST signed request with headers
func (client *XClient) PostSignedWithHeaders(url string, data interface{}, headers map[string]string) (*http.Response, error) {
	rawData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal body")
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(rawData))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new http request")
	}
	rg, err := client.SignRequest(req, rawData, headers)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}
	for key, value := range headers {
		rg.Header.Set(key, value)
	}

	return client.Do(rg)
}

// PostSignedWithHeaders create new signed GET request with headers
func (client *XClient) GetSignedWithHeaders(url string, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new http request")
	}
	rq, err := client.SignRequest(req, nil, headers)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}
	if headers != nil {
		for key, value := range headers {
			rq.Header.Set(key, value)
		}
	}

	return client.Do(rq)
}
