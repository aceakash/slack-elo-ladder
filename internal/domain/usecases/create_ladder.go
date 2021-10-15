package usecases

import "github.com/aceakash/slack-elo-ladder/internal/domain/models"

type CreateLadder struct {
}

func NewCreateLadder() CreateLadder {
	return CreateLadder{}
}

func (cl CreateLadder) Execute(startingScore int) *models.Ladder {
	return models.NewLadder(startingScore)
}
