package abstractions

import (
	abstractions "github.com/ldugdale/dropper/pkg/commonAbstractions"
)

type UserService interface {
	SignUp(user *abstractions.UserModel) (*abstractions.User, error)
	SignIn(signInDetails *abstractions.UserModel) (*abstractions.User, error)
}

type UserRepository interface {
	CreateUser(signUpDetails *abstractions.UserModel) (int64, error)
	GetUser(username string) (*abstractions.UserModel, error)
}