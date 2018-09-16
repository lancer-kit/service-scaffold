package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// сonn represents a connector to a database.
type сonn interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Rebind(sql string) string
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	Select(dest interface{}, query string, args ...interface{}) error
}

var conn *SQLConn

// Init initializes new connector with database.
func Init(dbConnStr string, logger *logrus.Entry) error {
	if conn == nil {
		conn = &SQLConn{}
	}

	var err error
	conn.logger = logger
	conn.db, err = sqlx.Connect("postgres", dbConnStr)
	if err != nil {
		return errors.Wrap(err, "failed to init db connection")
	}

	return nil
}

// GetConnector returns an instance of the SQLConn.
func GetConnector() *SQLConn {
	return conn.Clone()
}
