package repo

import (
	"database/sql"
	"time"

	"lancer-kit/service-scaffold/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/lancer-kit/armory/db"
)

type BuzzFeedQI interface {
	// Insert adds new `BuzzFeed` record to `buzzFeeds` table.
	Insert(buzzFeed *models.BuzzFeed) error
	// Update updates row with passed `uid`.
	Update(id int64, buzzFeed *models.BuzzFeed) error
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
	WithBuzzType(BuzzType models.ExampleType) BuzzFeedQI
	// WithDescription adds filter by `Description` column.
	WithDescription(Description string) BuzzFeedQI
	// WithDetails adds filter by `Details` column.
	WithDetails(Details models.Feed) BuzzFeedQI
	// WithCreatedAt adds filter by `CreatedAt` column.
	WithCreatedAt(CreatedAt int64) BuzzFeedQI
	// WithUpdatedAt adds filter by `UpdatedAt` column.
	WithUpdatedAt(UpdatedAt int64) BuzzFeedQI
	// Since sets lower time bound.
	Since(timestamp int64) BuzzFeedQI
	// Until sets upper time bound.
	Until(timestamp int64) BuzzFeedQI

	// Get returns first row of the result of query execution.
	Get() (*models.BuzzFeed, error)
	// GetByID returns one row with passed `id`.
	GetByID(id int64) (*models.BuzzFeed, error)
	// Select returns all records of the result of query execution.
	Select() ([]models.BuzzFeed, error)
	// SelectPage returns records according to given PageQuery params and the total count for the whole query.
	SelectPage(pq *db.PageQuery) ([]models.BuzzFeed, int64, error)
}

const (
	tableBuzzFeeds = "buzzFeeds"
	columnID       = "id"
)

type buzzFeedQ struct {
	parent *PGRepo
	table  db.Table

	Err error
}

func (repo *PGRepo) BuzzFeed() BuzzFeedQI {
	return &buzzFeedQ{
		parent: repo,
		table:  db.NewTable(tableBuzzFeeds, "bz", "*"),
	}
}

// Insert adds new `BuzzFeed` record to `buzzFeeds` table.
func (q *buzzFeedQ) Insert(buzzFeed *models.BuzzFeed) error {
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
func (q *buzzFeedQ) Update(id int64, buzzFeed *models.BuzzFeed) error {
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

// DeleteByID deletes row with passed `id`.
func (q *buzzFeedQ) DeleteByID(id int64) error {
	return q.parent.Exec(sq.Delete(q.table.Name).Where("id = ?", id))
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
func (q *buzzFeedQ) WithBuzzType(buzzType models.ExampleType) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("buzz_type = ?", buzzType)
	return q
}

// WithDescription adds filter by `Description` column.
func (q *buzzFeedQ) WithDescription(description string) BuzzFeedQI {
	q.table.QBuilder = q.table.QBuilder.Where("description = ?", description)
	return q
}

// WithDetails adds filter by `Details` column.
func (q *buzzFeedQ) WithDetails(details models.Feed) BuzzFeedQI {
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

// Get returns first row of the result of query execution.
func (q *buzzFeedQ) Get() (*models.BuzzFeed, error) {
	res := new(models.BuzzFeed)

	err := q.parent.Get(q.table.QBuilder, res)
	if err == sql.ErrNoRows {
		return res, nil
	}

	return res, err
}

// GetByID returns one row with passed `id`.
func (q *buzzFeedQ) GetByID(id int64) (*models.BuzzFeed, error) {
	res := new(models.BuzzFeed)
	err := q.parent.Get(q.table.QBuilder.Where("id = ?", id), res)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return res, err
}

// Select returns all records of the result of query execution.
func (q *buzzFeedQ) Select() ([]models.BuzzFeed, error) {
	res := make([]models.BuzzFeed, 0, 1)

	err := q.parent.Select(q.table.QBuilder.OrderBy(columnID), &res)
	if err == sql.ErrNoRows {
		return res, nil
	}

	return res, err
}

// SelectPage returns records according to given PageQuery params and the total count for the whole query.
func (q *buzzFeedQ) SelectPage(pq *db.PageQuery) ([]models.BuzzFeed, int64, error) {
	res := make([]models.BuzzFeed, 0, 1)

	total, err := q.table.SelectWithCount(q.parent.SQLConn, &res, pq.OrderBy, pq)
	if err == sql.ErrNoRows {
		return res, 0, nil
	}

	return res, total, err
}
