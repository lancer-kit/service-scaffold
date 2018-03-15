package vcgtools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyRequestSignature(t *testing.T) {
	privateKey, publicKey := GenKeyPair()

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
