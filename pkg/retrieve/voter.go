package retrieve

import (
	"time"
)

type HistoryMap map[int]VoterHistoryDTO

type VoterDTO struct {
	id       int
	name     string
	email    string
	history  HistoryMap
	created  time.Time
	modified time.Time
}

func NewVoterDTO(id int, name string, email string, history HistoryMap, created time.Time, modified time.Time) VoterDTO {
	return VoterDTO{
		id:       id,
		name:     name,
		email:    email,
		history:  history,
		created:  created,
		modified: modified,
	}
}

func (v *VoterDTO) GetBlankVoter() VoterDTO {
	return VoterDTO{}
}

func (v *VoterDTO) GetId() int {
	return v.id
}

func (v *VoterDTO) GetName() string {
	return v.name
}

func (v *VoterDTO) GetEmail() string {
	return v.email
}

func (v *VoterDTO) GetHistory() HistoryMap {
	return v.history
}

func (v *VoterDTO) GetCreated() time.Time {
	return v.created
}
func (v *VoterDTO) GetModified() time.Time {
	return v.modified
}
