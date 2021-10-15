package models

import (
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

type MatchPlayedEventDetails struct {
	Winner string
	Loser  string
}

type UserRegisteredEventDetails struct {
	UserId string
}
