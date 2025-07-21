package repository

import (
	"context"
	"errors"
	"github.com/compico/em-task/internal/pkg/entity"
	"github.com/compico/em-task/internal/pkg/query"
)

type Subscription interface {
	Create(ctx context.Context, sub *entity.Subscription) error
	Read(ctx context.Context, id int) (*entity.Subscription, error)
	Update(
		ctx context.Context,
		id int,
		fields query.UpdateSubscriptionFields,
		sub *entity.Subscription,
	) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, limit int, offset int) (entity.Subscriptions, error)
	Sum(ctx context.Context, fields query.SumSubscriptionsFields) (*SumResult, error)
}

var ErrNotAffectedRows = errors.New("not affected rows")

type SumResult struct {
	Sum int64
}
