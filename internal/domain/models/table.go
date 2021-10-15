package models

type TableEntry struct {
	PlayerId string
	Score    int
}

type Table []TableEntry
