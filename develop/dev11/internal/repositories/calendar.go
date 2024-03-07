package repositories

import (
	"fmt"
	"log"
	"task11/internal/entities"
	"task11/internal/storage"
	"time"
)

type CalendarRepository struct {
	db *storage.CalendarStorage
}

func NewCalendarRepository(db *storage.CalendarStorage) (*CalendarRepository, error) {
	return &CalendarRepository{db: db}, nil
}

func (r *CalendarRepository) CreateEvent(event entities.Event) error {
	event.Id = r.db.Len() + 1
	r.db.Set(event.Id, event)
	log.Print(r.db.Len())
	return nil
}

func (r *CalendarRepository) UpdateEvent(event entities.Event) error {
	if _, ok := r.db.Get(event.Id); ok {
		r.db.Set(event.Id, event)
	} else {
		return fmt.Errorf("event with id=%v is not defined", event.Id)
	}
	return nil
}

func (r *CalendarRepository) DeleteEvent(id int) error {
	if _, ok := r.db.Get(id); ok {
		r.db.Set(id, entities.Event{})
	} else {
		return fmt.Errorf("event with id=%v is not defined", id)
	}
	return nil
}

func (r *CalendarRepository) GetEventsForDay(userId int, date time.Time) ([]entities.Event, error) {
	result := make([]entities.Event, 0)
	r.db.Range(func(key int, event entities.Event) bool {
		if event.Date == date && event.UserId == userId {
			result = append(result, event)
		}
		return true
	})
	return result, nil
}

func (r *CalendarRepository) GetEventsForWeek(userId int, date time.Time) ([]entities.Event, error) {
	result := make([]entities.Event, 0)
	year, week := date.ISOWeek()
	r.db.Range(func(key int, event entities.Event) bool {
		yearTemp, weekTemp := event.Date.ISOWeek()
		if yearTemp == year && weekTemp == week && event.UserId == userId {
			result = append(result, event)
		}
		return true
	})
	return result, nil
}

func (r *CalendarRepository) GetEventsForMonth(userId int, date time.Time) ([]entities.Event, error) {
	result := make([]entities.Event, 0)
	month, year := date.Month(), date.Year()
	r.db.Range(func(key int, event entities.Event) bool {
		yearTemp, monthTemp := event.Date.Year(), event.Date.Month()
		if yearTemp == year && monthTemp == month && event.UserId == userId {
			result = append(result, event)
		}
		return true
	})
	return result, nil
}
