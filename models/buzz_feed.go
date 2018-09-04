package models

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"gitlab.inn4science.com/gophers/service-kit/db"
)

type BuzzFeed struct {
	ID          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	BuzzType    ExampleType `db:"buzz_type" json:"buzzType"`
	Description string      `db:"description"json:"description"`
	Details     Feed        `db:"details"json:"details"`
	CreatedAt   int64       `db:"created_at" json:"createdAt"`
	UpdatedAt   int64       `db:"updated_at" json:"updatedAt"`
}

type BuzzFeedQ struct {
	*Q
	db.Table
}

func NewBuzzFeedQ(q *Q) *BuzzFeedQ {
	return &BuzzFeedQ{
		Q: q,
		Table: db.Table{
			Name:      "buzz_feeds",
			QBuilder:  sq.Select("*").From("buzz_feeds"),
			IQBuilder: sq.Insert("buzz_feeds"),
			UQBuilder: sq.Update("buzz_feeds"),
			DQBuilder: sq.Delete("buzz_feeds"),
		},
	}
}

func (q *BuzzFeedQ) Insert(bf BuzzFeed) error {
	_, err := q.DBConn.Insert(
		q.IQBuilder.SetMap(map[string]interface{}{
			"name":        bf.Name,
			"buzz_type":   bf.BuzzType,
			"description": bf.Description,
			"details":     bf.Details,
			"created_at":  time.Now().UTC().Unix(),
			"updated_at":  time.Now().UTC().Unix(),
		}),
	)
	return err
}

func (q *BuzzFeedQ) ByID(id int64) (*BuzzFeed, error) {
	res := new(BuzzFeed)
	err := q.WithID(id).DBConn.Get(q.QBuilder, res)
	if err == sql.ErrNoRows {
		return res, nil
	}
	return res, err
}

func (q *BuzzFeedQ) WithName(name string) *BuzzFeedQ {
	q.QBuilder = q.QBuilder.Where("name = ?", name)
	return q
}

func (q *BuzzFeedQ) WithID(id int64) *BuzzFeedQ {
	q.QBuilder = q.QBuilder.Where("id = ?", id)
	return q
}

func (q *BuzzFeedQ) SetPage(query *db.PageQuery) *BuzzFeedQ {
	q.Table.SetPage(query)
	return q
}

func (q *BuzzFeedQ) Select() ([]BuzzFeed, error) {
	q.ApplyPage("id")
	res := make([]BuzzFeed, 0)
	err := q.DBConn.Select(q.QBuilder, &res)
	if err == sql.ErrNoRows {
		return res, nil
	}
	return res, err
}

func (q *BuzzFeedQ) UpdateDetails(id int64, details Feed) error {
	err := q.DBConn.Exec(
		q.UQBuilder.Set("details", details).Where("id = ?", id),
	)
	return err
}

func (q *BuzzFeedQ) UpdateBuzzDescription(id int64, description string) error {
	err := q.DBConn.Exec(
		q.UQBuilder.Set("description", description).Where("id = ?", id),
	)
	return err
}

func (q *BuzzFeedQ) DeleteBuzzByID(id int64) error {
	err := q.DBConn.Exec(
		q.DQBuilder.Where("id = ?", id),
	)
	return err
}
