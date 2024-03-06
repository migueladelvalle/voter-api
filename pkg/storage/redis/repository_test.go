package cache

import (
	"context"
	"errors"
	"testing"
	"time"

	"drexel.edu/voter-api/pkg/process"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

var mockDb redis.Client
var mock redismock.ClientMock
var mockRepository VoterCache

func init() {
	db, rmock := redismock.NewClientMock()

	mockDb = *db
	mock = rmock

	input := "2024-03-06T03:25:55.517739-05:00"

	timestamp, err := time.Parse(time.RFC3339Nano, input)
	if err != nil {
		panic("couldn't set timestamp in init")
	}

	mockRepository = VoterCache{
		cache: cache{
			&mockDb,
			context.TODO(),
		},
		Clock: &mockClock{
			now: timestamp,
		},
	}
}

func TestCreateVoterHappyPath(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	mock.ExpectDo(expectedArgsGet...).RedisNil()
	mock.ExpectExists("voter:1").RedisNil()

	newVoter := process.NewVoterDTO(1, "Sam", "samIAm@hotmail.com")

	newVoterJson := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":null,"created":"2024-03-06T03:25:55.517739-05:00","modified":"2024-03-06T03:25:55.517739-05:00"}`

	expectedArgsSet := []interface{}{"JSON.SET", "voter:1", ".", newVoterJson}
	mock.ExpectDo(expectedArgsSet...).SetVal("OK")

	err := mockRepository.CreateVoter(newVoter)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestCreateVoterAlreadyExists(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	existingVoterJson := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":null,"created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}`
	mock.ExpectDo(expectedArgsGet...).SetVal(existingVoterJson)

	newVoter := process.NewVoterDTO(1, "Sam", "samIAm@hotmail.com")

	err := mockRepository.CreateVoter(newVoter)
	assert.Error(t, err)
	assert.Equal(t, errors.New("the specified Voter Id already exists"), err)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestUpdateVoterHappyPath(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	existingVoterJson := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":null,"created":"0001-01-01T00:00:00Z","modified":"2024-03-06T03:25:55.517739-05:00"}`

	mock.ExpectDo(expectedArgsGet...).SetVal(existingVoterJson)
	mock.ExpectExists("voter:1").SetVal(int64(1))

	newVoter := process.NewVoterDTO(1, "Sam", "samIAm@hotmail.com")

	expectedArgsSet := []interface{}{"JSON.SET", "voter:1", ".", existingVoterJson}
	mock.ExpectDo(expectedArgsSet...).SetVal("OK")

	err := mockRepository.UpdateVoterInfo(newVoter)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestUpdateVoterNotFound(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	mock.ExpectDo(expectedArgsGet...).RedisNil()

	newVoter := process.NewVoterDTO(1, "Sam", "samIAm@hotmail.com")

	err := mockRepository.UpdateVoterInfo(newVoter)
	assert.Error(t, err)
	assert.Equal(t, errors.New("the specified Voter Id was not found"), err)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestDeleteVoterHappyPath(t *testing.T) {

	mock.ClearExpect()

	mock.ExpectExists("voter:1").SetVal(int64(1))
	mock.ExpectDel("voter:1").SetVal(int64(1))

	err := mockRepository.DeleteItem(1)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestDeleteVoterNotFound(t *testing.T) {

	mock.ClearExpect()

	mock.ExpectExists("voter:1").RedisNil()

	err := mockRepository.DeleteItem(1)
	assert.Error(t, err)
	assert.Equal(t, errors.New("Voter item with id 1 does not exist"), err)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestCreateHistoryHappyPath(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	existingVoterJson := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":null,"created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}`

	mock.ExpectDo(expectedArgsGet...).SetVal(existingVoterJson)

	input := "2024-03-06T03:25:55.517739-05:00"

	timestamp, err := time.Parse(time.RFC3339Nano, input)
	assert.NoError(t, err)

	newVoterHistory := process.NewVoterHistoryDTO(1, 1, timestamp)

	modifiedVoter := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":{"1":{"poll_id":1,"vote_id":1,"vote_date":"2024-03-06T03:25:55.517739-05:00","created":"2024-03-06T03:25:55.517739-05:00","modified":"2024-03-06T03:25:55.517739-05:00"}},"created":"0001-01-01T00:00:00Z","modified":"2024-03-06T03:25:55.517739-05:00"}`

	expectedArgsSet := []interface{}{"JSON.SET", "voter:1", ".", modifiedVoter}
	mock.ExpectDo(expectedArgsSet...).SetVal("OK")

	err = mockRepository.CreateVoterHistory(1, 1, newVoterHistory)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestCreateHistoryPollIdAlreadyExists(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	existingVoterJson := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":{"1":{"poll_id":1,"vote_id":1,"vote_date":"2024-03-06T03:25:55.517739-05:00","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}},"created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}`

	mock.ExpectDo(expectedArgsGet...).SetVal(existingVoterJson)

	input := "2024-03-06T03:25:55.517739-05:00"

	timestamp, err := time.Parse(time.RFC3339Nano, input)
	assert.NoError(t, err)

	newVoterHistory := process.NewVoterHistoryDTO(1, 1, timestamp)

	err = mockRepository.CreateVoterHistory(1, 1, newVoterHistory)
	assert.Error(t, err)
	assert.Equal(t, errors.New("the specified Poll Id already exists"), err)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestUpdateHistoryHappyPath(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	existingVoterJson := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":{"1":{"poll_id":1,"vote_id":1,"vote_date":"2024-03-06T03:25:55.517739-05:00","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}},"created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}`

	mock.ExpectDo(expectedArgsGet...).SetVal(existingVoterJson)

	input := "2024-03-06T03:25:55.517739-05:00"

	timestamp, err := time.Parse(time.RFC3339Nano, input)
	assert.NoError(t, err)

	newVoterHistory := process.NewVoterHistoryDTO(1, 2, timestamp)

	modifiedVoter := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":{"1":{"poll_id":1,"vote_id":2,"vote_date":"2024-03-06T03:25:55.517739-05:00","created":"0001-01-01T00:00:00Z","modified":"2024-03-06T03:25:55.517739-05:00"}},"created":"0001-01-01T00:00:00Z","modified":"2024-03-06T03:25:55.517739-05:00"}`

	expectedArgsSet := []interface{}{"JSON.SET", "voter:1", ".", modifiedVoter}
	mock.ExpectDo(expectedArgsSet...).SetVal("OK")

	err = mockRepository.UpdateVoterHistoryInfo(1, 1, newVoterHistory)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestUpdateHistoryPollNotFound(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	existingVoterJson := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":null,"created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}`

	mock.ExpectDo(expectedArgsGet...).SetVal(existingVoterJson)

	input := "2024-03-06T03:25:55.517739-05:00"

	timestamp, err := time.Parse(time.RFC3339Nano, input)
	assert.NoError(t, err)

	newVoterHistory := process.NewVoterHistoryDTO(1, 2, timestamp)

	err = mockRepository.UpdateVoterHistoryInfo(1, 1, newVoterHistory)
	assert.Error(t, err)
	assert.Equal(t, errors.New("the specified Poll Id could not be found"), err)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestDeleteHistoryHappyPath(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	existingVoterJson := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":{"1":{"poll_id":1,"vote_id":1,"vote_date":"2024-03-06T03:25:55.517739-05:00","created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}},"created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}`

	mock.ExpectDo(expectedArgsGet...).SetVal(existingVoterJson)

	modifiedVoter := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":{},"created":"0001-01-01T00:00:00Z","modified":"2024-03-06T03:25:55.517739-05:00"}`

	expectedArgsSet := []interface{}{"JSON.SET", "voter:1", ".", modifiedVoter}
	mock.ExpectDo(expectedArgsSet...).SetVal("OK")

	err := mockRepository.DeleteSingleVoterPoll(1, 1)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestDeleteHistoryNotFound(t *testing.T) {

	mock.ClearExpect()

	expectedArgsGet := []interface{}{"JSON.GET", "voter:1", "."}

	existingVoterJson := `{"id":1,"name":"Sam","email":"samIAm@hotmail.com","history":{},"created":"0001-01-01T00:00:00Z","modified":"0001-01-01T00:00:00Z"}`

	mock.ExpectDo(expectedArgsGet...).SetVal(existingVoterJson)

	err := mockRepository.DeleteSingleVoterPoll(1, 1)
	assert.Error(t, err)
	assert.Equal(t, errors.New("the specified Poll Id could not be found"), err)
	assert.NoError(t, mock.ExpectationsWereMet())

}
