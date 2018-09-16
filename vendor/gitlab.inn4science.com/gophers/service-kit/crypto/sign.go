package crypto

import (
	"encoding/json"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

// SignData serializes the `data` and signs it using the `privateKey`.
func SignData(privKey string, data interface{}) (string, error) {
	pk, err := Base32Decode(privKey)
	if err != nil {
		return "", errors.Wrap(err, "invalid private key")
	}
	message, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "can not to marshal data")
	}

	sig := ed25519.Sign(pk, message)
	return Base32Encode(sig), nil
}

// SignMessage signs the `message` using the `privateKey`.
func SignMessage(privKey, message string) (string, error) {
	pk, err := Base32Decode(privKey)
	if err != nil {
		return "", err
	}

	msg := []byte(message)
	sig := ed25519.Sign(pk, msg)
	return Base32Encode(sig), nil
}

// VerifySignature checks is valid `signature` of `message`.
func VerifySignature(pubKey, message, signature string) (bool, error) {
	var rawSignature []byte

	rawPubKey, err := Base32Decode(pubKey)
	if err != nil {
		return false, err
	}
	rawSignature, err = Base32Decode(signature)
	if err != nil {
		return false, err
	}

	return ed25519.Verify(rawPubKey, []byte(message), rawSignature), nil
}
