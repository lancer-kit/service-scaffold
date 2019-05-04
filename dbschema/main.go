package dbschema

import (
	"github.com/lancer-kit/armory/db"
	"github.com/rubenv/sql-migrate"
)

//go get -u github.com/lancer-kit/forge

//go:generate forge bindata --ignore .+\.go$ --pkg dbschema -o bindata.go -i ./...
//go:generate gofmt -w bindata.go

func Migrate(connStr string, dir db.MigrateDir) (int, error) {
	db.SetAssets(migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations",
	})
	return db.Migrate(connStr, dir)
}
