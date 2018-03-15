package vcgtools

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const (
	AuthHeaderHash        = "X-Auth-Hash"
	AuthHeaderSignature   = "X-Auth-Signature"
	AuthHeaderSigner      = "X-Auth-Signer"
	AuthHeaderService     = "X-Auth-Service"
	AuthHeaderContentType = "Content-Type"
)

// NewSignedGetRequest creates a new GET request, sings the request
// details using the `privateKey` and adds the auth headers.
func NewSignedGetRequest(privateKey, path, service string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new http request")
	}
	fullPath := req.URL.Path + req.URL.RawQuery
	msg := msgSchema(service, req.Method, fullPath, "", "")
	sign, err := SignMessage(privateKey, msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	req.Header.Set(AuthHeaderSignature, sign)
	req.Header.Set(AuthHeaderService, service)
	return req, nil
}

// NewSignedPostRequest creates a new POST request, hashes the body,
// sings the request details using the `privateKey` and adds the auth headers.
func NewSignedPostRequest(privateKey, path string, body []byte, mimeType, service string) (*http.Request, error) {
	bodyHash, err := HashData(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash body")
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new http request")
	}

	fullPath := req.URL.Path + req.URL.RawQuery
	msg := msgSchema(service, req.Method, fullPath, bodyHash, mimeType)
	sign, err := SignMessage(privateKey, msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	req.Header.Set("Content-Type", mimeType)
	req.Header.Set(AuthHeaderHash, bodyHash)
	req.Header.Set(AuthHeaderSignature, sign)
	req.Header.Set(AuthHeaderService, service)
	return req, nil
}

// VerifyRequestSignature checks the request auth headers.
func VerifyRequestSignature(r *http.Request, publicKey string) (bool, error) {
	bodyHash := r.Header.Get(AuthHeaderHash)
	mimeType := r.Header.Get(AuthHeaderContentType)
	service := r.Header.Get(AuthHeaderService)
	sign := r.Header.Get(AuthHeaderSignature)

	fullPath := r.URL.Path + r.URL.RawQuery
	msg := msgSchema(service, r.Method, fullPath, bodyHash, mimeType)

	return VerifySignature(publicKey, msg, sign)
}

func msgSchema(service, method, url, body, mime string) string {
	return fmt.Sprintf("service: %s; method: %s; path: %s; body: %s; content-type: %s",
		service, method, url, body, mime)
}
