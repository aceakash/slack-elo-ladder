package models

type LadderEntry struct {
	Player Player
	Score  int
}

type Ladder []LadderEntry
