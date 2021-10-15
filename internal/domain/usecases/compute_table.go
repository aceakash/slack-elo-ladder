package usecases

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/models"
)

type ComputeLadder struct {
	ladderTournament models.LadderTournament
}

func NewComputeLadder(ladderTournament models.LadderTournament) ComputeLadder {
	return ComputeLadder{
		ladderTournament: ladderTournament,
	}
}

func (cl ComputeLadder) Execute() (models.Ladder, error) {
	return cl.ladderTournament.ComputeLadder()
}
