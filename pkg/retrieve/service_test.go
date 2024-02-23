package retrieve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testService Service

func init() {
	testService = NewService(&MockRepository{})
}

func TestGetAllVoters(t *testing.T) {
	voters, err := testService.GetAllVoters()
	assert.NoError(t, err)

	for _, item := range voters {
		assert.Equal(t, SampleVoterDTO, item)
	}
}

func TestGetSingleVoter(t *testing.T) {
	voter, err := testService.GetSingleVoter(SampleVoterDTO.id)
	assert.NoError(t, err)
	assert.Equal(t, SampleVoterDTO, voter)
}

func TestErrorOnZeroValueId(t *testing.T) {
	_, err := testService.GetSingleVoter(0)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)
}

func TestErrorOnNegativeValueId(t *testing.T) {
	_, err := testService.GetSingleVoter(-1)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)
}

func TestErrorOnZeroValueIdVoterHistory(t *testing.T) {
	_, err := testService.GetVoterHistory(0)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)
}

func TestErrorOnNegativeValueIdGetVoterHistory(t *testing.T) {
	_, err := testService.GetVoterHistory(-1)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)
}

func TestGetVoterHistory(t *testing.T) {
	history, err := testService.GetVoterHistory(SampleVoterDTO.id)
	assert.NoError(t, err)

	for _, item := range history {
		assert.Equal(t, SampleVoterHistoryDTO, item)
	}
}

func TestErroOnZeroValueIdGetVoterHistory(t *testing.T) {
	_, err := testService.GetVoterHistory(0)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)
}

func TestErrorOnNegativeValueGetVoterHistory(t *testing.T) {
	_, err := testService.GetVoterHistory(-1)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)
}

func TestGetSingleEvent(t *testing.T) {
	history, err := testService.GetSingleEvent(SampleVoterDTO.id, SampleVoterHistoryDTO.pollId)
	assert.NoError(t, err)
	assert.Equal(t, SampleVoterHistoryDTO, history)
}

func TestErrorOnZeroIdGetSingleEvent(t *testing.T) {
	_, err := testService.GetSingleEvent(0, SampleVoterHistoryDTO.pollId)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)

	_, err = testService.GetSingleEvent(SampleVoterDTO.id, 0)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)
}

func TestErrorOnNegativeIdGetSingleEvent(t *testing.T) {
	_, err := testService.GetSingleEvent(-1, SampleVoterHistoryDTO.pollId)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)

	_, err = testService.GetSingleEvent(SampleVoterDTO.id, -1)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidId.Error(), err)
}
