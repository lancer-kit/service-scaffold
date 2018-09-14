package db

import (
	"database/sql"

	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

// MigrateDir represents a direction in which to perform schema migrations.
type MigrateDir string

const (
	// MigrateUp causes migrations to be run in the "up" direction.
	MigrateUp MigrateDir = "up"
	// MigrateDown causes migrations to be run in the "down" direction.
	MigrateDown MigrateDir = "down"
)

var directions = map[MigrateDir]migrate.MigrationDirection{
	MigrateUp:   migrate.Up,
	MigrateDown: migrate.Down,
}

// migrations represents all of the schema migration for service
var migrations *migrate.AssetMigrationSource

// SetAssets is a function for injection of precompiled by bindata migrations files.
func SetAssets(assets migrate.AssetMigrationSource) {
	migrations = &assets
}

// Migrate connects to the database and applies migrations.
func Migrate(connStr string, dir MigrateDir) (int, error) {
	if migrations == nil {
		return 0, errors.New("migrations isn't set")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return 0, errors.Wrap(err, "unable to connect to the database")
	}

	return migrate.Exec(db, "postgres", migrations, directions[dir])
}
