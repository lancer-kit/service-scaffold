package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

type Feed struct {
}

func (f *Feed) UnmarshalJSON(data []byte) error {
	return nil
}

func (f Feed) MarshalJSON() ([]byte, error) {
	return []byte("{}"), nil
}

func (f Feed) Value() (driver.Value, error) {
	j, err := json.Marshal(f)
	return j, err
}

func (f *Feed) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i Feed
	err := json.Unmarshal(source, &i)
	if err != nil {
		return errors.Wrap(err, "Feed: can't unmarshal column data")
	}

	*f = i
	return nil
}
