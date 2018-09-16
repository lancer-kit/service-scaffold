package db

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// SQLConn is a connector for interacting with the database.
type SQLConn struct {
	db     *sqlx.DB
	tx     *sqlx.Tx
	logger *logrus.Entry
}

// Clone clones the receiver, returning a new instance backed by the same
// context and db. The result will not be bound to any transaction that the
// source is currently within.
func (conn *SQLConn) Clone() *SQLConn {
	return &SQLConn{
		db:     conn.db,
		logger: conn.logger,
	}
}

// Get compile `sqq` to raw sql query, executes it and write first row into the `dest`.
func (conn *SQLConn) Get(sqq sq.Sqlizer, dest interface{}) error {
	query, args, err := sqq.ToSql()
	if err != nil {
		return err
	}
	return conn.GetRaw(dest, query, args...)
}

// GetRaw executes a raw sql query and write first row into the `dest`.
func (conn *SQLConn) GetRaw(dest interface{}, query string, args ...interface{}) error {
	query = conn.conn().Rebind(query)
	start := time.Now()
	err := conn.conn().Get(dest, query, args...)
	conn.log("select", start, query, args)

	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return err
	}

	return errors.Wrap(err, "failed to get raw")
}

// Select compile `sqq` to raw sql query, executes it, and write each row
// into dest, which must be a slice.
func (conn *SQLConn) Select(sqq sq.Sqlizer, dest interface{}) error {
	query, args, err := sqq.ToSql()
	if err != nil {
		return err
	}
	return conn.SelectRaw(dest, query, args...)
}

// SelectRaw executes a raw sql query, and write each row
// into dest, which must be a slice.
func (conn *SQLConn) SelectRaw(dest interface{}, query string, args ...interface{}) error {
	query = conn.conn().Rebind(query)
	start := time.Now()
	err := conn.conn().Select(dest, query, args...)
	conn.log("select", start, query, args)

	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return err
	}

	return errors.Wrap(err, "failed to select raw")
}

// Exec compile `sqq` to SQL and runs query.
func (conn *SQLConn) Exec(sqq sq.Sqlizer) error {
	query, args, err := sqq.ToSql()
	if err != nil {
		return err
	}

	return conn.ExecRaw(query, args...)
}

// ExecRaw runs `query` with `args`.
func (conn *SQLConn) ExecRaw(query string, args ...interface{}) error {
	query = conn.conn().Rebind(query)
	start := time.Now()
	_, err := conn.conn().Exec(query, args...)
	conn.log("exec", start, query, args)
	if err == sql.ErrNoRows {
		return err
	}

	return errors.Wrap(err, "failed to exec raw")
}

// Insert compile `sqq` to SQL and runs query. Return last inserted id
func (conn *SQLConn) Insert(sqq sq.InsertBuilder) (id interface{}, err error) {
	start := time.Now()
	err = sqq.Suffix(`RETURNING "id"`).
		RunWith(conn.baseRunner()).
		PlaceholderFormat(sq.Dollar).
		QueryRow().Scan(&id)

	query, args, _ := sqq.ToSql()
	conn.log("insert", start, query, args)

	return id, errors.Wrap(err, "failed to insert")
}

func (conn *SQLConn) conn() сonn {
	if conn.tx != nil {
		return conn.tx
	}
	return conn.db
}

func (conn *SQLConn) baseRunner() sq.BaseRunner {
	if conn.tx != nil {
		return conn.tx.Tx
	}
	return conn.db.DB
}

func (conn *SQLConn) log(typ string, start time.Time, query string, args []interface{}) {
	if conn.logger == nil {
		return
	}

	dur := time.Since(start)
	lEntry := conn.logger.
		WithFields(logrus.Fields{
			"args": args,
			"sql":  query,
			"dur":  dur.String(),
		})

	lEntry.Debugf("sql: %s", typ)
}
