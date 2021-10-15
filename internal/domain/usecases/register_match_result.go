package usecases

import "github.com/aceakash/slack-elo-ladder/internal/domain/models"

type RegisterMatchResult struct {
	ladderTournament *models.LadderTournament
}

func (r RegisterMatchResult) Execute(winner string, loser string) error {
	r.ladderTournament.RegisterMatchResult(winner, loser)
	return nil
}

func NewRegisterMatchResult(ladderTournament *models.LadderTournament) RegisterMatchResult {
	return RegisterMatchResult{
		ladderTournament: ladderTournament,
	}
}
