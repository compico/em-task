package entity

import (
	"time"
)

type Subscription struct {
	ID          int       `db:"id"`
	ServiceName string    `db:"service_name"`
	Price       int       `db:"price"`
	UserID      string    `db:"user_id"`
	StartDate   time.Time `db:"start_date"`
}

type Subscriptions []*Subscription
