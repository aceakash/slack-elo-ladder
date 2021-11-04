package acceptance_test

import (
	"github.com/aceakash/slack-elo-ladder/internal/domain/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSavingAndLoadingLadderTournamentFromFile(t *testing.T) {
	origLadderTournament := models.NewLadderTournament(2000, 32)
	origLadderTournament.RegisterUser("akash")
	origLadderTournament.RegisterUser("yash")
	origLadderTournament.RegisterUser("reshma")
	origLadderTournament.RegisterUser("saanika")
	origLadderTournament.RegisterMatchResult("akash", "reshma")
	origLadderTournament.RegisterMatchResult("akash", "yash")
	origLadderTournament.RegisterMatchResult("akash", "saanika")
	origLadderTournament.RegisterMatchResult("yash", "reshma")
	origLadderTournament.RegisterMatchResult("yash", "saanika")
	origLadderTournament.RegisterMatchResult("reshma", "saanika")

	jsonFileRepo := NewLadderTournamentJsonFileRepo("./events.json")
	// todo temp file?
	// todo this test should work for any interface implementation

	err := jsonFileRepo.SaveLadderTournament(origLadderTournament)
	assert.NoError(t, err)

	loadedLadderTournament, err := jsonFileRepo.LoadLadderTournament()
	assert.NoError(t, err)

	assert.Equal(t, loadedLadderTournament, origLadderTournament)
}

type LadderTournamentRepo interface {
	LoadLadderTournament() (*models.LadderTournament, error)
	SaveLadderTournament(ladderTourney *models.LadderTournament) error
}

type LadderTournamentJsonFileRepo struct {
	filePath string
}

func (t LadderTournamentJsonFileRepo) SaveLadderTournament(ladderTourney *models.LadderTournament) error {
	json, err := ladderTourney.Json()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(t.filePath, json, 0o644)
	if err != nil {
		return errors.Wrapf(err, "error writing to file %s", t.filePath)
	}
	return nil
}

func NewLadderTournamentJsonFileRepo(filePath string) LadderTournamentJsonFileRepo {
	return LadderTournamentJsonFileRepo{filePath: filePath}
}

func (t LadderTournamentJsonFileRepo) LoadLadderTournament() (*models.LadderTournament, error) {
	bytes, err := ioutil.ReadFile(t.filePath)
	if err != nil {
		return nil, err
	}
	lt, err := models.NewLadderTournamentFromJson(bytes)
	if err != nil {
		return nil, err
	}
	return lt, nil
}


