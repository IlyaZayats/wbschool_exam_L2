package requests

import (
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const dateLayout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(dateLayout, s)
	return
}

type CreateEventRequest struct {
	UserId int        `json:"user_id"`
	Name   string     `json:"name"`
	Date   CustomTime `json:"date"`
}

type UpdateEventRequest struct {
	Id     int        `json:"id"`
	UserId int        `json:"user_id"`
	Name   string     `json:"name"`
	Date   CustomTime `json:"date"`
}

type DeleteEventRequest struct {
	Id int `json:"id"`
}

type GetEventsForDayRequest struct {
	UserId int       `json:"user_id"`
	Date   time.Time `json:"date"`
}

type GetEventsForWeekRequest struct {
	UserId int       `json:"user_id"`
	Date   time.Time `json:"date"`
}

type GetEventsForMonthRequest struct {
	UserId int       `json:"user_id"`
	Date   time.Time `json:"date"`
}
