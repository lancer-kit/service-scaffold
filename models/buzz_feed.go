package models

import (
	"database/sql/driver"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"

	"github.com/pkg/errors"
)

type BuzzFeed struct {
	ID          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description"json:"description"`
	Details     Feed   `db:"details"json:"details"`
}

type BuzzFeedQ struct {
	*Q
	sqBuilder sq.SelectBuilder
}

// NewPriceQ returns the new instance of the `PriceQ`.
func NewBuzzFeedQ(q *Q) *BuzzFeedQ {
	return &BuzzFeedQ{
		Q:         q,
		sqBuilder: sq.Select("*").From("buzz_feed").OrderBy("id"),
	}
}

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
