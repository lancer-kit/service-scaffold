package crypto

import (
	"encoding/base32"
	"encoding/base64"
	"strings"
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

//No padding decode base 64
func Base64DecodeNP(str string) ([]byte, error) {
	return base64.StdEncoding.
		WithPadding(base64.NoPadding).
		DecodeString(str)
}

//No padding encode base 64
func Base64EncodeNP(data []byte) string {
	return base64.StdEncoding.
		WithPadding(base64.NoPadding).
		EncodeToString(data)
}

func Base32DecodeNP(str string) ([]byte, error) {
	return base32.StdEncoding.
		WithPadding(base32.NoPadding).
		DecodeString(strings.ToUpper(str))
}

func Base32EncodeNP(data []byte) string {
	return strings.ToLower(base32.StdEncoding.
		WithPadding(base32.NoPadding).
		EncodeToString(data))
}
