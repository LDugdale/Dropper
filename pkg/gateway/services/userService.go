package services

import (
	pb "github.com/LDugdale/dropper/proto"
	"github.com/ldugdale/dropper/pkg/commonAbstractions"
	"github.com/ldugdale/dropper/pkg/log"
	"golang.org/x/net/context"
)

type UserService struct {
	logger                      log.Logger
	userServiceClient           pb.UserServiceClient
	authenticationServiceClient pb.AuthenticationServiceClient
}

func NewUserService(logger log.Logger, userServiceClient pb.UserServiceClient, authenticationServiceClient pb.AuthenticationServiceClient) *UserService {
	userService := &UserService{
		logger:                      logger,
		userServiceClient:           userServiceClient,
		authenticationServiceClient: authenticationServiceClient,
	}

	return userService
}

func (us *UserService) SignUp(user commonAbstractions.UserModel) (*commonAbstractions.User, error) {

	userDetails := &pb.UserDetails{
		Username: user.Username,
		Password: user.Password,
	}

	result, err := us.userServiceClient.SignUp(context.Background(), userDetails)
	us.logger.LogTrace("result :", result, err)
	if err != nil {
		return nil, err
	}

	claimsDetails := &pb.ClaimsDetails{
		Username: user.Username,
	}

	tokenResult, err := us.authenticationServiceClient.CreateToken(context.Background(), claimsDetails)
	if err != nil {
		return nil, err
	}

	

	userReturn := &commonAbstractions.User{
		Username: result.Username,
	}

	return userReturn, nil
}

func (us *UserService) SignIn(user commonAbstractions.UserModel) (*commonAbstractions.User, error) {

	userDetails := &pb.UserDetails{
		Username: user.Username,
		Password: user.Password,
	}

	result, err := us.userServiceClient.SignIn(context.Background(), userDetails)
	us.logger.LogTrace("result :", result, err)
	if err != nil {
		return nil, err
	}

	userReturn := &commonAbstractions.User{
		Username: result.Username,
	}

	return userReturn, nil
}
