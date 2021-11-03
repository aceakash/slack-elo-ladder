package models

import (
	"github.com/google/uuid"
	"time"
)

type EventType string

const (
	UserRegisteredEventType EventType = "USER_REGISTERED"
	MatchPlayedEventType    EventType = "MATCH_PLAYED"
)

type Event struct {
	Type      EventType   `json:"type"`
	ID        string      `json:"id"`
	Timestamp time.Time   `json:"timestamp"`
	Details   interface{} `json:"details"`
}

func NewUserRegisteredEvent(userId string) Event {
	return newEvent(UserRegisteredEventType, UserRegisteredEventDetails{UserId: userId})
}

func NewMatchPlayedEvent(winnerId string, loserId string) Event {
	return newEvent(MatchPlayedEventType, MatchPlayedEventDetails{Winner: winnerId, Loser: loserId})
}

func newEvent(eventType EventType, details interface{}) Event {
	return Event{
		Type:      eventType,
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Details:   details,
	}
}
