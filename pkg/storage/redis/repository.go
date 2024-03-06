package cache

import (
	"errors"

	"drexel.edu/voter-api/pkg/process"
	"drexel.edu/voter-api/pkg/retrieve"
)

func (t *VoterCache) CreateVoter(voter process.VoterDTO) error {
	newVoter := t.processVoterDTOToStorageVoter(voter)

	if _, err := t.GetItem(newVoter.Id); err == nil {
		return errors.New("the specified Voter Id already exists")
	}

	currentTime := t.Clock.Now()
	newVoter.Created = currentTime
	newVoter.Modified = currentTime

	return t.AddItem(&newVoter)
}

func (t *VoterCache) UpdateVoterInfo(updatedVoter process.VoterDTO) error {
	newVoter := t.processVoterDTOToStorageVoter(updatedVoter)

	if _, err := t.GetItem(newVoter.Id); err != nil {
		return errors.New("the specified Voter Id was not found")
	}

	currentTime := t.Clock.Now()
	newVoter.Modified = currentTime

	return t.UpdateItem(&newVoter)
}

func (t *VoterCache) DeleteSingleVoter(id int) error {

	if !t.doesKeyExist(id) {
		return errors.New("the specified Voter Id was not found")
	}

	return t.DeleteItem(id)
}

func (t *VoterCache) CreateVoterHistory(voterId int, pollId int, history process.VoterHistoryDTO) error {

	targetVoter, err := t.GetItem(voterId)
	if err != nil {
		return errors.New("the specified Voter Id was not found")
	}

	if _, exists := targetVoter.VoterHistory[pollId]; exists {
		return errors.New("the specified Poll Id already exists")
	}

	if targetVoter.VoterHistory == nil {
		targetVoter.VoterHistory = make(HistoryMap)
	}

	newHistory := t.processHistoryDTOToStorageHistory(history)

	currentTime := t.Clock.Now()
	targetVoter.Modified = currentTime
	newHistory.Created = currentTime
	newHistory.Modified = currentTime

	targetVoter.VoterHistory[pollId] = newHistory

	return t.upsertVoter(targetVoter)
}

func (t *VoterCache) UpdateVoterHistoryInfo(voterId int, pollId int, history process.VoterHistoryDTO) error {

	targetVoter, err := t.GetItem(voterId)
	if err != nil {
		return errors.New("the specified Voter Id was not found")
	}

	if _, exists := targetVoter.VoterHistory[pollId]; !exists {
		return errors.New("the specified Poll Id could not be found")
	}

	newHistory := t.processHistoryDTOToStorageHistory(history)

	currentTime := t.Clock.Now()
	targetVoter.Modified = currentTime
	newHistory.Modified = currentTime

	targetVoter.VoterHistory[pollId] = newHistory

	return t.upsertVoter(targetVoter)
}

func (t *VoterCache) DeleteSingleVoterPoll(voterId int, pollId int) error {

	targetVoter, err := t.GetItem(voterId)
	if err != nil {
		return errors.New("the specified Voter Id was not found")
	}

	if _, exists := targetVoter.VoterHistory[pollId]; !exists {
		return errors.New("the specified Poll Id could not be found")
	}

	delete(targetVoter.VoterHistory, pollId)

	currentTime := t.Clock.Now()
	targetVoter.Modified = currentTime

	return t.upsertVoter(targetVoter)

}

func (t *VoterCache) GetAllVoters() ([]retrieve.VoterDTO, error) {
	allItems, err := t.GetAllItems()
	if err != nil {
		return nil, nil
	}

	var voters []retrieve.VoterDTO

	for _, item := range allItems {

		voters = append(voters, t.StorageVoterToRetrieveDTO(item))
	}

	if len(voters) < 1 {
		return nil, errors.New("no Voter Data was found")
	}

	return voters, nil
}

func (t *VoterCache) GetSingleVoter(id int) (retrieve.VoterDTO, error) {
	voter, err := t.GetItem(id)
	if err != nil {
		if err.Error() == "unexpected end of JSON input" {
			return retrieve.VoterDTO{}, errors.New("the specified Voter Id was not found")
		}
		return retrieve.VoterDTO{}, err
	}
	return t.StorageVoterToRetrieveDTO(*voter), nil
}

func (t *VoterCache) GetVoterHistory(id int) ([]retrieve.VoterHistoryDTO, error) {

	targetVoter, err := t.GetItem(id)
	if err != nil {
		return nil, errors.New("the specified Voter Id was not found")
	}

	var history []retrieve.VoterHistoryDTO

	for _, item := range targetVoter.VoterHistory {
		history = append(history, t.StoragehistoryToRetrieveDTO(item))
	}

	if len(history) < 1 {
		return nil, errors.New("no history was found for the specified voter")
	}

	return history, nil
}

func (t *VoterCache) GetSingleEvent(voterId int, pollId int) (retrieve.VoterHistoryDTO, error) {

	targetVoter, err := t.GetItem(voterId)
	if err != nil {
		return retrieve.SampleVoterDTO.GetBlankHistory(), errors.New("the specified Voter Id was not found")
	}

	if targetVoter.VoterHistory == nil {
		return retrieve.SampleVoterDTO.GetBlankHistory(), errors.New("the specified Poll Id could not be found")
	}

	if history, exists := targetVoter.VoterHistory[pollId]; exists {
		return t.StoragehistoryToRetrieveDTO(history), nil
	}

	return retrieve.SampleVoterDTO.GetBlankHistory(), errors.New("the specified Poll Id could not be found")
}

func (t *VoterCache) processVoterDTOToStorageVoter(voter process.VoterDTO) Voter {
	return Voter{
		Id:    voter.GetId(),
		Name:  voter.GetName(),
		Email: voter.GetEmail(),
	}
}

func (t *VoterCache) processHistoryDTOToStorageHistory(history process.VoterHistoryDTO) VoterHistory {
	return VoterHistory{
		PollId:   history.GetPollID(),
		VoteId:   history.GetVoteID(),
		VoteDate: history.GetVoteDate(),
	}
}

func (t *VoterCache) StorageVoterToRetrieveDTO(voter Voter) retrieve.VoterDTO {

	tHistory := make(retrieve.HistoryMap)

	for _, item := range voter.VoterHistory {
		tHistory[item.PollId] = retrieve.NewVoterHistoryDTO(
			item.PollId,
			item.VoteId,
			item.VoteDate,
			item.Created,
			item.Modified,
		)
	}

	return retrieve.NewVoterDTO(
		voter.Id,
		voter.Name,
		voter.Email,
		tHistory,
		voter.Created,
		voter.Modified,
	)
}

func (t *VoterCache) StoragehistoryToRetrieveDTO(history VoterHistory) retrieve.VoterHistoryDTO {

	historyDTO := retrieve.NewVoterHistoryDTO(
		history.PollId,
		history.VoteId,
		history.VoteDate,
		history.Created,
		history.Modified,
	)

	return historyDTO
}
