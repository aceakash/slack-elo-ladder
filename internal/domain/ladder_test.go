package domain_test

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/models"
	"github.com/aceakash/slack-elo-ladder/internal/domain/usecases"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLadderTournament(t *testing.T) {

	t.Run("When no matches have been played, and the ladder is computed, all players are at the starting score on the ladder", func(t *testing.T) {
		startingScore := 2000
		createLadderTournamentUseCase := usecases.NewCreateLadderTournament()
		ladderTournament := createLadderTournamentUseCase.Execute(startingScore, 32)

		// Given two players are registered
		registerUser := usecases.NewRegisterUser(ladderTournament)
		playerCount := 2
		err := registerUser.Execute("bruce")
		assert.NoError(t, err)
		err = registerUser.Execute("diana")
		assert.NoError(t, err)

		// And no matches have been played

		// When the ladderTournament is computed
		computeLadder := usecases.NewComputeLadder(*ladderTournament)
		ladder, err := computeLadder.Execute()
		assert.NoError(t, err)

		// Then there are two entries in the ladder
		assert.Equal(t, playerCount, len(ladder), "unexpected number of players on the ladder")
		// Then each player is at the starting score
		for _, entry := range ladder {
			assert.Equal(t, startingScore, entry.Score)
		}
	})

	t.Run("Ladder scoring", func(t *testing.T) {
		t.Run("When diana beats bruce, the ladder reflects that", func(t *testing.T) {
			startingScore := 2000
			createLadderTournamentUseCase := usecases.NewCreateLadderTournament()
			ladderTournament := createLadderTournamentUseCase.Execute(startingScore, 32)

			// Given two players are registered
			registerUser := usecases.NewRegisterUser(ladderTournament)
			err := registerUser.Execute("bruce")
			assert.NoError(t, err)
			err = registerUser.Execute("diana")
			assert.NoError(t, err)

			// When diana beats bruce
			registerMatchResult := usecases.NewRegisterMatchResult(ladderTournament)
			err = registerMatchResult.Execute("diana", "bruce")
			assert.NoError(t, err)

			// And the ladder is computed
			computeLadder := usecases.NewComputeLadder(*ladderTournament)
			ladder, err := computeLadder.Execute()
			assert.NoError(t, err)

			// Then the diana is above bruce and the scores are correct
			expectedLadder := models.Ladder{
				models.LadderEntry{PlayerId: "diana", Score: 2016},
				models.LadderEntry{PlayerId: "bruce", Score: 1984},
			}
			assert.Equal(t, expectedLadder, ladder)
		})
	})
}
