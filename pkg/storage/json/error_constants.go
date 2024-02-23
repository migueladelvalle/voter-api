package json

import "errors"

type RepositoryError string

const (
	ErrFailedToLoadDB       RepositoryError = "Failed to load the database."
	ErrGettingVoter         RepositoryError = "Unhandled Exception Occured While attempting to retrieve a Voter."
	ErrVoterAlreadyExists   RepositoryError = "Attempted to create a voter but the id already exists."
	ErrVoterNotFound        RepositoryError = "The Voter Id was not found."
	ErrSaveFailed           RepositoryError = "Error saving to the database."
	ErrHistoryNotFound      RepositoryError = "The History Id for the Voter was not found"
	ErrHistoryAlreadyExists RepositoryError = "Attempted to create new history for the voter but the poll Id already exists"
	ErrNoVoterHistory       RepositoryError = "No history was found for the voter Id"
)

func (e RepositoryError) Error() error {
	return errors.New(string(e))
}
