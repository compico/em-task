package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Ping(ctx context.Context) error
}

type database struct {
	conn *pgxpool.Pool
}

func NewDatabase(ctx context.Context, conn *pgxpool.Pool) DB {
	return &database{conn: conn}
}

func (db *database) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return db.conn.Query(ctx, sql, args...)
}

func (db *database) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return db.conn.QueryRow(ctx, sql, args...)
}

func (db *database) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return db.conn.Exec(ctx, sql, args...)
}

func (db *database) Begin(ctx context.Context) (pgx.Tx, error) {
	return db.conn.Begin(ctx)
}

func (db *database) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return db.conn.BeginTx(ctx, txOptions)
}

func (db *database) Ping(ctx context.Context) error {
	if db.conn == nil {
		return fmt.Errorf("database has not been initialized")
	}
	return db.conn.Ping(ctx)
}
