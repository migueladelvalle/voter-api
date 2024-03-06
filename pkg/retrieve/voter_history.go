package retrieve

import (
	"time"
)

type VoterHistoryDTO struct {
	pollId   int
	voteId   int
	voteDate time.Time
	created  time.Time
	modified time.Time
}

func NewVoterHistoryDTO(id int, voteId int, voteDate time.Time, created time.Time, modified time.Time) VoterHistoryDTO {
	return VoterHistoryDTO{
		pollId:   id,
		voteId:   voteId,
		voteDate: voteDate,
		created:  created,
		modified: modified,
	}
}

func (v *VoterDTO) GetBlankHistory() VoterHistoryDTO {
	return VoterHistoryDTO{}
}

func (v *VoterHistoryDTO) GetPollID() int {
	return v.pollId
}

func (v *VoterHistoryDTO) GetVoteID() int {
	return v.voteId
}

func (v *VoterHistoryDTO) GetVoteDate() time.Time {
	return v.voteDate
}

func (v *VoterHistoryDTO) GetCreated() time.Time {
	return v.created
}

func (v *VoterHistoryDTO) GetModified() time.Time {
	return v.modified
}
