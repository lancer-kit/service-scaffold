package db

import "github.com/pkg/errors"

// Transaction is generic helper method for specific Q's to implement Transaction capabilities
func (conn *SQLConn) Transaction(fn func() error) (err error) {
	if err = conn.Begin(); err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}
	defer func() {
		if err != nil {
			// swallowing rollback err,
			// should not affect data consistency
			conn.Rollback()
		}
	}()

	if err = fn(); err != nil {
		return errors.Wrap(err, "failed to execute statements")
	}

	if err = conn.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit tx")
	}

	return
}

// Begin binds this SQLConn to a new transaction.
func (conn *SQLConn) Begin() error {
	if conn.tx != nil {
		return errors.New("already in transaction")
	}

	tx, err := conn.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to begin tx")
	}
	//conn.logBegin()

	conn.tx = tx
	return nil
}

// Commit commits the current transaction.
func (conn *SQLConn) Commit() error {
	if conn.tx == nil {
		return errors.New("not in transaction")
	}

	err := conn.tx.Commit()
	//conn.logCommit()
	conn.tx = nil
	return err
}

// Rollback rolls back the current transaction
func (conn *SQLConn) Rollback() error {
	if conn.tx == nil {
		return errors.New("not in transaction")
	}

	err := conn.tx.Rollback()
	//conn.logRollback()
	conn.tx = nil
	return err
}

// IsInTx return `true` if transaction started.
func (conn *SQLConn) IsInTx() bool {
	return conn.tx == nil
}
