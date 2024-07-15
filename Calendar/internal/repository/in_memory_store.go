package repository

import (
	"errors"
	"github.com/NikKazzzzzz/OTUS-PRACTIC/Calendar/internal/domain"
	"sync"
)

type InMemoryStore struct {
	mu     sync.Mutex
	events map[string]domain.Event
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		events: make(map[string]domain.Event),
	}
}

func (s *InMemoryStore) AddEvent(event domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event.ID] = event
	return nil
}

func (s *InMemoryStore) DeleteEvent(eventID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.events[eventID]; !exists {
		return errors.New("event not found")
	}
	delete(s.events, eventID)
	return nil
}

func (s *InMemoryStore) UpdateEvent(event domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.events[event.ID]; !exists {
		return errors.New("event not found")
	}
	s.events[event.ID] = event
	return nil
}

func (s *InMemoryStore) ListEvents() ([]domain.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	events := make([]domain.Event, 0, len(s.events))
	for _, event := range s.events {
		events = append(events, event)
	}
	return events, nil
}
