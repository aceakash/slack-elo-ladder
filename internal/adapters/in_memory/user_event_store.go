package in_memory

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/ports"
)

type UserEventStore struct {
	events []ports.UserEvent
}

func (i *UserEventStore) StoreEvent(userEvent ports.UserEvent) error {
	i.events = append(i.events, userEvent)
	return nil
}

func (i *UserEventStore) GetUserEvents() ([]ports.UserEvent, error) {
	return i.events, nil
}

func NewUserEventStore() UserEventStore {
	return UserEventStore{}
}
