package retrieve

type Service interface {
	GetAllVoters() ([]VoterDTO, error)
	GetSingleVoter(id int) (VoterDTO, error)
	GetVoterHistory(id int) ([]VoterHistoryDTO, error)
	GetSingleEvent(voterId int, pollId int) (VoterHistoryDTO, error)
}

type Repository interface {
	GetAllVoters() ([]VoterDTO, error)
	GetSingleVoter(id int) (VoterDTO, error)
	GetVoterHistory(id int) ([]VoterHistoryDTO, error)
	GetSingleEvent(voterId int, pollId int) (VoterHistoryDTO, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAllVoters() ([]VoterDTO, error) {

	voters, err := s.r.GetAllVoters()
	if err != nil {
		return nil, err
	}

	return voters, nil
}

func (s *service) GetSingleVoter(id int) (VoterDTO, error) {

	if id < 1 {
		return VoterDTO{}, ErrInvalidId.Error()
	}

	voter, err := s.r.GetSingleVoter(id)
	if err != nil {
		return VoterDTO{}, err
	}

	return voter, nil
}

func (s *service) GetVoterHistory(id int) ([]VoterHistoryDTO, error) {

	if id < 1 {
		return nil, ErrInvalidId.Error()
	}

	history, err := s.r.GetVoterHistory(id)
	if err != nil {
		return nil, err
	}

	return history, nil
}

func (s *service) GetSingleEvent(voterId int, pollId int) (VoterHistoryDTO, error) {

	if voterId < 1 || pollId < 1 {
		return VoterHistoryDTO{}, ErrInvalidId.Error()
	}

	history, err := s.r.GetSingleEvent(voterId, pollId)
	if err != nil {
		return VoterHistoryDTO{}, err
	}

	return history, nil
}
