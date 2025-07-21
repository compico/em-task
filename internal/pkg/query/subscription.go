package query

import (
	"fmt"
	"strings"
	"time"
)

type UpdateSubscriptionFields struct {
	ServiceName *string
	Price       *int
	UserID      *string
	StartDate   *time.Time
}

func (fields *UpdateSubscriptionFields) ToQuery(id int) (string, []interface{}, error) {
	var setClauses []string
	var args []interface{}
	i := 1

	if fields.ServiceName != nil {
		setClauses = append(setClauses, fmt.Sprintf("service_name = $%d", i))
		args = append(args, *fields.ServiceName)
		i++
	}

	if fields.Price != nil {
		setClauses = append(setClauses, fmt.Sprintf("price = $%d", i))
		args = append(args, *fields.Price)
		i++
	}

	if fields.UserID != nil {
		setClauses = append(setClauses, fmt.Sprintf("user_id = $%d", i))
		args = append(args, *fields.UserID)
		i++
	}

	if fields.StartDate != nil {
		setClauses = append(setClauses, fmt.Sprintf("start_date = $%d", i))
		args = append(args, *fields.StartDate)
		i++
	}

	if len(setClauses) == 0 {
		return "", nil, fmt.Errorf("ничего не передано для обновления")
	}

	args = append(args, id)
	query := fmt.Sprintf(`
        UPDATE subscriptions
        SET %s
        WHERE id = $%d
        RETURNING id, service_name, price, user_id, start_date
    `, strings.Join(setClauses, ", "), len(args))

	return query, args, nil
}

type SumSubscriptionsFields struct {
	ServiceName *string
	UserID      *string
	From        *time.Time
	To          *time.Time
}

func (fields *SumSubscriptionsFields) ToQuery() (string, []interface{}, error) {
	var whereClauses []string
	var args []interface{}
	i := 1

	if fields.ServiceName != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("service_name = $%d", i))
		args = append(args, *fields.ServiceName)
		i++
	}

	if fields.UserID != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("user_id = $%d", i))
		args = append(args, *fields.UserID)
		i++
	}

	if fields.From != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("start_date >= $%d", i))
		args = append(args, *fields.From)
		i++
	}

	if fields.To != nil {
		whereClauses = append(whereClauses, fmt.Sprintf("start_date <= $%d", i))
		args = append(args, *fields.To)
		i++
	}

	var query string
	if len(whereClauses) == 0 {
		query = "SELECT COALESCE(SUM(price), 0) FROM subscriptions"
	} else {
		query = fmt.Sprintf(`
            SELECT COALESCE(SUM(price), 0)
            FROM subscriptions
            WHERE %s
        `, strings.Join(whereClauses, " AND "))
	}

	return query, args, nil
}
