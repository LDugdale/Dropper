package services

import (
	"errors"

	"github.com/ldugdale/dropper/pkg/commonAbstractions"
	"github.com/LDugdale/Dropper/pkg/services/userService/abstractions"
	"github.com/ldugdale/dropper/pkg/log"
	"github.com/ldugdale/dropper/pkg/cryptography"
)

type UserService struct {
	logger log.Logger
	userRepository abstractions.UserRepository
	passwordHasher cryptography.PasswordHasher
}

func NewUserService(logger log.Logger, userRepository abstractions.UserRepository, passwordHasher *cryptography.PasswordHasher) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger: logger,
		passwordHasher: *passwordHasher,
	}
}

func (us *UserService) SignUp(user *commonAbstractions.UserModel) (*commonAbstractions.User, error) {

	us.logger.LogTrace("SignUp", user)


	userPassword := user.Password
	hashedUserPassword, err := us.passwordHasher.HashAndSalt(userPassword)
	if err != nil {
		return nil, err
	}
	user.Password = *hashedUserPassword

	rowsAffected, err := us.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	if rowsAffected > 0 {
		
	}

	signUpResult := &commonAbstractions.User{
		Username: user.Username,
	}

	return signUpResult, nil
}

func (us *UserService) SignIn(user *commonAbstractions.UserModel) (*commonAbstractions.User, error) {

	returnedUser, err := us.userRepository.GetUser(user.Password)
	if err != nil {
		return nil, err
	}
	
	hashedUserPassword, err := us.passwordHasher.HashAndSalt(user.Password)
	if err != nil {
		return nil, err
	}

	isPasswordMatch := us.passwordHasher.ComparePasswords(*hashedUserPassword, returnedUser.Password)
	if !isPasswordMatch {
		return nil, errors.New("")
	}

	signInResult := &commonAbstractions.User{
		Username: user.Username,
	}

	return signInResult, nil
}