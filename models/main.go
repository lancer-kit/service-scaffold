package models

import "gitlab.inn4science.com/gophers/service-kit/db"

// QI is a top level interface for interaction with database.
type QI interface {
	db.Transactional

	BuzzFeed() BuzzFeedQI
}

// Q implementation of the `QI` interface.
type Q struct {
	*db.SQLConn
}

// NewQ returns initialized instance of the `QI`.
func NewQ(dbConn *db.SQLConn) *Q {
	if dbConn == nil {
		dbConn = db.GetConnector()
	}
	return &Q{
		SQLConn: dbConn,
	}
}
