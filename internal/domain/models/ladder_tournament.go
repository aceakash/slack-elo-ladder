package models

import (
	"encoding/json"
	"github.com/aceakash/slack-elo-ladder/internal/domain"
)

type LadderTournament struct {
	startingScore  int
	constantFactor int
	events         []Event
}

func NewLadderTournament(startingScore int, constantFactor int) *LadderTournament {
	return &LadderTournament{
		startingScore:  startingScore,
		constantFactor: constantFactor,
		events:         []Event{},
	}
}

func (l *LadderTournament) RegisterUser(userId string) {
	l.events = append(l.events, NewUserRegisteredEvent(userId))
}

func (l *LadderTournament) ComputeLadder() (Ladder, error) {
	ladder := Ladder{}

	scores := map[string]int{}

	for _, event := range l.events {
		if event.Type == UserRegisteredEventType {
			eventDetails := event.Details.(UserRegisteredEventDetails)
			scores[eventDetails.UserId] = l.startingScore
		}
		if event.Type == MatchPlayedEventType {
			eventDetails := event.Details.(MatchPlayedEventDetails)
			winnerOldScore, loserOldScore := scores[eventDetails.Winner], scores[eventDetails.Winner]
			winnerNewScore, loserNewScore := domain.CalculateEloRating(winnerOldScore, loserOldScore, l.constantFactor)
			scores[eventDetails.Winner] = winnerNewScore
			scores[eventDetails.Loser] = loserNewScore
		}
	}
	for user, score := range scores {
		ladder = append(ladder, LadderEntry{
			PlayerId: user,
			Score:    score,
		})
	}
	ladder.SortByScoreDesc()
	return ladder, nil
}

func (lt *LadderTournament) RegisterMatchResult(winner string, loser string) {
	lt.events = append(lt.events, NewMatchPlayedEvent(winner, loser))
}

func (l *LadderTournament) Json() ([]byte, error) {
	return json.Marshal(l.Dto())
}

func (l *LadderTournament) Dto() LadderTournamentDTO {
	return LadderTournamentDTO{
		StartingScore:  l.startingScore,
		ConstantFactor: l.constantFactor,
		Events:         l.events[:],
	}
}

func NewLadderTournamentFromDto(dto LadderTournamentDTO) *LadderTournament {
	return &LadderTournament{
		startingScore:  dto.StartingScore,
		constantFactor: dto.ConstantFactor,
		events:         dto.Events[:],
	}
}

func NewLadderTournamentFromJson(jsonBytes []byte) (*LadderTournament, error) {
	var ltDto LadderTournamentDTO
	err := json.Unmarshal(jsonBytes, &ltDto)
	if err != nil {
		return nil, err
	}
	return NewLadderTournamentFromDto(ltDto), nil
}

type MatchPlayedEventDetails struct {
	Winner string
	Loser  string
}

type UserRegisteredEventDetails struct {
	UserId string
}

type LadderTournamentDTO struct {
	StartingScore  int     `json:"starting_score"`
	ConstantFactor int     `json:"constant_factor"`
	Events         []Event `json:"events"`
}


