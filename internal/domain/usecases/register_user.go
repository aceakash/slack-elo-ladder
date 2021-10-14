package usecases

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/ports"
	"github.com/google/uuid"
	"time"
)

type RegisterUser struct {
	userEventStore ports.UserEventStore
}

func (ru RegisterUser) Execute(userId string) error {
	err := ru.userEventStore.StoreEvent(ports.UserEvent{
		Event: ports.Event{
			Type:      "UserRegistered",
			ID:        uuid.New().String(),
			Timestamp: time.Time{},
		},
		UserID: userId,
	})
	return err
}

func NewRegisterUser(userEventStore ports.UserEventStore) RegisterUser {
	return RegisterUser{
		userEventStore: userEventStore,
	}
}
