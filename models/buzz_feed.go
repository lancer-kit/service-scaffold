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
			Name:      "buzz_feed",
			QBuilder:  sq.Select("*").From("buzz_feed"),
			IQBuilder: sq.Insert("buzz_feed"),
			UQBuilder: sq.Update("buzz_feed"),
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

func (q *BuzzFeedQ) ByID() (*BuzzFeed, error) {
	res := new(BuzzFeed)
	err := q.DBConn.Get(q.QBuilder, &res)
	if err == sql.ErrNoRows {
		return res, nil
	}
	return res, err
}

func (q *BuzzFeedQ) WithName(name string) *BuzzFeedQ {
	q.QBuilder = q.QBuilder.Where("name = ?", name)
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
