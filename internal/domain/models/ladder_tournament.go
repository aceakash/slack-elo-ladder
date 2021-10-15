package models

import (
	"fmt"
	"github.com/google/uuid"
	"sort"
	"time"
)

type Event struct {
	Type      string
	ID        string
	Timestamp time.Time
	Details   interface{}
}

type LadderTournament struct {
	startingScore int
	events        []Event
}

func NewLadderTournament(startingScore int) *LadderTournament {
	return &LadderTournament{
		startingScore: startingScore,
		events:        []Event{},
	}
}

func (l *LadderTournament) RegisterUser(userId string) {
	l.events = append(l.events, Event{
		Type:      "UserRegistered",
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Details: UserRegisteredEventDetails{
			UserId: userId,
		},
	})
}

func (l *LadderTournament) ComputeLadder() (Ladder, error) {
	ladder := Ladder{}

	scores := map[string]int{}

	for _, event := range l.events {
		if event.Type == "UserRegistered" {
			eventDetails := event.Details.(UserRegisteredEventDetails)
			scores[eventDetails.UserId] = l.startingScore
		}
		if event.Type == "MatchPlayed" {
			eventDetails := event.Details.(MatchPlayedEventDetails)
			winnerOldScore, loserOldScore := scores[eventDetails.Winner], scores[eventDetails.Winner]
			winnerNewScore, loserNewScore := CalculateEloScores(l.startingScore, 32, winnerOldScore, loserOldScore)
			fmt.Println(winnerNewScore, loserNewScore)
			scores[eventDetails.Winner] = winnerNewScore
			scores[eventDetails.Loser] = loserNewScore
		}
	}
	fmt.Println(scores)
	for user, score := range scores {
		ladder = append(ladder, LadderEntry{
			PlayerId: user,
			Score:    score,
		})
	}
	sort.Stable(ladder)
	return ladder, nil
}

func CalculateEloScores(startingScore int, constantFactor int, winnerOldScore int, loserOldScore int) (winnerNewScore int, loserNewScore int) {
	return 2016, 1984
}

func (lt *LadderTournament) RegisterMatchResult(winner string, loser string) {
	lt.events = append(lt.events, Event{
		Type:      "MatchPlayed",
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Details: MatchPlayedEventDetails{
			Winner: winner,
			Loser:  loser,
		},
	})
}

type MatchPlayedEventDetails struct {
	Winner string
	Loser  string
}

type UserRegisteredEventDetails struct {
	UserId string
}
