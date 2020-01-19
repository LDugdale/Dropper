package services

import (
	"github.com/LDugdale/dropper/pkg/services/userService/abstractions"
	"github.com/ldugdale/dropper/pkg/commonAbstractions"
	"github.com/ldugdale/dropper/pkg/cryptography"
	"github.com/ldugdale/dropper/pkg/log"
)

type UserService struct {
	logger         log.Logger
	userRepository abstractions.UserRepository
	passwordHasher cryptography.PasswordHasher
}

func NewUserService(logger log.Logger, userRepository abstractions.UserRepository, passwordHasher *cryptography.PasswordHasher) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
		passwordHasher: *passwordHasher,
	}
}

func (us *UserService) SignUp(user *commonAbstractions.UserModel) (*commonAbstractions.User, error) {

	userPassword := user.Password
	hashedUserPassword, err := us.passwordHasher.HashAndSalt(userPassword)
	if err != nil {
		us.logger.LogError("Error occured whilst hashing password", err)

		return nil, abstractions.UnexpectedError(err)
	}
	user.Password = *hashedUserPassword

	rowsAffected, err := us.userRepository.CreateUser(user)
	if err != nil {
		us.logger.LogTrace("Error occured whilst calling CreateUser from repository", err)
		return nil, abstractions.UnexpectedError(err)
	}

	if rowsAffected < 1 {
		us.logger.LogTrace(abstractions.UserExists, user)

		return nil, abstractions.UserExistsError
	}

	signUpResult := &commonAbstractions.User{
		Username: user.Username,
	}

	return signUpResult, nil
}

func (us *UserService) SignIn(user *commonAbstractions.UserModel) (*commonAbstractions.User, error) {

	returnedUser, err := us.userRepository.GetUser(user.Username)
	if err != nil {
		us.logger.LogError("Error occured whilst getting user: ", err)
		return nil, abstractions.UserDoesNotExistWrapError(err)
	}

	isPasswordMatch := us.passwordHasher.ComparePasswords(returnedUser.Password, user.Password)
	if !isPasswordMatch {
		return nil, abstractions.PasswordIncorrectError
	}

	signInResult := &commonAbstractions.User{
		Username: user.Username,
	}

	return signInResult, nil
}
