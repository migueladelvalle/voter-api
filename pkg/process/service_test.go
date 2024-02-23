package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testService Service

func init() {
	testService = NewService(&MockRepository{})
}

func TestInvalidRequestFailuresCreateVoter(t *testing.T) {
	err := testService.CreateVoter(SampleVoterZeroId)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.CreateVoter(SampleVoterNegativeId)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.CreateVoter(SampleVoterNoName)
	assert.Equal(t, ErrInvalidName.Error(), err)

	err = testService.CreateVoter(SampleVoterInvalidEmail)
	assert.Equal(t, ErrInvalidEmail.Error(), err)
}

func TestInvalidRequestFailuresUpdateVoterInfo(t *testing.T) {
	err := testService.UpdateVoterInfo(SampleVoterZeroId)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.UpdateVoterInfo(SampleVoterNegativeId)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.UpdateVoterInfo(SampleVoterNoName)
	assert.Equal(t, ErrInvalidName.Error(), err)

	err = testService.UpdateVoterInfo(SampleVoterInvalidEmail)
	assert.Equal(t, ErrInvalidEmail.Error(), err)
}

func TestInvalidRequestFailuresDeleteSingleVoter(t *testing.T) {
	err := testService.DeleteSingleVoter(-1)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.DeleteSingleVoter(0)
	assert.Equal(t, ErrInvalidId.Error(), err)
}

func TestInvalidRequestFailuresCreateVoterHistory(t *testing.T) {
	err := testService.CreateVoterHistory(0, 1, SampleValidVoterHistory)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.CreateVoterHistory(-1, 1, SampleValidVoterHistory)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.CreateVoterHistory(1, 0, SampleValidVoterHistory)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.CreateVoterHistory(1, -1, SampleValidVoterHistory)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.CreateVoterHistory(1, 1, SampleVoterHistoryMissingDate)
	assert.Equal(t, ErrInvalidDate.Error(), err)
}

func TestInvalidRequestFailuresUpdateVoterHistoryInfo(t *testing.T) {
	err := testService.UpdateVoterHistoryInfo(0, 1, SampleValidVoterHistory)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.UpdateVoterHistoryInfo(-1, 1, SampleValidVoterHistory)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.UpdateVoterHistoryInfo(1, 0, SampleValidVoterHistory)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.UpdateVoterHistoryInfo(1, -1, SampleValidVoterHistory)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.UpdateVoterHistoryInfo(1, 1, SampleVoterHistoryMissingDate)
	assert.Equal(t, ErrInvalidDate.Error(), err)
}

func TestInvalidRequestFailuresDeleteSingleVoterPoll(t *testing.T) {
	err := testService.DeleteSingleVoterPoll(-1, 1)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.DeleteSingleVoterPoll(0, 1)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.DeleteSingleVoterPoll(1, 0)
	assert.Equal(t, ErrInvalidId.Error(), err)

	err = testService.DeleteSingleVoterPoll(1, -1)
	assert.Equal(t, ErrInvalidId.Error(), err)
}

func TestValidCreateVoter(t *testing.T) {
	err := testService.CreateVoter(SampleValidrequest)
	assert.NoError(t, err)
}

func TestValidUpdateVoterInfo(t *testing.T) {
	err := testService.UpdateVoterInfo(SampleValidrequest)
	assert.NoError(t, err)
}

func TestValidDeleteSingleVoter(t *testing.T) {
	err := testService.DeleteSingleVoter(1)
	assert.NoError(t, err)
}

func TestValidCreateVoterHistory(t *testing.T) {
	err := testService.CreateVoterHistory(1, 1, SampleValidVoterHistory)
	assert.NoError(t, err)
}

func TestValidUpdateVoterHistoryInfo(t *testing.T) {
	err := testService.UpdateVoterHistoryInfo(1, 1, SampleValidVoterHistory)
	assert.NoError(t, err)
}

func DeleteSingleVoterPoll(t *testing.T) {
	err := testService.DeleteSingleVoterPoll(1, 1)
	assert.NoError(t, err)
}
