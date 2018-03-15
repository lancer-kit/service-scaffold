package auth

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/inn4sc/vcg-go-common/crypto"
	"github.com/pkg/errors"
)

const (
	HeaderHash        = "X-Auth-Hash"
	HeaderSignature   = "X-Auth-Signature"
	HeaderSigner      = "X-Auth-Signer"
	HeaderService     = "X-Auth-Service"
	HeaderContentType = "Content-Type"
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
	sign, err := crypto.SignMessage(privateKey, msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	req.Header.Set(HeaderSignature, sign)
	req.Header.Set(HeaderService, service)
	return req, nil
}

// NewSignedPostRequest creates a new POST request, hashes the body,
// sings the request details using the `privateKey` and adds the auth headers.
func NewSignedPostRequest(privateKey, path string, body []byte, mimeType, service string) (*http.Request, error) {
	bodyHash, err := crypto.HashData(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash body")
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new http request")
	}

	fullPath := req.URL.Path + req.URL.RawQuery
	msg := msgSchema(service, req.Method, fullPath, bodyHash, mimeType)
	sign, err := crypto.SignMessage(privateKey, msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	req.Header.Set("Content-Type", mimeType)
	req.Header.Set(HeaderHash, bodyHash)
	req.Header.Set(HeaderSignature, sign)
	req.Header.Set(HeaderService, service)
	return req, nil
}

// VerifyRequestSignature checks the request auth headers.
func VerifyRequestSignature(r *http.Request, publicKey string) (bool, error) {
	bodyHash := r.Header.Get(HeaderHash)
	mimeType := r.Header.Get(HeaderContentType)
	service := r.Header.Get(HeaderService)
	sign := r.Header.Get(HeaderSignature)

	fullPath := r.URL.Path + r.URL.RawQuery
	msg := msgSchema(service, r.Method, fullPath, bodyHash, mimeType)

	return crypto.VerifySignature(publicKey, msg, sign)
}

func msgSchema(service, method, url, body, mime string) string {
	return fmt.Sprintf("service: %s; method: %s; path: %s; body: %s; content-type: %s",
		service, method, url, body, mime)
}
