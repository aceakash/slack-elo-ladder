package ports

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/models"
	"time"
)

type MatchEvent struct {
	Event
	Winner models.Player
	Loser  models.Player
}

type Event struct {
	Type      string
	ID        string
	Timestamp time.Time
}

type MatchEventStore interface {
	GetMatchEvents() ([]MatchEvent, error)
}
