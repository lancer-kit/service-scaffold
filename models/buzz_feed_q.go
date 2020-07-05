package models

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/lancer-kit/armory/db"
)

type BuzzFeedQI interface {
	// Insert adds new `BuzzFeed` record to `buzzFeeds` table.
	Insert(buzzFeed *BuzzFeed) error
	// Update updates row with passed `uid`.
	Update(id int64, buzzFeed *BuzzFeed) error
	// UpdateBuzzDescription sets new value of
	// `description` column for row with passed `uid`.
	UpdateBuzzDescription(id int64, description string) error
	// DeleteByID deletes row with passed `id`.
	DeleteByID(id int64) error

	// WithID adds filter by `ID` column.
	WithID(ID int64) BuzzFeedQI
	// WithName adds filter by `Name` column.
	WithName(Name string) BuzzFeedQI
	// WithBuzzType adds filter by `BuzzType` column.
	WithBuzzType(BuzzType ExampleType) BuzzFeedQI
	// WithDescription adds filter by `Description` column.
	WithDescription(Description string) BuzzFeedQI
	// WithDetails adds filter by `Details` column.
	WithDetails(Details Feed) BuzzFeedQI
	// WithCreatedAt adds filter by `CreatedAt` column.
	WithCreatedAt(CreatedAt int64) BuzzFeedQI
	// WithUpdatedAt adds filter by `UpdatedAt` column.
	WithUpdatedAt(UpdatedAt int64) BuzzFeedQI

	// Get returns first row of the result of query execution.
	Get() (*BuzzFeed, error)
	// Select returns all records of the result of query execution.
	Select() ([]BuzzFeed, error)
	// GetByID returns one row with passed `id`.
	GetByID(id int64) (*BuzzFeed, error)

	// Until sets lower time bound.
	Since(timestamp int64) BuzzFeedQI
	// Until sets upper time bound.
	Until(timestamp int64) BuzzFeedQI
	// SetPage applies pagination parameters.
	SetPage(pq *db.PageQuery) BuzzFeedQI
}

const tableBuzzFeeds = "buzzFeeds"

type buzzFeedQ struct {
	parent *Q
	table  db.Table

	Err error
}

func (q *Q) BuzzFeed() BuzzFeedQI {
	return &buzzFeedQ{
		parent: q,
		table: db.Table{
			Name:     tableBuzzFeeds,
			QBuilder: sq.Select("*").From(tableBuzzFeeds),
		},
	}
}

// Insert adds new `BuzzFeed` record to `buzzFeeds` table.
func (q *buzzFeedQ) Insert(buzzFeed *BuzzFeed) error {
	query := sq.Insert(q.table.Name).SetMap(map[string]interface{}{

		"id":          buzzFeed.ID,
		"name":        buzzFeed.Name,
		"buzz_type":   buzzFeed.BuzzType,
		"description": buzzFeed.Description,
		"details":     buzzFeed.Details,
		"created_at":  buzzFeed.CreatedAt,
		"updated_at":  buzzFeed.UpdatedAt,
	})

	var err error
	_, err = q.parent.Insert(query)
	return err
}

// Update updates row with passed `uid`.
//fixme: check that this is the correct update
func (q *buzzFeedQ) Update(id int64, buzzFeed *BuzzFeed) error {
	query := sq.Update(q.table.Name).SetMap(map[string]interface{}{

		"id":          buzzFeed.ID,
		"name":        buzzFeed.Name,
		"buzz_type":   buzzFeed.BuzzType,
		"description": buzzFeed.Description,
		"details":     buzzFeed.Details,
		"created_at":  buzzFeed.CreatedAt,
		"updated_at":  buzzFeed.UpdatedAt,
	}).Where("id = ?", id)
	return q.parent.Exec(query)
}

func (q *buzzFeedQ) UpdateBuzzDescription(id int64, description string) error {
	query := sq.Update(q.table.Name).SetMap(map[string]interface{}{
		"description": description,
		"updated_at":  time.Now().Unix(),
	}).Where("id = ?", id)
	return q.parent.Exec(query)
}

// WithID adds filter by `ID` column.
func (q *buzzFeedQ) WithID(id int64) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("id = ?", id)
	return q
}

// WithName adds filter by `Name` column.
func (q *buzzFeedQ) WithName(name string) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("name = ?", name)
	return q
}

// WithBuzzType adds filter by `BuzzType` column.
func (q *buzzFeedQ) WithBuzzType(buzzType ExampleType) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("buzz_type = ?", buzzType)
	return q
}

// WithDescription adds filter by `Description` column.
func (q *buzzFeedQ) WithDescription(description string) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("description = ?", description)
	return q
}

// WithDetails adds filter by `Details` column.
func (q *buzzFeedQ) WithDetails(details Feed) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("details = ?", details)
	return q
}

// WithCreatedAt adds filter by `CreatedAt` column.
func (q *buzzFeedQ) WithCreatedAt(createdAt int64) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("created_at = ?", createdAt)
	return q
}

// WithUpdatedAt adds filter by `UpdatedAt` column.
func (q *buzzFeedQ) WithUpdatedAt(updatedAt int64) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("updated_at = ?", updatedAt)
	return q
}

// Until sets lower time bound.
func (q *buzzFeedQ) Since(timestamp int64) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("created_at >= ?", timestamp)
	return q
}

// Until sets upper time bound.
func (q *buzzFeedQ) Until(timestamp int64) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("created_at <= ?", timestamp)
	return q
}

// SetPage applies pagination parameters.
func (q *buzzFeedQ) SetPage(pq *db.PageQuery) BuzzFeedQI {
	q.table.SetPage(pq)
	return q
}

// Select returns all records of the result of query execution.
func (q *buzzFeedQ) Select() ([]BuzzFeed, error) {
	res := make([]BuzzFeed, 0, 1)
	q.table.ApplyPage("id")

	err := q.parent.Select(q.table.QBuilder, &res)
	if err == sql.ErrNoRows {
		return res, nil
	}

	return res, err
}

// Get returns first row of the result of query execution.
func (q *buzzFeedQ) Get() (*BuzzFeed, error) {
	res := new(BuzzFeed)
	q.table.ApplyPage("id")

	err := q.parent.Get(q.table.QBuilder, res)
	if err == sql.ErrNoRows {
		return res, nil
	}

	return res, err
}

// GetByID returns one row with passed `id`.
// fixme: check that this is the correct getter
func (q *buzzFeedQ) GetByID(id int64) (*BuzzFeed, error) {
	res := new(BuzzFeed)
	err := q.parent.Get(q.table.QBuilder.Where("id = ?", id), res)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return res, err
}

// DeleteByID deletes row with passed `id`.
// fixme: check that this is the correct getter
func (q *buzzFeedQ) DeleteByID(id int64) error {
	return q.parent.Exec(sq.Delete(q.table.Name).Where("id = ?", id))
}
