package domain_test

import "testing"
import "github.com/matryer/is"

func TestLadder(t *testing.T) {
	is := is.New(t)

	t.Run("When no matches have been played, all players are at the starting score", func(t *testing.T) {
		// Given two players are registered
		registerUser := NewRegisterUser()
		registerUser.Execute("bruce")
		registerUser.Execute("diana")

		// And no matches have been played

		// When the ladder is computed
		computeLadder := NewComputeLadder()
		ladder := computeLadder.Execute()

		// Then each player is at the starting score
		for _, player := range ladder {
			is.Equal(player.Score, settings.StartingScore())
		}
	})
}
