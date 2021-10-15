package models

type LadderEntry struct {
	PlayerId string
	Score    int
}

type Ladder []LadderEntry
