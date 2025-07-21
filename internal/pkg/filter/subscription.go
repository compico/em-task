package filter

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"net/url"
	"time"
)

type UpdateSubscription struct {
	ServiceName *string
	Price       *int
	UserID      *string
	StartDate   *time.Time
}

type SumSubscription struct {
	ServiceName *string
	UserID      *string
	From        *time.Time
	To          *time.Time
}

func ListSubscriptionsFromQuery(values url.Values) (SumSubscription, error) {
	params := SumSubscription{}
	var errs []error

	if values.Has("service_name") {
		if err := params.setServiceName(values.Get("service_name")); err != nil {
			errs = append(errs, err)
		}
	}

	if values.Has("user_id") {
		if err := params.SetUserId(values.Get("user_id")); err != nil {
			errs = append(errs, err)
		}
	}

	if userID := values.Get("user_id"); userID != "" {

		params.UserID = &userID
	}

	return params, errors.Join(errs...)
}

func (filter *SumSubscription) setServiceName(serviceName string) error {
	if serviceName == "" {
		return fmt.Errorf("service_name cannot be empty")
	}

	filter.ServiceName = &serviceName

	return nil
}

func (filter *SumSubscription) SetUserId(userId string) error {
	if userId == "" {
		return fmt.Errorf("user_id cannot be empty")
	}

	if _, err := uuid.FromString(userId); err != nil {
		return fmt.Errorf("user_id must be a valid UUID")
	}

	filter.UserID = &userId

	return nil
}

func (filter *SumSubscription) SetFrom(from string) error {
	if from == "" {
		return fmt.Errorf("from cannot be empty")
	}

	t, err := time.Parse("01-2006", from)
	if err != nil {
		return fmt.Errorf("from should be a format 01-2006")
	}

	filter.From = &t

	return nil
}

func (filter *SumSubscription) SetTo(to string) error {
	if to == "" {
		return fmt.Errorf("to cannot be empty")
	}

	t, err := time.Parse("01-2006", to)
	if err != nil {
		return fmt.Errorf("to should be a format 01-2006")
	}

	filter.To = &t

	return nil
}
