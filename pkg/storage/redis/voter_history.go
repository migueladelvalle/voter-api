package cache

import (
	"time"
)

type VoterHistory struct {
	PollId   int       `json:"poll_id"`
	VoteId   int       `json:"vote_id"`
	VoteDate time.Time `json:"vote_date"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}
