package domain

import "time"

type Event struct {
	ID          string
	Title       string
	Description string
	StartTime   time.Time
	EndTime     time.Time
}

type EventStore interface {
	AddEvent(event Event) error
	DeleteEvent(eventID string) error
	UpdateEvent(event Event) error
	ListEvents() ([]Event, error)
}
