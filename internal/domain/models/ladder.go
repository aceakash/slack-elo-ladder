package models

type LadderEntry struct {
	PlayerId string
	Score    int
}

type Ladder []LadderEntry

func (l Ladder) Len() int {
	return len(l)
}

func (l Ladder) Less(i, j int) bool {
	return l[j].Score < l[i].Score
}

func (l Ladder) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
