package repo

import "github.com/lancer-kit/armory/db"

// PGRepoI is a top level interface for interaction with database.
type PGRepoI interface {
	db.Transactional
	Clone() PGRepoI
	BuzzFeed() BuzzFeedQI
}

// PGRepo implementation of the `PGRepoI` interface.
type PGRepo struct {
	*db.SQLConn
}

// NewPGRepo returns initialized instance of the `PGRepoI`.
func NewPGRepo(dbConn *db.SQLConn) PGRepoI {
	if dbConn == nil {
		dbConn = db.GetConnector()
	}

	return &PGRepo{
		SQLConn: dbConn,
	}
}

func (repo *PGRepo) Clone() PGRepoI {
	return &PGRepo{
		SQLConn: repo.SQLConn.Clone(),
	}
}
