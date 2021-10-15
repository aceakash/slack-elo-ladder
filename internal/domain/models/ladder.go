package models

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Type      string
	ID        string
	Timestamp time.Time
	Details   interface{}
}

type Ladder struct {
	startingScore int
	events        []Event
}

func NewLadder(startingScore int) *Ladder {
	return &Ladder{
		startingScore: startingScore,
		events:        []Event{},
	}
}

func (l *Ladder) RegisterUser(userId string) {
	l.events = append(l.events, Event{
		Type:      "UserRegistered",
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Details:   NewUserRegisteredDetails(userId),
	})
}

func (l *Ladder) ComputeTable() (Table, error) {
	table := Table{}
	for _, event := range l.events {
		if event.Type == "UserRegistered" {
			fmt.Println(event)
			eventDetails := event.Details.(UserRegisteredDetails)
			table = append(table, TableEntry{
				PlayerId: eventDetails.userId,
				Score:    l.startingScore,
			})
		}
	}
	return table, nil
}

type UserRegisteredDetails struct {
	userId string
}

func NewUserRegisteredDetails(userId string) UserRegisteredDetails {
	return UserRegisteredDetails{userId: userId}
}
