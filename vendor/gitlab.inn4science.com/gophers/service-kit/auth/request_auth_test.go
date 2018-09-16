package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.inn4science.com/gophers/service-kit/crypto"
)

type TestModel struct {
	TestData string `json:"testData"`
}

func TestVerifyRequestSignature(t *testing.T) {
	privateKey, publicKey := crypto.GenKeyPair()

	testRequest, err := NewSignedGetRequest(privateKey,
		"https://localhost:8080/test/42?que=ctulhu", "dummy")
	ok, _ := VerifyRequestSignature(testRequest, publicKey)
	assert.Equal(t, true, ok)
	fmt.Println(err)

	testRequest, err = NewSignedPostRequest(privateKey,
		"https://localhost:8080/test/42?que=ctulhu",
		[]byte("testing signature"),
		"plain/text", "dummy")
	fmt.Println(err)

	ok, err = VerifyRequestSignature(testRequest, publicKey)
	fmt.Println(err)

	assert.Equal(t, true, ok)

}
func TestNewSignedDataRequest(t *testing.T) {
	privateKey, publicKey := crypto.GenKeyPair()
	println("[", privateKey, ",", publicKey, "]")
	x := &TestModel{
		TestData: "Vegas",
	}
	testRequest, err := NewSignedDataRequest("POST", privateKey,
		"https://localhost:8080/test/42?que=ctulhu", x, "dummy")

	assert.NoError(t, err)
	ok, err := VerifyRequestSignature(testRequest, publicKey)
	assert.NoError(t, err)
	assert.Equal(t, true, ok)

}
