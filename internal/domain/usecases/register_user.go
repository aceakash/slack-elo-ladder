package usecases

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/models"
)

type RegisterUser struct {
	ladder *models.LadderTournament
}

func (ru RegisterUser) Execute(userId string) error {
	ru.ladder.RegisterUser(userId)
	//err := ru.userEventStore.StoreEvent(ports.UserEvent{
	//	Event: ports.Event{
	//		Type:      "UserRegistered",
	//		ID:        uuid.New().String(),
	//		Timestamp: time.Time{},
	//	},
	//	UserID: userId,
	//})
	return nil
}

func NewRegisterUser(ladder *models.LadderTournament) RegisterUser {
	return RegisterUser{
		ladder: ladder,
	}
}
