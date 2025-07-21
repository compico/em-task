package pgrepo

import (
	"context"
	"github.com/compico/em-task/internal/pkg/entity"
	"github.com/compico/em-task/internal/pkg/query"
	"github.com/compico/em-task/internal/pkg/repository"
	"github.com/compico/em-task/pkg/logger"
	"github.com/compico/em-task/pkg/postgres"
)

type subscription struct {
	logger logger.Logger
	db     postgres.DB
}

func NewSubscriptionRepository(logger logger.Logger, db postgres.DB) repository.Subscription {
	return &subscription{
		logger: logger,
		db:     db,
	}
}

func (r *subscription) Create(ctx context.Context, sub *entity.Subscription) error {
	q := `
		INSERT INTO subscriptions (service_name, price, user_id, start_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	r.logger.DebugContext(ctx, "database_query", "query", q, "subscription", sub)
	return r.db.QueryRow(
		ctx,
		`
			INSERT INTO subscriptions (service_name, price, user_id, start_date)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`,
		sub.ServiceName,
		sub.Price,
		sub.UserID,
		sub.StartDate,
	).Scan(&sub.ID)
}

func (r *subscription) Read(ctx context.Context, id int) (*entity.Subscription, error) {
	q := `SELECT * FROM subscriptions WHERE id = $1`
	r.logger.DebugContext(ctx, "database query", "query", q, "id", id)

	row := r.db.QueryRow(ctx, q, id)

	sub := &entity.Subscription{}

	return sub, row.Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate)
}

func (r *subscription) Update(
	ctx context.Context,
	id int,
	fields query.UpdateSubscriptionFields,
	sub *entity.Subscription,
) error {
	queryStr, args, err := fields.ToQuery(id)
	if err != nil {
		return err
	}
	r.logger.DebugContext(ctx, "database_query", "query", queryStr, "update_fields", fields)

	row := r.db.QueryRow(ctx, queryStr, args...)
	if err := row.Scan(
		&sub.ID,
		&sub.ServiceName,
		&sub.Price,
		&sub.UserID,
		&sub.StartDate,
	); err != nil {
		return err
	}

	return nil
}

func (r *subscription) Delete(ctx context.Context, id int) error {
	q := `DELETE FROM subscriptions WHERE id = $1`
	r.logger.DebugContext(ctx, "database_query", "query", q, "id", id)

	tags, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tags.RowsAffected() < 1 {
		return repository.ErrNotAffectedRows
	}

	return nil
}

func (r *subscription) List(ctx context.Context, limit int, offset int) (entity.Subscriptions, error) {
	q := `
		SELECT id, service_name, price, user_id, start_date
		FROM subscriptions ORDER BY id LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r.logger.DebugContext(ctx, "database query", "query", q, "limit", limit, "offset", offset)
	list := entity.Subscriptions{}
	for rows.Next() {
		sub := &entity.Subscription{}

		err := rows.Scan(
			&sub.ID,
			&sub.ServiceName,
			&sub.Price,
			&sub.UserID,
			&sub.StartDate,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, sub)
	}

	return list, nil
}

func (r *subscription) Sum(ctx context.Context, fields query.SumSubscriptionsFields) (*repository.SumResult, error) {
	queryStr, args, err := fields.ToQuery()
	if err != nil {
		return nil, err
	}
	r.logger.DebugContext(ctx, "database_query", "query", queryStr, "sum_fields", fields)

	var sum int64
	row := r.db.QueryRow(ctx, queryStr, args...)
	if err := row.Scan(&sum); err != nil {
		return nil, err
	}

	return &repository.SumResult{
		Sum: sum,
	}, nil
}
