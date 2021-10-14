package in_memory

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/ports"
)

type InMemoryMatchEventStore struct {
}

func (i InMemoryMatchEventStore) GetMatchEvents() ([]ports.MatchEvent, error) {
	panic("implement me")
}

func NewMatchEventStore() InMemoryMatchEventStore {
	return InMemoryMatchEventStore{}
}
