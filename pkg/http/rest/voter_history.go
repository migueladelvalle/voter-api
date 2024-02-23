package rest

type VoterHistory struct {
	PollId   int    `json:"poll_id"`
	VoteId   int    `json:"vote_id"`
	VoteDate string `json:"vote_date"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
}
