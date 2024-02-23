package retrieve

import "errors"

type RetrieveServiceError string

const (
	ErrInvalidId RetrieveServiceError = "Id must be a positive non-zero integer."
)

func (e RetrieveServiceError) Error() error {
	return errors.New(string(e))
}
