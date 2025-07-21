package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"net/http"
	"time"
)

// CreateSubscriptionData структура для создания подписки
type CreateSubscriptionData struct {
	ServiceName string `json:"service_name" example:"Yandex Plus" binding:"required"`
	Price       int    `json:"price" example:"1299" binding:"required"`
	UserID      string `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000" binding:"required"`
	StartDate   string `json:"start_date" example:"03-2024" binding:"required"`
	startDate   time.Time
} //@name CreateSubscriptionRequest

func CreateSubscriptionDataFromRequest(r *http.Request) (*CreateSubscriptionData, error) {
	data := &CreateSubscriptionData{}
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, data.validate()
}

func (d *CreateSubscriptionData) validate() error {
	var errs []error

	if d.ServiceName == "" {
		errs = append(errs, errors.New("service name is required"))
	}

	if d.Price <= 0 {
		errs = append(errs, errors.New("price must be greater than zero"))
	}

	if startDate, err := time.Parse("01-2006", d.StartDate); err != nil {
		errs = append(errs, fmt.Errorf("invalid start date format"))
	} else {
		d.startDate = startDate
	}

	if _, err := uuid.FromString(d.UserID); err != nil {
		errs = append(errs, fmt.Errorf("invalid user ID format"))
	}

	return errors.Join(errs...)
}

func (d *CreateSubscriptionData) GetStartDateAsTime() time.Time {
	return d.startDate
}
