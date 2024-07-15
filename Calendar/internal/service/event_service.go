package service

import (
	"errors"
	"github.com/NikKazzzzzz/OTUS-PRACTIC/Calendar/internal/domain"
)

var ErrDateBusy = errors.New("this date busy")

type EventService struct {
	store domain.EventStore
}

func NewEventService(store domain.EventStore) *EventService {
	return &EventService{store: store}
}

func (s *EventService) AddEvent(event domain.Event) error {
	events, _ := s.store.ListEvents()
	for _, event := range events {
		if event.StartTime.Before(event.EndTime) && event.EndTime.After(event.StartTime) {
			return ErrDateBusy
		}
	}
	return s.store.AddEvent(event)
}

func (s *EventService) DeleteEvent(eventId string) error {
	return s.store.DeleteEvent(eventId)
}

func (s *EventService) UpdateEvent(event domain.Event) error {
	events, _ := s.store.ListEvents()
	for _, e := range events {
		if e.ID != event.ID && (event.StartTime.Before(event.EndTime) && event.EndTime.After(event.StartTime)) {
			return ErrDateBusy
		}
	}
	return s.store.UpdateEvent(event)
}

func (s *EventService) ListEvents() ([]domain.Event, error) {
	return s.store.ListEvents()
}
