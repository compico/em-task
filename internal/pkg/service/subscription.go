package service

import (
	"context"
	"errors"
	"github.com/compico/em-task/internal/pkg/entity"
	"github.com/compico/em-task/internal/pkg/filter"
	"github.com/compico/em-task/internal/pkg/query"
	"github.com/compico/em-task/internal/pkg/repository"
	"time"
)

var CrudOperationTimeout = 5 * time.Second
var ErrResourceNotFound = errors.New("resource not found")

type Subscription interface {
	Create(context.Context, *entity.Subscription) error
	Read(context.Context, int) (*entity.Subscription, error)
	Update(context.Context, int, filter.UpdateSubscription) (*entity.Subscription, error)
	Delete(ctx context.Context, id int) error
	List(ctx context.Context, page int, perPage int) (entity.Subscriptions, error)
	Sum(ctx context.Context, filter filter.SumSubscription) (*SumResult, error)
}

type SumResult struct {
	Sum int64 `json:"sum" example:"1000"`
} //@name SumResult

type subscription struct {
	repo           repository.Subscription
	requestTimeout time.Duration
}

func NewSubscription(subRepo repository.Subscription) Subscription {
	return &subscription{
		repo:           subRepo,
		requestTimeout: CrudOperationTimeout,
	}
}

func (s *subscription) Create(ctx context.Context, data *entity.Subscription) error {
	ctx, cancelFunc := context.WithTimeout(ctx, s.requestTimeout)
	defer cancelFunc()

	return s.repo.Create(ctx, data)
}

func (s *subscription) Read(ctx context.Context, id int) (*entity.Subscription, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, s.requestTimeout)
	defer cancelFunc()

	return s.repo.Read(ctx, id)
}

type UpdateSubscriptionParams struct {
	ServiceName *string
	Price       *int
	UserID      *string
	StartDate   *time.Time
}

func (s *subscription) Update(
	ctx context.Context,
	id int,
	params filter.UpdateSubscription,
) (*entity.Subscription, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, s.requestTimeout)
	defer cancelFunc()
	sub := &entity.Subscription{}

	if err := s.repo.Update(
		ctx,
		id,
		query.UpdateSubscriptionFields{
			ServiceName: params.ServiceName,
			Price:       params.Price,
			UserID:      params.UserID,
			StartDate:   params.StartDate,
		},
		sub,
	); err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *subscription) Delete(ctx context.Context, id int) error {
	ctx, cancelFunc := context.WithTimeout(ctx, s.requestTimeout)
	defer cancelFunc()

	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotAffectedRows) {
			return errors.Join(err, ErrResourceNotFound)
		}
		return err
	}

	return nil
}

func (s *subscription) List(ctx context.Context, page int, perPage int) (entity.Subscriptions, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, s.requestTimeout)
	defer cancelFunc()

	offset := (page - 1) * perPage

	return s.repo.List(ctx, perPage, offset)
}

func (s *subscription) Sum(ctx context.Context, filter filter.SumSubscription) (*SumResult, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, s.requestTimeout)
	defer cancelFunc()

	sum, err := s.repo.Sum(ctx, query.SumSubscriptionsFields{
		ServiceName: filter.ServiceName,
		UserID:      filter.UserID,
		From:        filter.From,
		To:          filter.To,
	})
	if err != nil {
		return nil, err
	}

	return &SumResult{Sum: sum.Sum}, nil
}
