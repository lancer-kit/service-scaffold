package models

import "github.com/lancer-kit/armory/db"

// RepoI is a top level interface for interaction with database.
type RepoI interface {
	db.Transactional

	BuzzFeed() BuzzFeedQI
}

// Repo implementation of the `RepoI` interface.
type Repo struct {
	*db.SQLConn
}

// NewQ returns initialized instance of the `RepoI`.
func NewQ(dbConn *db.SQLConn) *Repo {
	if dbConn == nil {
		dbConn = db.GetConnector()
	}
	return &Repo{
		SQLConn: dbConn,
	}
}
