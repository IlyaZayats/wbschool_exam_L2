package entities

import "time"

type Event struct {
	Id     int       `json:"id"`
	UserId int       `json:"user_id"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
}

type EventSort []Event

func (e EventSort) Len() int { return len(e) }

func (e EventSort) Less(i, j int) bool {
	return e[i].Date.Before(e[j].Date)
}

func (e EventSort) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
