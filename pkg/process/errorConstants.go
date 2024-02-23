package process

import "errors"

type processServiceError string

const (
	ErrInvalidId    processServiceError = "id must be a positive non-zero integer."
	ErrInvalidName  processServiceError = "name must not be blank"
	ErrInvalidEmail processServiceError = "email must be in the format of <adddress>@<domain> "
	ErrInvalidDate  processServiceError = "date must not be nil"
)

func (e processServiceError) Error() error {
	return errors.New(string(e))
}
