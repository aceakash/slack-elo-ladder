package usecases

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/models"
	"github.com/aceakash/slack-elo-ladder/internal/domain/ports"
)

type ComputeLadder struct {
	UserEventStore  ports.UserEventStore
	MatchEventStore ports.MatchEventStore
	startingScore   int
}

func NewComputeLadder(matchEventStore ports.MatchEventStore, userEventStore ports.UserEventStore, startingScore int) ComputeLadder {
	return ComputeLadder{
		MatchEventStore: matchEventStore,
		UserEventStore: userEventStore,
		startingScore: startingScore,
	}
}

func (cl ComputeLadder) Execute() (models.Ladder, error) {
	userEvents, err := cl.UserEventStore.GetUserEvents()
	if err != nil {
		panic("todo: implement me")
	}

	ladder := models.Ladder{}
	for _, userEvent := range userEvents {
		if userEvent.Type == "UserRegistered" {
			ladder = append(ladder, models.LadderEntry{
				Player: models.Player{ID: userEvent.UserID},
				Score:  cl.startingScore,
			})
		}
	}

	return ladder, nil
}


