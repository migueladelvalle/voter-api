package process

import "time"

type VoterHistoryDTO struct {
	pollId   int
	voteId   int
	voteDate time.Time
}

func NewVoterHistoryDTO(id int, voteId int, voteDate time.Time) VoterHistoryDTO {
	return VoterHistoryDTO{
		pollId:   id,
		voteId:   voteId,
		voteDate: voteDate,
	}
}

func (v *VoterHistoryDTO) GetPollID() int {
	return v.pollId
}

func (v *VoterHistoryDTO) GetVoteID() int {
	return v.pollId
}

func (v *VoterHistoryDTO) GetVoteDate() time.Time {
	return v.voteDate
}
