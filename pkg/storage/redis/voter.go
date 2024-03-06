package cache

import (
	"time"
)

type HistoryMap map[int]VoterHistory

type Voter struct {
	Id           int        `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	VoterHistory HistoryMap `json:"history"`
	Created      time.Time  `json:"created"`
	Modified     time.Time  `json:"modified"`
}
