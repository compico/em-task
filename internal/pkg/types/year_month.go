package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type MonthYear struct {
	Month int
	Year  int
}

func NewMonthYear(s string) (MonthYear, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return MonthYear{}, fmt.Errorf("invalid format: expected MM-YYYY, got %s", s)
	}

	month, err := strconv.Atoi(parts[0])
	if err != nil || month < 1 || month > 12 {
		return MonthYear{}, fmt.Errorf("invalid month: %s", parts[0])
	}

	year, err := strconv.Atoi(parts[1])
	if err != nil {
		return MonthYear{}, fmt.Errorf("invalid year: %s", parts[1])
	}

	return MonthYear{Month: month, Year: year}, nil
}

func (my *MonthYear) String() string {
	return fmt.Sprintf("%02d-%04d", my.Month, my.Year)
}

func (my *MonthYear) ToDate() time.Time {
	return time.Date(my.Year, time.Month(my.Month), 1, 0, 0, 0, 0, time.UTC)
}

func MonthYearFromDate(t time.Time) MonthYear {
	return MonthYear{
		Month: int(t.Month()),
		Year:  t.Year(),
	}
}

func (my *MonthYear) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*my = MonthYearFromDate(v)
		return nil
	case string:
		parsed, err := NewMonthYear(v)
		if err != nil {
			return err
		}
		*my = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan %T into MonthYear", value)
	}
}

func (my *MonthYear) Value() (driver.Value, error) {
	return my.ToDate(), nil
}
