package services

import (
	"sort"
	"task11/internal/entities"
	"task11/internal/interfaces"
	"task11/internal/requests"
	"task11/internal/responses"
)

type CalendarService struct {
	repo interfaces.CalendarInterface
}

func NewCalendarService(repo interfaces.CalendarInterface) (*CalendarService, error) {
	return &CalendarService{repo: repo}, nil
}

func (s *CalendarService) CreateEvent(event requests.CreateEventRequest) error {
	return s.repo.CreateEvent(entities.Event{UserId: event.UserId, Name: event.Name, Date: event.Date.Time})
}

func (s *CalendarService) UpdateEvent(event requests.UpdateEventRequest) error {
	return s.repo.UpdateEvent(entities.Event{Id: event.Id, UserId: event.UserId, Name: event.Name, Date: event.Date.Time})
}

func (s *CalendarService) DeleteEvent(id int) error {
	return s.repo.DeleteEvent(id)
}

func (s *CalendarService) GetEvents(req interface{}) (responses.GetEventsResponse, error) {
	var events []entities.Event
	var err error
	switch req.(type) {
	case requests.GetEventsForDayRequest:
		events, err = s.repo.GetEventsForDay(
			req.(requests.GetEventsForDayRequest).UserId,
			req.(requests.GetEventsForDayRequest).Date)
		break
	case requests.GetEventsForWeekRequest:
		events, err = s.repo.GetEventsForWeek(
			req.(requests.GetEventsForWeekRequest).UserId,
			req.(requests.GetEventsForWeekRequest).Date)
		break
	case requests.GetEventsForMonthRequest:
		events, err = s.repo.GetEventsForMonth(
			req.(requests.GetEventsForMonthRequest).UserId,
			req.(requests.GetEventsForMonthRequest).Date)
		break

	}
	if err != nil {
		return responses.GetEventsResponse{}, err
	}
	sort.Sort(entities.EventSort(events))
	var resp responses.GetEventsResponse
	for _, e := range events {
		resp.Events = append(resp.Events, s.EntityToMap(e))
	}
	return resp, nil
}

func (s *CalendarService) EntityToMap(e entities.Event) map[string]interface{} {
	event := map[string]interface{}{}
	event[`id`] = e.Id
	event[`user_id`] = e.UserId
	event[`name`] = e.Name
	event[`date`] = e.Date
	return event
}
