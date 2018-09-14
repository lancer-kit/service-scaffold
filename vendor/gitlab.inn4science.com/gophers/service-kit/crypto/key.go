package crypto

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

// Key is a type for
type Key []byte

// ToBytes cast `key` to `[]byte `.
func (key Key) ToBytes() []byte {
	return []byte(key)
}

// ToPrivate cast `key` to `ed25519.PrivateKey`.
func (key Key) ToPrivate() ed25519.PrivateKey {
	return ed25519.PrivateKey(key)
}

// ToPublicKey cast `key` to `ed25519.PublicKey`.
func (key Key) ToPublicKey() ed25519.PublicKey {
	return ed25519.PublicKey(key)
}

// String is a realization of the `Stringer` interface.
func (key Key) String() string {
	return Base32Encode(key)
}

// UnmarshalJSON is a realization of the `Unmarshaller` interface.
func (key *Key) UnmarshalJSON(data []byte) error {
	var rawStr string
	err := json.Unmarshal(data, &rawStr)
	if err != nil {
		return errors.Wrap(err, "can not to unmarshal Key into string")
	}
	rawKey, err := Base32Decode(rawStr)
	if err != nil {
		return errors.Wrap(err, "can not to decode Key")
	}
	*key = Key(rawKey)
	return nil
}

// MarshalJSON is a realization of the `Marshaller` interface.
func (key Key) MarshalJSON() ([]byte, error) {
	return []byte(key.String()), nil
}

// Value is a realization of the `Valuer` interface.
func (key Key) Value() (driver.Value, error) {
	j, err := json.Marshal(key)
	return j, err
}

// Scan is realization of the `Scanner` interface.
func (key *Key) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed")
	}

	var i Key
	err := json.Unmarshal(source, &i)
	if err != nil {
		return errors.Wrap(err, "Key: can't unmarshal column data")
	}

	*key = i
	return nil
}
