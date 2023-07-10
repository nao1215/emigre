// Package mysql provides a MySQL database.
package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/nao1215/emigre/app/domain/repository"
)

var _ repository.DB = (*DB)(nil)

// DB is a struct that contains the settings for the MySQL database.
type DB struct {
	*sql.DB
}

// NewDB returns a new DB struct.
func NewDB(config *mysql.Config) (*DB, error) {
	const (
		maxOpenConns = 10
		maxIdleConns = 10
		maxLifetime  = 10 * time.Second // 10[s]
	)

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxLifetime)
	return &DB{DB: db}, nil
}

// BeginTx begins a transaction.
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (repository.Tx, error) {
	return db.DB.BeginTx(ctx, opts)
}

// BeginReadOnlyTx begins a read-only transaction.
func (db *DB) BeginReadOnlyTx(ctx context.Context) (repository.Tx, error) {
	return db.DB.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
}

// Close closes the database and prevents new queries from starting.
func (db *DB) Close() error {
	return db.DB.Close()
}
