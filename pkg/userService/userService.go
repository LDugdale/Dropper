package userService

import (
	"golang.org/x/net/context"
	"github.com/LDugdale/Dropper/pkg/types"
	"github.com/ldugdale/dropper/pkg/logger"
	pb "github.com/LDugdale/Dropper/proto"
)

type IUserRepository interface {
	CreateUser(signUpDetails *types.UserModel) (int64, error)
	GetUser(username string) (*types.User, error)
}

type UserService struct {
	logger logger.ILogger
	userRepository IUserRepository
	passwordHasher PasswordHasher
}

func NewUserService(userRepository IUserRepository, logger logger.ILogger) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger: logger,
	}
}

func (us *UserService) SignUp(context context.Context, userDetails *pb.UserDetails) (*pb.UserResult, error) {

	us.logger.LogTrace("SignUp")

	userPassword := userDetails.Password
	hashedUserPassword, err := us.passwordHasher.hashAndSalt(userPassword)
	if err != nil {
		return nil, err
	}

	userModel := &types.UserModel {
		Username: userDetails.Username,
		Password: *hashedUserPassword,
	}

	rowsAffected, err := us.userRepository.CreateUser(userModel)
	if err != nil {
		return nil, err
	}

	signUpResult := &pb.UserResult{
		Username: userDetails.Username,
	}

	if rowsAffected > 0 {
		
	}

	return signUpResult, nil
}

func (us *UserService) SignIn(context context.Context, signInDetails *pb.UserDetails) (*pb.UserResult, error){
	return nil, nil
}