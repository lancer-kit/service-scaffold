package dbschema

import (
	"github.com/rubenv/sql-migrate"
	"gitlab.inn4science.com/gophers/service-kit/db"
)

//go get -u github.com/jteeuwen/go-bindata/...

//go:generate go-bindata -ignore .+\.go$ -pkg dbschema -o bindata.go ./...
//go:generate gofmt -w bindata.go

func Migrate(connStr string, dir db.MigrateDir) (int, error) {
	db.SetAssets(migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations",
	})
	return db.Migrate(connStr, dir)
}
