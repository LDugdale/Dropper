package controller

import (
	"golang.org/x/net/context"
	"github.com/LDugdale/Dropper/pkg/services/userService/abstractions"
	"github.com/ldugdale/dropper/pkg/logger"
	pb "github.com/LDugdale/Dropper/proto"
)

type UserServiceServer struct {
	logger logger.ILogger
	userService abstractions.UserService
}

func NewUserServiceServer(logger logger.ILogger, userService abstractions.UserService) *UserServiceServer {
	return &UserServiceServer{
		logger: logger,
		userService: userService,
	}
}

func (us *UserServiceServer) SignUp(context context.Context, userDetails *pb.UserDetails) (*pb.UserResult, error) {

	us.logger.LogTrace("SignUp")

	userModel := &abstractions.UserModel{
		Username: userDetails.Username,
		Password: userDetails.Password,
	}

	result, err := us.userService.SignUp(userModel)
	if err != nil {
		return nil, err
	}	

	signUpResult := &pb.UserResult{
		Username: result.Username,
	}

	return signUpResult, nil
}

func (us *UserServiceServer) SignIn(context context.Context, userDetails *pb.UserDetails) (*pb.UserResult, error) {
	
	us.logger.LogTrace("SignIn")

	userModel := &abstractions.UserModel{
		Username: userDetails.Username,
		Password: userDetails.Password,
	}

	result, err := us.userService.SignIn(userModel)
	if err != nil {
		return nil, err
	}	

	signInResult := &pb.UserResult{
		Username: result.Username,
	}

	return signInResult, nil}