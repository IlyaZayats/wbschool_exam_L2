package storage

import (
	"sync"
	"task11/internal/entities"
)

type CalendarStorage struct {
	m    *sync.RWMutex
	data map[int]entities.Event
}

func NewCalendarStorage(m *sync.RWMutex) *CalendarStorage {
	return &CalendarStorage{
		m:    m,
		data: map[int]entities.Event{},
	}
}

func (c *CalendarStorage) Get(key int) (entities.Event, bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	entity, ok := c.data[key]
	return entity, ok
}

func (c *CalendarStorage) Set(key int, event entities.Event) {
	c.m.Lock()
	defer c.m.Unlock()
	c.data[key] = event
}

func (c *CalendarStorage) Len() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return len(c.data)
}

func (c *CalendarStorage) Range(f func(key int, event entities.Event) bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	for k, e := range c.data {
		if !f(k, e) {
			break
		}
	}
}
