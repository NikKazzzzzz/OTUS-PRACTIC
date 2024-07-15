package service

import (
	"errors"
	"github.com/NikKazzzzzz/OTUS-PRACTIC/Calendar/internal/domain"
	"github.com/NikKazzzzzz/OTUS-PRACTIC/Calendar/internal/repository"
	"testing"
	"time"
)

func TestAddEvent(t *testing.T) {
	store := repository.NewInMemoryStore()
	service := NewEventService(store)

	event1 := domain.Event{
		ID:        "1",
		Title:     "Meeting",
		StartTime: time.Date(2020, time.April, 29, 9, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, time.April, 29, 10, 0, 0, 0, time.UTC),
	}

	event2 := domain.Event{
		ID:        "2",
		Title:     "Workshop",
		StartTime: time.Date(2020, time.April, 29, 9, 30, 0, 0, time.UTC),
		EndTime:   time.Date(2020, time.April, 29, 11, 0, 0, 0, time.UTC),
	}

	err := service.AddEvent(event1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = service.AddEvent(event2)
	if !errors.Is(err, ErrDateBusy) {
		t.Fatalf("expected ErrDateBusy, got %v", err)
	}
}

func TestDeleteEvent(t *testing.T) {
	store := repository.NewInMemoryStore()
	service := NewEventService(store)

	event := domain.Event{
		ID:        "1",
		Title:     "Meeting",
		StartTime: time.Date(2020, time.April, 29, 9, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, time.April, 29, 10, 0, 0, 0, time.UTC),
	}

	_ = service.AddEvent(event)
	err := service.DeleteEvent("1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = service.DeleteEvent("1")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestUpdateEvent(t *testing.T) {
	store := repository.NewInMemoryStore()
	service := NewEventService(store)

	event := domain.Event{
		ID:        "1",
		Title:     "Meeting",
		StartTime: time.Date(2020, time.April, 29, 9, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, time.April, 29, 10, 0, 0, 0, time.UTC),
	}

	_ = service.AddEvent(event)

	updatedEvent := domain.Event{
		ID:        "1",
		Title:     "Updated Meeting",
		StartTime: time.Date(2020, time.April, 29, 10, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2020, time.April, 29, 11, 0, 0, 0, time.UTC),
	}

	err := service.UpdateEvent(updatedEvent)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	events, _ := service.ListEvents()
	if len(events) != 1 || events[0].Title != "Updated Meeting" {
		t.Fatalf("expected update event, got %+v", events[0])
	}
}
