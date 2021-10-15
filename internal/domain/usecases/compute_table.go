package usecases

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/models"
)

type ComputeTable struct {
	ladder models.Ladder
}

func NewComputeTable(ladder models.Ladder) ComputeTable {
	return ComputeTable{
		ladder: ladder,
	}
}

func (cl ComputeTable) Execute() (models.Table, error) {
	return cl.ladder.ComputeTable()

	//userEvents, err := cl.UserEventStore.GetUserEvents()
	//if err != nil {
	//	panic("todo: implement me")
	//}
	//
	//ladder := models.Ladder{}
	//for _, userEvent := range userEvents {
	//	if userEvent.Type == "UserRegistered" {
	//		ladder = append(ladder, models.TableEntry{
	//			Player: models.Player{ID: userEvent.UserID},
	//			Score:  cl.startingScore,
	//		})
	//	}
	//}

	return models.Table{}, nil
}
