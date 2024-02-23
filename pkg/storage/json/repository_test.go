package json

import (
	"fmt"
	"os"
	"testing"

	"drexel.edu/voter-api/pkg/process"
	"drexel.edu/voter-api/pkg/retrieve"
	fake "github.com/brianvoe/gofakeit/v6" //aliasing package name
	"github.com/stretchr/testify/assert"
)

var db VoterDB

func init() {
	Refresh()
}

func TestMain(m *testing.M) {
	exitCode := m.Run()

	//clean test files
	os.Remove("./tmp_test")
	os.Remove("./tmp_test2")

	os.Exit(exitCode)
}

func Refresh() {
	os.Remove("./tmp_test")
	testDB, err := NewJsonDB("./tmp_test")
	if err != nil {
		fmt.Print("ERROR CREATING DB:", err)
		os.Exit(1)
	}

	db = *testDB
}

func TestCreateNewJsonDB(t *testing.T) {
	_, err := NewJsonDB("./tmp_test")
	assert.NoError(t, err)
}

func TestRestore(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.NoError(t, err)

	fmt.Println(fmt.Printf("Testing in directory: %s", currentDir))

	filePath := "./tmp_test2"
	backUpFile := "../../../Data.Bak"

	//Should overwrite with a blank file
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)

	assert.NoError(t, err, "Couldn't create a blank db file")
	defer file.Close()

	dbTemp, err := NewJsonDB(filePath)
	assert.NoError(t, err)

	err = dbTemp.RestoreDB(backUpFile)
	assert.NoError(t, err, "Error while restoring database")

	areFilesEqual, err := areFilesEqual(t, filePath, backUpFile)
	assert.NoError(t, err, "Error occurred comparing files")
	assert.Equal(t, true, areFilesEqual)

	os.Remove("./tmp_test2")
}

func areFilesEqual(t *testing.T, file1, file2 string) (bool, error) {
	contentFromFile1, err := os.Open(file1)
	assert.NoError(t, err, "Could not read from file1")

	contentFromFile2, err := os.Open(file2)
	assert.NoError(t, err, "Could not read from file2")

	bufferFile1 := make([]byte, 1024)
	bufferFile2 := make([]byte, 1024)

	for {
		bytesReadFromFile1, err := contentFromFile1.Read(bufferFile1)
		if err != nil && err.Error() != "EOF" {
			return false, err
		}

		bytesReadFromFile2, err := contentFromFile2.Read(bufferFile2)
		if err != nil && err.Error() != "EOF" {
			return false, err
		}

		if bytesReadFromFile1 != bytesReadFromFile2 {
			return false, nil

		}

		if bytesReadFromFile1 == 0 {
			break
		}

		if len(bufferFile1) != len(bufferFile2) {
			return false, nil
		}

		for i := range bufferFile1 {
			if bufferFile1[i] != bufferFile2[i] {
				return false, nil
			}
		}
	}
	return true, nil
}

func TestCreateRetrieveVoterWithRandomData(t *testing.T) {

	expectedVoter := process.NewVoterDTO(
		fake.IntRange(1, 10),
		fake.Name(),
		fake.Email(),
	)

	err := db.CreateVoter(expectedVoter)
	assert.NoError(t, err)

	actualVoter, err := db.GetSingleVoter(expectedVoter.GetId())
	assert.NoError(t, err)
	assert.Equal(t, expectedVoter.GetId(), actualVoter.GetId())
	assert.Equal(t, expectedVoter.GetName(), actualVoter.GetName())
	assert.Equal(t, expectedVoter.GetEmail(), actualVoter.GetEmail())
}

func TestUpdateRetrieveVoterWithRandomData(t *testing.T) {
	expectedVoter := process.NewVoterDTO(
		fake.IntRange(11, 20),
		fake.Name(),
		fake.Email(),
	)

	err := db.CreateVoter(expectedVoter)
	assert.NoError(t, err)

	expectedVoter = process.NewVoterDTO(
		expectedVoter.GetId(),
		fake.Name(),
		fake.Email(),
	)

	err = db.UpdateVoterInfo(expectedVoter)
	assert.NoError(t, err)

	actualVoter, err := db.GetSingleVoter(expectedVoter.GetId())
	assert.NoError(t, err)
	assert.Equal(t, expectedVoter.GetId(), actualVoter.GetId())
	assert.Equal(t, expectedVoter.GetName(), actualVoter.GetName())
	assert.Equal(t, expectedVoter.GetEmail(), actualVoter.GetEmail())
}

func TestDeleteVoter(t *testing.T) {
	expectedVoter := process.NewVoterDTO(
		fake.IntRange(21, 30),
		fake.Name(),
		fake.Email(),
	)

	err := db.CreateVoter(expectedVoter)
	assert.NoError(t, err)

	actualVoter, err := db.GetSingleVoter(expectedVoter.GetId())
	assert.NoError(t, err)
	assert.Equal(t, expectedVoter.GetId(), actualVoter.GetId())
	assert.Equal(t, expectedVoter.GetName(), actualVoter.GetName())
	assert.Equal(t, expectedVoter.GetEmail(), actualVoter.GetEmail())

	err = db.DeleteSingleVoter(expectedVoter.GetId())
	assert.NoError(t, err)

	nullVoter := retrieve.VoterDTO{}

	actualVoter, err = db.GetSingleVoter(expectedVoter.GetId())
	assert.Error(t, err)
	assert.Equal(t, nullVoter, actualVoter)
	assert.Equal(t, ErrVoterNotFound.Error(), err)
}

func TestCreateRetrievePollWithRandomData(t *testing.T) {
	expectedVoter := process.NewVoterDTO(
		fake.IntRange(31, 40),
		fake.Name(),
		fake.Email(),
	)

	err := db.CreateVoter(expectedVoter)
	assert.NoError(t, err)

	expectedPoll := process.NewVoterHistoryDTO(
		fake.IntRange(41, 50),
		fake.IntRange(51, 60),
		fake.Date(),
	)

	err = db.CreateVoterHistory(
		expectedVoter.GetId(),
		expectedPoll.GetPollID(),
		expectedPoll)
	assert.NoError(t, err)

	actualVoter, err := db.GetSingleVoter(expectedVoter.GetId())
	assert.NoError(t, err)
	assert.Equal(t, expectedVoter.GetId(), actualVoter.GetId())
	assert.Equal(t, expectedVoter.GetName(), actualVoter.GetName())
	assert.Equal(t, expectedVoter.GetEmail(), actualVoter.GetEmail())

	actualPoll, err := db.GetSingleEvent(actualVoter.GetId(), expectedPoll.GetPollID())
	assert.NoError(t, err)

	assert.Equal(t, expectedPoll.GetPollID(), actualPoll.GetPollID())
	assert.Equal(t, expectedPoll.GetVoteID(), actualPoll.GetVoteID())
	assert.Equal(t, expectedPoll.GetVoteDate(), actualPoll.GetVoteDate())
}

func TestUpdateRetrievePollWithRandomData(t *testing.T) {
	expectedVoter := process.NewVoterDTO(
		fake.IntRange(61, 70),
		fake.Name(),
		fake.Email(),
	)

	err := db.CreateVoter(expectedVoter)
	assert.NoError(t, err)

	expectedPoll := process.NewVoterHistoryDTO(
		fake.IntRange(71, 80),
		fake.IntRange(81, 90),
		fake.Date(),
	)

	err = db.CreateVoterHistory(
		expectedVoter.GetId(),
		expectedPoll.GetPollID(),
		expectedPoll)
	assert.NoError(t, err)

	expectedPoll = process.NewVoterHistoryDTO(
		expectedPoll.GetPollID(),
		fake.IntRange(91, 100),
		fake.Date(),
	)

	err = db.UpdateVoterHistoryInfo(
		expectedVoter.GetId(),
		expectedPoll.GetPollID(),
		expectedPoll)
	assert.NoError(t, err)

	actualVoter, err := db.GetSingleVoter(expectedVoter.GetId())
	assert.NoError(t, err)
	assert.Equal(t, expectedVoter.GetId(), actualVoter.GetId())
	assert.Equal(t, expectedVoter.GetName(), actualVoter.GetName())
	assert.Equal(t, expectedVoter.GetEmail(), actualVoter.GetEmail())

	actualPoll, err := db.GetSingleEvent(actualVoter.GetId(), expectedPoll.GetPollID())
	assert.NoError(t, err)

	assert.Equal(t, expectedPoll.GetPollID(), actualPoll.GetPollID())
	assert.Equal(t, expectedPoll.GetVoteID(), actualPoll.GetVoteID())
	assert.Equal(t, expectedPoll.GetVoteDate(), actualPoll.GetVoteDate())
}

func TestDeletePoll(t *testing.T) {
	expectedVoter := process.NewVoterDTO(
		fake.IntRange(101, 110),
		fake.Name(),
		fake.Email(),
	)

	err := db.CreateVoter(expectedVoter)
	assert.NoError(t, err)

	expectedPoll := process.NewVoterHistoryDTO(
		fake.IntRange(111, 120),
		fake.IntRange(121, 130),
		fake.Date(),
	)

	err = db.CreateVoterHistory(
		expectedVoter.GetId(),
		expectedPoll.GetPollID(),
		expectedPoll)
	assert.NoError(t, err)

	actualVoter, err := db.GetSingleVoter(expectedVoter.GetId())
	assert.NoError(t, err)
	assert.Equal(t, expectedVoter.GetId(), actualVoter.GetId())
	assert.Equal(t, expectedVoter.GetName(), actualVoter.GetName())
	assert.Equal(t, expectedVoter.GetEmail(), actualVoter.GetEmail())

	actualPoll, err := db.GetSingleEvent(actualVoter.GetId(), expectedPoll.GetPollID())
	assert.NoError(t, err)

	assert.Equal(t, expectedPoll.GetPollID(), actualPoll.GetPollID())
	assert.Equal(t, expectedPoll.GetVoteID(), actualPoll.GetVoteID())
	assert.Equal(t, expectedPoll.GetVoteDate(), actualPoll.GetVoteDate())

	err = db.DeleteSingleVoterPoll(
		expectedVoter.GetId(),
		expectedPoll.GetPollID(),
	)
	assert.NoError(t, err)

	actualPoll, err = db.GetSingleEvent(actualVoter.GetId(), expectedPoll.GetPollID())
	assert.Error(t, err)
	assert.Equal(t, ErrHistoryNotFound.Error(), err)
}

func TestGetAllVoterHistory(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.NoError(t, err)

	fmt.Println(fmt.Printf("Testing in directory: %s", currentDir))

	filePath := "./tmp_test2"

	//Should overwrite with a blank file
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)

	assert.NoError(t, err, "Couldn't create a blank db file")
	defer file.Close()

	dbTemp, err := NewJsonDB(filePath)
	assert.NoError(t, err)

	expectedVoter := process.NewVoterDTO(
		160,
		fake.Name(),
		fake.Email(),
	)
	err = dbTemp.CreateVoter(expectedVoter)
	assert.NoError(t, err)

	var history [3]process.VoterHistoryDTO

	iterator := 1

	for _, item := range history {
		item = process.NewVoterHistoryDTO(
			iterator,
			iterator,
			fake.Date(),
		)
		history[iterator-1] = item
		err = dbTemp.CreateVoterHistory(
			expectedVoter.GetId(),
			item.GetPollID(),
			item)
		assert.NoError(t, err)
		iterator++
	}

	actualVoter, err := dbTemp.GetSingleVoter(expectedVoter.GetId())
	assert.NoError(t, err)

	assert.Equal(t, 3, len(actualVoter.GetHistory()))

	for _, item := range history {
		actualPoll, err := dbTemp.GetSingleEvent(expectedVoter.GetId(), item.GetPollID())
		assert.NoError(t, err)
		assert.Equal(t, item.GetPollID(), actualPoll.GetPollID())
		assert.Equal(t, item.GetVoteID(), actualPoll.GetVoteID())
		assert.Equal(t, item.GetVoteDate(), actualPoll.GetVoteDate())
	}

	os.Remove("./tmp_test2")
}

func TestGetAllVoters(t *testing.T) {

	filePath := "./tmp_test3"

	os.Remove(filePath)

	currentDir, err := os.Getwd()
	assert.NoError(t, err)

	fmt.Println(fmt.Printf("Testing in directory: %s", currentDir))

	//Should overwrite with a blank file
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)

	assert.NoError(t, err, "Couldn't create a blank db file")
	defer file.Close()

	dbTemp, err := NewJsonDB(filePath)
	assert.NoError(t, err)

	var expectedVoters [3]process.VoterDTO

	item1 := process.NewVoterDTO(
		1,
		fake.Name(),
		fake.Email(),
	)
	expectedVoters[0] = item1
	err = dbTemp.CreateVoter(item1)
	assert.NoError(t, err)

	item2 := process.NewVoterDTO(
		2,
		fake.Name(),
		fake.Email(),
	)
	expectedVoters[1] = item2
	err = dbTemp.CreateVoter(item2)
	assert.NoError(t, err)

	item3 := process.NewVoterDTO(
		3,
		fake.Name(),
		fake.Email(),
	)
	expectedVoters[2] = item3
	err = dbTemp.CreateVoter(item3)
	assert.NoError(t, err)

	actualVoters, err := dbTemp.GetAllVoters()
	assert.NoError(t, err)

	assert.Equal(t, 3, len(actualVoters))

	assert.Equal(
		t,
		expectedVoters[0].GetId(),
		item1.GetId(),
	)

	assert.Equal(
		t,
		expectedVoters[0].GetEmail(),
		item1.GetEmail(),
	)

	assert.Equal(
		t,
		expectedVoters[0].GetName(),
		item1.GetName(),
	)

	assert.Equal(
		t,
		expectedVoters[1].GetId(),
		item2.GetId(),
	)

	assert.Equal(
		t,
		expectedVoters[1].GetEmail(),
		item2.GetEmail(),
	)

	assert.Equal(
		t,
		expectedVoters[1].GetName(),
		item2.GetName(),
	)

	assert.Equal(
		t,
		expectedVoters[2].GetId(),
		item3.GetId(),
	)

	assert.Equal(
		t,
		expectedVoters[2].GetEmail(),
		item3.GetEmail(),
	)

	assert.Equal(
		t,
		expectedVoters[2].GetName(),
		item3.GetName(),
	)

	os.Remove(filePath)
}
