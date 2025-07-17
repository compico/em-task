package entity

import (
	"github.com/compico/em-task/internal/pkg/types"
)

type Subscription struct {
	ID          int             `db:"id"`
	ServiceName string          `db:"service_name"`
	Price       int             `db:"price"`
	UserID      string          `db:"user_id"`
	StartDate   types.MonthYear `db:"start_date"`
}
