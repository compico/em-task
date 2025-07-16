package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"sync"
)

type Url string

var instance *Database

type Database struct {
	conn *pgx.Conn
}

func NewConnection(ctx context.Context, url Url) (*Database, error) {
	var err error
	sync.OnceFunc(func() {
		instance = &Database{}
		instance.conn, err = pgx.Connect(ctx, string(url))
	})()

	return instance, err
}

func (c *Database) Close(ctx context.Context) error {
	return c.conn.Close(ctx)
}
