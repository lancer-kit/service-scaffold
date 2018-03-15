package dbschema

import (
	"database/sql"

	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

//go:generate go-bindata -ignore .+\.go$ -pkg dbschema -o bindata.go ./...
//go:generate gofmt -w bindata.go

// MigrateDir represents a direction in which to perform schema migrations.
type MigrateDir string

const (
	// MigrateUp causes migrations to be run in the "up" direction.
	MigrateUp MigrateDir = "up"
	// MigrateDown causes migrations to be run in the "down" direction.
	MigrateDown MigrateDir = "down"
)

// migrations represents all of the schema migration for horizon
var migrations = &migrate.AssetMigrationSource{
	Asset:    Asset,
	AssetDir: AssetDir,
	Dir:      "migrations",
}

var directions = map[MigrateDir]migrate.MigrationDirection{
	MigrateUp:   migrate.Up,
	MigrateDown: migrate.Down,
}

// Migrate connects to the database and applies migrations.
func Migrate(connStr string, dir MigrateDir) (int, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return 0, errors.Wrap(err, "unable to connect to the database")
	}

	return migrate.Exec(db, "postgres", migrations, directions[dir])
}
