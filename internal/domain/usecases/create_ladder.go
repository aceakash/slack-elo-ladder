package usecases

import "github.com/aceakash/slack-elo-ladder/internal/domain/models"

type CreateLadder struct {
}

func NewCreateLadderTournament() CreateLadder {
	return CreateLadder{}
}

func (cl CreateLadder) Execute(startingScore int) *models.LadderTournament {
	return models.NewLadderTournament(startingScore)
}
