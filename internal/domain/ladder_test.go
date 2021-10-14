package domain_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLadder(t *testing.T) {

	t.Run("When no matches have been played, all players are at the starting score", func(t *testing.T) {
		// Given two players are registered
		registerUser := NewRegisterUser()
		playerCount := 2
		err := registerUser.Execute("bruce")
		assert.NoError(t, err)
		err = registerUser.Execute("diana")
		assert.NoError(t, err)

		// And no matches have been played

		// When the ladder is computed
		startingScore := 2000
		computeLadder := NewComputeLadder()
		ladder, err := computeLadder.Execute()
		assert.NoError(t, err)

		// Then each player is at the starting score
		assert.Equal(t, playerCount, len(ladder), "unexpected number of players in the ladder")
		for _, entry := range ladder {
			assert.Equal(t, startingScore, entry.Score)
		}
	})
}

type ComputeLadder struct {

}

type Ladder []struct {
	Score int
}

func (cl ComputeLadder) Execute() (Ladder, error) {
	return Ladder{}, nil
}

func NewComputeLadder() ComputeLadder {
	return ComputeLadder{}
}

type RegisterUser struct {

}

func (ru RegisterUser) Execute(userId string) error {
	return nil
}

func NewRegisterUser() RegisterUser {
	return RegisterUser{}
}
