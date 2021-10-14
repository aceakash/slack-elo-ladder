package domain_test

import (
	"github.com/aceakash/slack-elo-ladder/internal/adapters/in_memory"
	"github.com/aceakash/slack-elo-ladder/internal/domain/usecases"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLadder(t *testing.T) {

	t.Run("When no matches have been played, all players are at the starting score", func(t *testing.T) {
		// Given two players are registered
		userEventStore := in_memory.NewUserEventStore()
		registerUser := usecases.NewRegisterUser(&userEventStore)
		playerCount := 2
		err := registerUser.Execute("bruce")
		assert.NoError(t, err)
		err = registerUser.Execute("diana")
		assert.NoError(t, err)

		// And no matches have been played

		// When the ladder is computed
		startingScore := 2000
		matchEventStore := in_memory.NewMatchEventStore()
		computeLadder := usecases.NewComputeLadder(matchEventStore, &userEventStore, startingScore)
		ladder, err := computeLadder.Execute()
		assert.NoError(t, err)

		// Then each player is at the starting score
		assert.Equal(t, playerCount, len(ladder), "unexpected number of players in the ladder")
		for _, entry := range ladder {
			assert.Equal(t, startingScore, entry.Score)
		}
	})
}

