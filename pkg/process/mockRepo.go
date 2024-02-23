package process

import (
	"time"

	fake "github.com/brianvoe/gofakeit/v6" //aliasing package name
)

type MockRepository struct{}

var SampleValidrequest = NewVoterDTO(
	fake.IntRange(1, 10),
	fake.Name(),
	fake.Email(),
)

var SampleVoterNegativeId = NewVoterDTO(
	-1,
	fake.Name(),
	fake.Email(),
)

var SampleVoterZeroId = NewVoterDTO(
	0,
	fake.Name(),
	fake.Email(),
)

var SampleVoterNoName = NewVoterDTO(
	fake.IntRange(1, 10),
	"",
	fake.Email(),
)

var SampleVoterInvalidEmail = NewVoterDTO(
	fake.IntRange(1, 10),
	fake.Name(),
	"badEmail",
)

var SampleVoterHistoryZeroValueId = NewVoterHistoryDTO(
	0,
	fake.IntRange(1, 10),
	fake.Date(),
)

var SampleVoterHistoryNegativeValueId = NewVoterHistoryDTO(
	-1,
	fake.IntRange(1, 10),
	fake.Date(),
)

var SampleVoterHistoryNegativePollValueId = NewVoterHistoryDTO(
	fake.IntRange(1, 10),
	-1,
	fake.Date(),
)

var SampleVoterHistoryZeroPollValueId = NewVoterHistoryDTO(
	fake.IntRange(1, 10),
	0,
	fake.Date(),
)

var SampleVoterHistoryMissingDate = NewVoterHistoryDTO(
	fake.IntRange(1, 10),
	fake.IntRange(1, 10),
	time.Time{},
)

var SampleValidVoterHistory = NewVoterHistoryDTO(
	fake.IntRange(1, 10),
	fake.IntRange(1, 10),
	fake.Date(),
)

func (m *MockRepository) CreateVoter(voter VoterDTO) error {
	return nil
}

func (m *MockRepository) UpdateVoterInfo(updatedVoter VoterDTO) error {
	return nil
}

func (m *MockRepository) DeleteSingleVoter(id int) error {
	return nil
}

func (m *MockRepository) CreateVoterHistory(voterId int, pollId int, history VoterHistoryDTO) error {
	return nil
}

func (m *MockRepository) UpdateVoterHistoryInfo(voterId int, pollId int, history VoterHistoryDTO) error {
	return nil
}

func (m *MockRepository) DeleteSingleVoterPoll(voterId int, pollId int) error {
	return nil
}
