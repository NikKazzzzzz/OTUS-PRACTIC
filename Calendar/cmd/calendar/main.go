package main

import (
	"fmt"
	"github.com/NikKazzzzzz/OTUS-PRACTIC/Calendar/internal/domain"
	"github.com/NikKazzzzzz/OTUS-PRACTIC/Calendar/internal/repository"
	"github.com/NikKazzzzzz/OTUS-PRACTIC/Calendar/internal/service"
	"time"
)

func main() {
	store := repository.NewInMemoryStore()
	eventService := service.NewEventService(store)

	event1 := domain.Event{
		ID:          "1",
		Title:       "Meeting",
		Description: "Team meeting",
		StartTime:   time.Date(2020, time.April, 28, 9, 0, 0, 0, time.UTC),
		EndTime:     time.Date(2020, time.April, 28, 10, 0, 0, 0, time.UTC),
	}

	err := eventService.AddEvent(event1)
	if err != nil {
		fmt.Println("Error adding event:", err)
	} else {
		fmt.Println("Added event")
	}

	events, _ := eventService.ListEvents()
	fmt.Println("Current events", events)
}
