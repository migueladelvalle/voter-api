package retrieve

import (
	"time"
)

type MockRepository struct{}

var refTime, _ = time.Parse(
	time.RFC3339,
	"2024-02-14T16:01:55Z")

var SampleVoterDTO = NewVoterDTO(
	1,
	"test",
	"123@abc.com",
	make(HistoryMap),
	refTime,
	refTime,
)

var SampleVoterHistoryDTO = NewVoterHistoryDTO(
	1,
	1,
	refTime,
	refTime,
	refTime,
)

func (m *MockRepository) GetAllVoters() ([]VoterDTO, error) {

	var voters []VoterDTO

	voters = append(voters, SampleVoterDTO)
	voters = append(voters, SampleVoterDTO)
	voters = append(voters, SampleVoterDTO)

	return voters, nil
}

func (m *MockRepository) GetSingleVoter(id int) (VoterDTO, error) {

	return SampleVoterDTO, nil
}

func (m *MockRepository) GetVoterHistory(id int) ([]VoterHistoryDTO, error) {

	var history []VoterHistoryDTO

	history = append(history, SampleVoterHistoryDTO)
	history = append(history, SampleVoterHistoryDTO)
	history = append(history, SampleVoterHistoryDTO)

	return history, nil

}

func (m *MockRepository) GetSingleEvent(voterId int, pollId int) (VoterHistoryDTO, error) {

	return SampleVoterHistoryDTO, nil
}
