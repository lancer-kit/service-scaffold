package crypto

import (
	"encoding/base32"
	"encoding/base64"
)

func Base32Decode(str string) ([]byte, error) {
	return base32.StdEncoding.
		WithPadding(base32.StdPadding).
		DecodeString(str)
}

func Base32Encode(data []byte) string {
	return base32.StdEncoding.
		WithPadding(base32.StdPadding).
		EncodeToString(data)
}

func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.
		WithPadding(base64.StdPadding).
		DecodeString(str)
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.
		WithPadding(base64.StdPadding).
		EncodeToString(data)
}
