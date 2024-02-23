package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"drexel.edu/voter-api/pkg/process"
	"drexel.edu/voter-api/pkg/retrieve"
)

type DbMap map[int]Voter

type VoterDB struct {
	voterList  DbMap
	dbFileName string
}

func NewJsonDB(dbFile string) (*VoterDB, error) {

	if _, err := os.Stat(dbFile); err != nil {
		//If the file doesn't exist, create it
		err := initDB(dbFile)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Open(dbFile)
	if err != nil {
		return &VoterDB{}, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if stat.Size() == 0 {
		err := initDB(dbFile)
		if err != nil {
			return nil, err
		}
	}

	voterList := &VoterDB{
		voterList:  make(map[int]Voter),
		dbFileName: dbFile,
	}

	return voterList, nil
}

func (v *VoterDB) RestoreDB(targetFileName string) error {

	dbFileName := v.dbFileName
	backupFileName := targetFileName

	dbFile, err := os.OpenFile(dbFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		msg := fmt.Sprintf("failed to open %s", dbFileName)
		return errors.New(msg)
	}

	defer dbFile.Close()

	backupFile, err := os.Open(backupFileName)
	if err != nil {
		msg := fmt.Sprintf("failed to open %s", backupFileName)
		return errors.New(msg)
	}

	defer backupFile.Close()

	buffer := make([]byte, 1024)

	for {
		bytesRead, err := backupFile.Read(buffer)
		if err != nil && err != io.EOF {
			msg := fmt.Sprintf("failed to read from %s", backupFileName)
			return errors.New(msg)
		}
		if bytesRead == 0 {
			fmt.Println("finished copying from ", backupFileName, " to ", dbFileName)
			break
		}
		if _, err := dbFile.Write(buffer[:bytesRead]); err != nil {
			fmt.Print("failed to write to ", dbFileName)
		}
	}

	return nil
}

func (v *VoterDB) CreateVoter(voter process.VoterDTO) error {
	if err := v.loadDB(); err != nil {
		return ErrFailedToLoadDB.Error()
	}

	if _, exists := v.voterList[voter.GetId()]; exists {
		return ErrVoterAlreadyExists.Error()
	}

	currentTime := time.Now()

	newVoter := Voter{
		Id:           voter.GetId(),
		Name:         voter.GetName(),
		Email:        voter.GetEmail(),
		VoterHistory: nil,
		Created:      currentTime,
		Modified:     currentTime,
	}

	v.voterList[voter.GetId()] = newVoter

	if err := v.saveDB(); err != nil {
		return ErrSaveFailed.Error()
	}

	fmt.Println("The voter was successfully registered.")

	v.PrintItem(newVoter)

	return nil
}

func (v *VoterDB) UpdateVoterInfo(voter process.VoterDTO) error {

	if err := v.loadDB(); err != nil {
		return ErrFailedToLoadDB.Error()
	}

	if previousVoter, exists := v.voterList[voter.GetId()]; exists {
		currentTime := time.Now()

		updatedVoter := Voter{
			Id:           voter.GetId(),
			Name:         voter.GetName(),
			Email:        voter.GetEmail(),
			VoterHistory: previousVoter.VoterHistory,
			Created:      previousVoter.Created,
			Modified:     currentTime,
		}

		v.voterList[voter.GetId()] = updatedVoter

		if err := v.saveDB(); err != nil {
			return ErrSaveFailed.Error()
		}

		fmt.Println("The voter was successfully registered.")

		v.PrintItem(updatedVoter)

		return nil
	}

	return ErrVoterNotFound.Error()
}

func (v *VoterDB) DeleteSingleVoter(id int) error {

	if err := v.loadDB(); err != nil {
		return ErrFailedToLoadDB.Error()
	}

	if _, exists := v.voterList[id]; exists {
		delete(v.voterList, id)

		if err := v.saveDB(); err != nil {
			return ErrSaveFailed.Error()
		}

		fmt.Println("The voter was successfully deleted.")

		return nil
	}

	return ErrVoterNotFound.Error()
}

func (v *VoterDB) CreateVoterHistory(voterId int, pollId int, history process.VoterHistoryDTO) error {
	if err := v.loadDB(); err != nil {
		return ErrFailedToLoadDB.Error()
	}

	voter, exists := v.voterList[voterId]
	if !exists {
		return ErrVoterNotFound.Error()
	}

	if _, exists := v.voterList[voterId].VoterHistory[pollId]; exists {
		return ErrHistoryAlreadyExists.Error()
	}

	currentTime := time.Now()

	if voter.VoterHistory == nil {
		voter.VoterHistory = make(HistoryMap)
	}

	voter.VoterHistory[pollId] = VoterHistory{
		PollId:   pollId,
		VoteId:   history.GetVoteID(),
		VoteDate: history.GetVoteDate(),
		Created:  currentTime,
		Modified: currentTime,
	}

	v.voterList[voterId] = voter

	if err := v.saveDB(); err != nil {
		return ErrSaveFailed.Error()
	}

	fmt.Println("The voter poll was successfully registered.")

	v.PrintItemHistory(v.voterList[voterId].VoterHistory[pollId])

	return nil
}

func (v *VoterDB) UpdateVoterHistoryInfo(voterId int, pollId int, history process.VoterHistoryDTO) error {

	if err := v.loadDB(); err != nil {
		return ErrFailedToLoadDB.Error()
	}

	if previousHistory, exists := v.voterList[voterId].VoterHistory[pollId]; exists {

		currentTime := time.Now()

		newHistory := VoterHistory{
			PollId:   pollId,
			VoteId:   history.GetVoteID(),
			VoteDate: history.GetVoteDate(),
			Created:  previousHistory.Created,
			Modified: currentTime,
		}

		v.voterList[voterId].VoterHistory[pollId] = newHistory

		if err := v.saveDB(); err != nil {
			return ErrSaveFailed.Error()
		}

		fmt.Println("The voter poll was successfully updated.")

		v.PrintItemHistory(newHistory)

		return nil
	}

	return ErrHistoryNotFound.Error()

}

func (v *VoterDB) DeleteSingleVoterPoll(voterId int, pollId int) error {
	if err := v.loadDB(); err != nil {
		return ErrFailedToLoadDB.Error()
	}

	if _, exists := v.voterList[voterId].VoterHistory[pollId]; exists {
		delete(v.voterList[voterId].VoterHistory, pollId)

		if err := v.saveDB(); err != nil {
			return ErrSaveFailed.Error()
		}

		fmt.Println("The voter poll was successfully deleted.")

		return nil
	}

	return ErrHistoryNotFound.Error()
}
func (v *VoterDB) GetAllVoters() ([]retrieve.VoterDTO, error) {

	if err := v.loadDB(); err != nil {
		return nil, ErrFailedToLoadDB.Error()
	}

	var votersList []retrieve.VoterDTO

	for _, voter := range v.voterList {

		voterDTO := retrieve.NewVoterDTO(
			voter.Id,
			voter.Name,
			voter.Email,
			v.copyVoterHistoryMap(voter.VoterHistory),
			voter.Created,
			voter.Modified,
		)

		votersList = append(votersList, voterDTO)
	}

	return votersList, nil
}

func (v *VoterDB) GetSingleVoter(id int) (retrieve.VoterDTO, error) {
	if err := v.loadDB(); err != nil {
		return retrieve.VoterDTO{}, ErrFailedToLoadDB.Error()
	}

	if voter, exists := v.voterList[id]; exists {
		return retrieve.NewVoterDTO(
			voter.Id,
			voter.Name,
			voter.Email,
			v.copyVoterHistoryMap(voter.VoterHistory),
			voter.Created,
			voter.Modified,
		), nil
	}

	return retrieve.VoterDTO{}, ErrVoterNotFound.Error()
}

func (v *VoterDB) GetVoterHistory(voterId int) ([]retrieve.VoterHistoryDTO, error) {
	if err := v.loadDB(); err != nil {
		return nil, ErrFailedToLoadDB.Error()
	}

	if _, exists := v.voterList[voterId]; exists {

		if v.voterList[voterId].VoterHistory == nil {
			return nil, ErrNoVoterHistory.Error()
		}

		var historyList []retrieve.VoterHistoryDTO

		for _, item := range v.voterList[voterId].VoterHistory {
			newHistory := retrieve.NewVoterHistoryDTO(
				item.PollId,
				item.VoteId,
				item.VoteDate,
				item.Created,
				item.Modified,
			)
			historyList = append(historyList, newHistory)
		}

		return historyList, nil
	}

	return nil, ErrVoterNotFound.Error()
}

func (v *VoterDB) GetSingleEvent(voterId int, pollId int) (retrieve.VoterHistoryDTO, error) {
	if err := v.loadDB(); err != nil {
		return retrieve.VoterHistoryDTO{}, ErrFailedToLoadDB.Error()
	}

	if _, exists := v.voterList[voterId]; !exists {
		return retrieve.VoterHistoryDTO{}, ErrVoterNotFound.Error()
	}

	if history, exists := v.voterList[voterId].VoterHistory[pollId]; exists {

		return retrieve.NewVoterHistoryDTO(
			history.PollId,
			history.VoteId,
			history.VoteDate,
			history.Created,
			history.Modified,
		), nil
	}

	return retrieve.VoterHistoryDTO{}, ErrHistoryNotFound.Error()
}

func (v *VoterDB) PrintItem(item Voter) {
	jsonBytes, _ := json.MarshalIndent(item, "", "  ")
	fmt.Println(string(jsonBytes))
}

func (v *VoterDB) PrintItemHistory(item VoterHistory) {
	jsonBytes, _ := json.MarshalIndent(item, "", "  ")
	fmt.Println(string(jsonBytes))
}

func (v *VoterDB) PrintAllItems(itemList []Voter) {
	for _, item := range itemList {
		v.PrintItem(item)
	}
}

func initDB(dbFileName string) error {
	f, err := os.Create(dbFileName)
	if err != nil {
		return err
	}

	_, err = f.Write([]byte("[]"))
	if err != nil {
		return err
	}

	f.Close()

	return nil
}

func (v *VoterDB) saveDB() error {

	var voterList []Voter
	for _, item := range v.voterList {
		voterList = append(voterList, item)
	}

	data, err := json.MarshalIndent(voterList, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(v.dbFileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (v *VoterDB) loadDB() error {
	data, err := os.ReadFile(v.dbFileName)
	if err != nil {
		return err
	}

	var voterList []Voter
	err = json.Unmarshal(data, &voterList)
	if err != nil {
		return err
	}

	for _, item := range voterList {
		v.voterList[item.Id] = item
	}

	return nil
}

func (v *VoterDB) copyVoterHistoryMap(history HistoryMap) retrieve.HistoryMap {
	returnMap := make(retrieve.HistoryMap)

	for _, item := range history {
		newHistory := retrieve.NewVoterHistoryDTO(
			item.PollId,
			item.PollId,
			item.VoteDate,
			item.Created,
			item.Modified,
		)

		returnMap[item.PollId] = newHistory
	}

	return returnMap
}
