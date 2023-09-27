package repository

import (
	"context"
	"database/sql"
)

// Preparer is an interface for sql.DB and sql.Tx. It is used to prepare sql.Stmt.
type Preparer interface {
	// PrepareContext creates a prepared statement for later queries or executions.
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// Execer is an interface for sql.DB and sql.Tx. It is a combination of Preparer and Execer.
type Execer interface {
	// ExecContext executes a query without returning any rows.
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Preparer
}

// Queryer is an interface for sql.DB and sql.Tx. It is used to query rows.
type Queryer interface {
	// QueryContext executes a query that returns rows, typically a SELECT.
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// QueryRowContext executes a query that is expected to return at most one row.
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Preparer
}

// QueryExecer is an interface for sql.DB and sql.Tx. It is a combination of Queryer and Execer.
type QueryExecer interface {
	Queryer
	Execer
}

// Tx is an interface for sql.Tx.
type Tx interface {
	// Commit commits the transaction.
	Commit() error
	// Rollback aborts the transaction.
	Rollback() error
	QueryExecer
}

// TxBeginner is an interface for sql.DB. It is used to begin transaction.
type TxBeginner interface {
	// BeginTx starts a transaction.
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error)
}

// ReadOnlyTxBegginer is an interface for sql.DB. It is used to begin read only transaction.
type ReadOnlyTxBegginer interface {
	// BeginReadOnlyTx starts a read only transaction.
	BeginReadOnlyTx(ctx context.Context) (Tx, error)
}

// Closer is an interface for sql.DB and sql.Tx. It is used to close database connection.
type Closer interface {
	Close() error
}

// DB is an interface for sql.DB.
type DB interface {
	TxBeginner
	ReadOnlyTxBegginer
	Queryer
	Execer
	Closer
}
