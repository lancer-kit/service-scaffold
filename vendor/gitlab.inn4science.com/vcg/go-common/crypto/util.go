package crypto

import "encoding/base32"

func base32Decode(str string) ([]byte, error) {
	return base32.StdEncoding.
		WithPadding(base32.StdPadding).
		DecodeString(str)
}

func base32Encode(data []byte) string {
	return base32.StdEncoding.
		WithPadding(base32.StdPadding).
		EncodeToString(data)
}
