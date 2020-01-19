package abstractions

import (
	"github.com/ldugdale/dropper/pkg/customErrors"
)

const (
	UserExists        = "This user already exists."
	UserDoesNotExist  = "Username does not exist."
	PasswordIncorrect = "Password is incorrect."
	Unexpected        = "An unexpected error occured."
)

var (
	UserExistsError           = customErrors.AlreadyExists.NewWithDetails("User Exists", UserExists)
	UserDoesNotExistError     = customErrors.AlreadyExists.NewWithDetails("User Does Not Exist", UserDoesNotExist)
	PasswordIncorrectError    = customErrors.AlreadyExists.NewWithDetails("Password Incorrect", PasswordIncorrect)
	UnexpectedError           = func(err error) error { return customErrors.Unknown.WrapWithDetails(err, "Unexpected", Unexpected) }
	UserDoesNotExistWrapError = func(err error) error {
		return customErrors.AlreadyExists.WrapWithDetails(err, "User Does Not Exist", UserDoesNotExist)
	}
)
