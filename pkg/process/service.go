package process

import (
	"regexp"
	"strings"
	"time"
)

type Service interface {
	CreateVoter(voter VoterDTO) error
	UpdateVoterInfo(updatedVoter VoterDTO) error
	DeleteSingleVoter(id int) error
	CreateVoterHistory(voterId int, pollId int, history VoterHistoryDTO) error
	UpdateVoterHistoryInfo(voterId int, pollId int, history VoterHistoryDTO) error
	DeleteSingleVoterPoll(voterId int, pollId int) error
}

type Repository interface {
	CreateVoter(voter VoterDTO) error
	UpdateVoterInfo(voter VoterDTO) error
	DeleteSingleVoter(id int) error
	CreateVoterHistory(voterId int, pollId int, history VoterHistoryDTO) error
	UpdateVoterHistoryInfo(voterId int, pollId int, history VoterHistoryDTO) error
	DeleteSingleVoterPoll(voterId int, pollId int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateVoter(voter VoterDTO) error {

	err := s.validateVoter(voter)
	if err != nil {
		return err
	}

	err = s.r.CreateVoter(voter)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateVoterInfo(voter VoterDTO) error {

	err := s.validateVoter(voter)
	if err != nil {
		return err
	}

	err = s.r.UpdateVoterInfo(voter)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteSingleVoter(id int) error {

	if id < 1 {
		return ErrInvalidId.Error()
	}

	err := s.r.DeleteSingleVoter(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateVoterHistory(voterId int, pollId int, history VoterHistoryDTO) error {

	err := s.validateVoterHistory(voterId, pollId, history)
	if err != nil {
		return err
	}

	err = s.r.CreateVoterHistory(voterId, pollId, history)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateVoterHistoryInfo(voterId int, pollId int, history VoterHistoryDTO) error {

	err := s.validateVoterHistory(voterId, pollId, history)
	if err != nil {
		return err
	}

	err = s.r.UpdateVoterHistoryInfo(voterId, pollId, history)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteSingleVoterPoll(voterId int, pollId int) error {

	if voterId < 1 || pollId < 1 {
		return ErrInvalidId.Error()
	}

	return s.r.DeleteSingleVoterPoll(voterId, pollId)

}

func isValidEmail(email string) bool {
	// Regular expression pattern for basic email validation
	// This pattern is a simplified version and may not cover all edge cases
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression pattern
	regExp := regexp.MustCompile(pattern)

	// Check if the email matches the pattern
	return regExp.MatchString(email)
}

func isInvalidString(s string) bool {

	trimmed := strings.TrimSpace(s)

	return trimmed == ""
}

func (s *service) validateVoter(voter VoterDTO) error {
	if voter.id < 1 {
		return ErrInvalidId.Error()
	}
	if isInvalidString(voter.name) {
		return ErrInvalidName.Error()
	}
	if !isValidEmail(voter.email) {
		return ErrInvalidEmail.Error()
	}

	return nil
}

func (s *service) validateVoterHistory(voterId int, pollId int, history VoterHistoryDTO) error {
	if voterId < 1 {
		return ErrInvalidId.Error()
	}
	if pollId < 1 {
		return ErrInvalidId.Error()
	}

	if (history.voteDate == time.Time{}) {
		return ErrInvalidDate.Error()
	}

	return nil
}
