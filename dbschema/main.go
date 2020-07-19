package dbschema

import (
	"github.com/lancer-kit/armory/db"
	migrate "github.com/lancer-kit/sql-migrate"
	"github.com/pkg/errors"
)

// go get -u github.com/lancer-kit/forge

//go:generate forge bindata --ignore .+\.go$ --pkg dbschema -o bindata.go -i ./...
//go:generate gofmt -w bindata.go

func Migrate(connStr string, dir db.MigrateDir) (int, error) {
	return db.MigrateSet(connStr, db.DriverPostgres, dir, db.Migrations{
		Table:           "_migrations",
		EnablePatchMode: false,
		IgnoreUnknown:   false,
		Assets: &migrate.AssetMigrationSource{
			Asset:    Asset,
			AssetDir: AssetDir,
			Dir:      "pg",
		},
	})
}

func DropSchema(connStr string) error {
	sqlConn, err := db.NewConnector(db.Config{ConnURL: connStr}, nil)
	if err != nil {
		return errors.Wrap(err, "unable to init dbConn")
	}

	err = sqlConn.ExecRaw(`DROP SCHEMA public CASCADE; CREATE SCHEMA public;`, nil)
	if err != nil {
		return errors.Wrap(err, "drop query failed")
	}
	return nil
}
