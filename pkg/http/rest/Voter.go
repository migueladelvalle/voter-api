package rest

type Voter struct {
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	VoterHistory []VoterHistory `json:"voter_history"`
	Created      string         `json:"created"`
	Modified     string         `json:"modified"`
}
