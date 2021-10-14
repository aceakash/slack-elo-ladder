package usecases

import "github.com/aceakash/slack-elo-ladder/internal/domain/models"

type ComputeLadder struct {
}

func (cl ComputeLadder) Execute() (models.Ladder, error) {
	return models.Ladder{
		{Score: 2000},
		{Score: 2000},
	}, nil
}

func NewComputeLadder() ComputeLadder {
	return ComputeLadder{}
}
