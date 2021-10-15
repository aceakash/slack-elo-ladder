package models

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Type      string
	ID        string
	Timestamp time.Time
	Details   interface{}
}

type LadderTournament struct {
	startingScore int
	events        []Event
}

func NewLadderTournament(startingScore int) *LadderTournament {
	return &LadderTournament{
		startingScore: startingScore,
		events:        []Event{},
	}
}

func (l *LadderTournament) RegisterUser(userId string) {
	l.events = append(l.events, Event{
		Type:      "UserRegistered",
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Details:   NewUserRegisteredDetails(userId),
	})
}

func (l *LadderTournament) ComputeLadder() (Ladder, error) {
	ladder := Ladder{}
	for _, event := range l.events {
		if event.Type == "UserRegistered" {
			eventDetails := event.Details.(UserRegisteredDetails)
			ladder = append(ladder, LadderEntry{
				PlayerId: eventDetails.userId,
				Score:    l.startingScore,
			})
		}
	}
	return ladder, nil
}

func (lt *LadderTournament) RegisterMatchResult(winner string, loser string) {
}

type UserRegisteredDetails struct {
	userId string
}

func NewUserRegisteredDetails(userId string) UserRegisteredDetails {
	return UserRegisteredDetails{userId: userId}
}
