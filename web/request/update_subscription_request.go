package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// UpdateSubscriptionData структура для обновления подписки
type UpdateSubscriptionData struct {
	ServiceName *string `json:"service_name,omitempty" example:"Yandex Plus"`
	Price       *int    `json:"price,omitempty" example:"1599"`
	UserID      *string `json:"user_id,omitempty" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	StartDate   *string `json:"start_date,omitempty" example:"07-2024"`
	startDate   *time.Time
} //@name UpdateSubscriptionRequest

func UpdateSubscriptionDataFromRequest(r *http.Request) (*UpdateSubscriptionData, error) {
	data := &UpdateSubscriptionData{}
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

func (data *UpdateSubscriptionData) validate() error {
	var errs []error
	if data.ServiceName != nil && *data.ServiceName == "" {
		errs = append(errs, fmt.Errorf("service_name cannot be empty"))
	}

	if data.Price != nil && *data.Price <= 0 {
		errs = append(errs, fmt.Errorf("price must be greater than 0"))
	}

	if data.UserID != nil && *data.UserID == "" {
		errs = append(errs, fmt.Errorf("user_id cannot be empty"))
	}

	if data.StartDate != nil {
		if *data.StartDate == "" {
			errs = append(errs, fmt.Errorf("start_date cannot be empty"))
		} else {
			var err error
			if data.startDate, err = data.parseStartDate(*data.StartDate); err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errors.Join(errs...)
}

func (data *UpdateSubscriptionData) parseStartDate(startDate string) (*time.Time, error) {
	sd, err := time.Parse("01-2006", startDate)
	if err != nil {
		return nil, err
	}

	return &sd, nil
}

func (data *UpdateSubscriptionData) GetStartDate() *time.Time {
	return data.startDate
}
