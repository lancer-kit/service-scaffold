package models

import "gitlab.inn4science.com/gophers/service-kit/db"

// QI is a top level interface for interaction with database.
type QI interface {
	TxBegin() error
	TxCommit() error
	TxRollback() error
	IsInTx() bool
	BuzzFeed() *BuzzFeedQ
}

// Q implementation of the `QI` interface.
type Q struct {
	DBConn *db.SQLConn
}

// NewQ returns initialized instance of the `QI`.
func NewQ(dbConn *db.SQLConn) *Q {
	if dbConn == nil {
		dbConn = db.GetConnector()
	}
	return &Q{
		DBConn: dbConn,
	}
}

// IsInTx checks whether the transaction was started.
func (q *Q) IsInTx() bool {
	return q.DBConn.IsInTx()
}

// TxBegin starts new database transaction.
func (q *Q) TxBegin() error {
	return q.DBConn.Begin()
}

//TxCommit commits the current database transaction.
func (q *Q) TxCommit() error {
	return q.DBConn.Commit()
}

// TxRollback rolls back the current database transaction.
func (q *Q) TxRollback() error {
	return q.DBConn.Rollback()
}

func (q *Q) BuzzFeed() *BuzzFeedQ {
	return NewBuzzFeedQ(q)
}
