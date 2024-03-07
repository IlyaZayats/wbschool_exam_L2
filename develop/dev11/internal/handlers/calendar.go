package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"task11/internal/middleware"
	"task11/internal/requests"
	"task11/internal/responses"
	"task11/internal/services"
	"time"
)

type CalendarHandlers struct {
	service *services.CalendarService
	mux     *http.ServeMux
}

func NewCalendarHandlers(mux *http.ServeMux, service *services.CalendarService) (*CalendarHandlers, error) {
	return &CalendarHandlers{mux: mux, service: service}, nil
}

func (h *CalendarHandlers) initRoute() {
	h.mux.Handle(`POST /create_event`, middleware.Logging(http.HandlerFunc(h.CreateEvent)))
	h.mux.Handle(`POST /update_event`, middleware.Logging(http.HandlerFunc(h.UpdateEvent)))
	h.mux.Handle(`POST /delete_event`, middleware.Logging(http.HandlerFunc(h.DeleteEvent)))

	h.mux.Handle(`GET /events_for_day`, middleware.Logging(http.HandlerFunc(h.GetEventsForDay)))
	h.mux.Handle(`GET /events_for_week`, middleware.Logging(http.HandlerFunc(h.GetEventsForWeek)))
	h.mux.Handle(`GET /events_for_month`, middleware.Logging(http.HandlerFunc(h.GetEventsForMonth)))
}

func (h *CalendarHandlers) Run(host string) error {
	h.initRoute()
	return http.ListenAndServe(host, h.mux)
}

func (h *CalendarHandlers) handleError(w http.ResponseWriter, r *http.Request, err string, statusCode int) {
	log.Printf("%s %s %s %s", r.Method, r.RequestURI, "error:", err)
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error" : "%s"}`, err)))
}

func (h *CalendarHandlers) buildResponseBody(events responses.GetEventsResponse) string {
	var builder strings.Builder
	builder.WriteString(`{"result" : [`)
	for i, event := range events.Events {
		builder.WriteString(`{`)
		counter := 0
		for key, value := range event {
			switch value.(type) {
			case int:
				builder.WriteString(fmt.Sprintf(`"%s" : %v`, key, value.(int)))
				break
			case time.Time:
				builder.WriteString(fmt.Sprintf(`"%s" : "%s"`, key, value.(time.Time).Format(`2006-01-02`)))
				break
			default:
				builder.WriteString(fmt.Sprintf(`"%s" : "%s"`, key, value.(string)))
			}
			if counter != len(event)-1 {
				builder.WriteString(`, `)
			}
			counter++
		}
		if i != len(events.Events)-1 {
			builder.WriteString(`}, `)
		} else {
			builder.WriteString(`}`)
		}
	}
	builder.WriteString(`]}`)
	return builder.String()
}

func (h *CalendarHandlers) validateGetRequest(r *http.Request) (int, time.Time, error) {
	if err := r.ParseForm(); err != nil {
		return 0, time.Time{}, err
	}
	userIdString, dateString := r.Form.Get(`user_id`), r.Form.Get(`date`)
	if userIdString == "" || dateString == "" || len(r.Form) > 2 {
		return 0, time.Time{}, errors.New(`empty user_id or date fields`)
	}
	if matched, err := regexp.MatchString(`\d+`, userIdString); !matched || err != nil {
		return 0, time.Time{}, errors.New(`user_id validation error`)
	}
	userId, _ := strconv.Atoi(userIdString)
	date, err := time.Parse(`2006-01-02`, dateString)
	if err != nil {
		return 0, time.Time{}, errors.New(`date validation error`)
	}
	return userId, date, nil
}

func (h *CalendarHandlers) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	userId, date, err := h.validateGetRequest(r)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := h.service.GetEvents(requests.GetEventsForDayRequest{UserId: userId, Date: date})
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(h.buildResponseBody(events)))
}

func (h *CalendarHandlers) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	userId, date, err := h.validateGetRequest(r)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := h.service.GetEvents(requests.GetEventsForWeekRequest{UserId: userId, Date: date})
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(h.buildResponseBody(events)))
}

func (h *CalendarHandlers) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	userId, date, err := h.validateGetRequest(r)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := h.service.GetEvents(requests.GetEventsForMonthRequest{UserId: userId, Date: date})
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(h.buildResponseBody(events)))
}

func UnmarshallBody[T any](body []byte) (T, error) {
	var request T
	if err := json.Unmarshal(body, &request); err != nil {
		return request, err
	}
	return request, nil
}

func (h *CalendarHandlers) CreateEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	request, err := UnmarshallBody[requests.CreateEventRequest](body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateEvent(request); err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"result":"ok"}`))
}

func (h *CalendarHandlers) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	request, err := UnmarshallBody[requests.UpdateEventRequest](body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateEvent(request); err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"result":"ok"}`))
}

func (h *CalendarHandlers) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	request, err := UnmarshallBody[requests.DeleteEventRequest](body)
	if err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteEvent(request.Id); err != nil {
		h.handleError(w, r, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(`{"result":"ok"}`))
}
