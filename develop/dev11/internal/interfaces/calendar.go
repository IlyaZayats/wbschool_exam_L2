package interfaces

import (
	"task11/internal/entities"
	"time"
)

type CalendarInterface interface {
	CreateEvent(event entities.Event) error
	UpdateEvent(event entities.Event) error
	DeleteEvent(id int) error

	GetEventsForDay(userId int, date time.Time) ([]entities.Event, error)
	GetEventsForWeek(userId int, date time.Time) ([]entities.Event, error)
	GetEventsForMonth(userId int, date time.Time) ([]entities.Event, error)
}
