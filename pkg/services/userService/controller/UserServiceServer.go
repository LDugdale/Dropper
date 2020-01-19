package controller

import (
	"github.com/LDugdale/Dropper/pkg/gRpc"

	pb "github.com/LDugdale/Dropper/proto"
	"github.com/LDugdale/dropper/pkg/services/userService/abstractions"
	"github.com/ldugdale/dropper/pkg/commonAbstractions"
	"github.com/ldugdale/dropper/pkg/customErrors"
	"github.com/ldugdale/dropper/pkg/log"
	"golang.org/x/net/context"
)

type UserServiceServer struct {
	logger      log.Logger
	userService abstractions.UserService
}

func NewUserServiceServer(logger log.Logger, userService abstractions.UserService) *UserServiceServer {
	return &UserServiceServer{
		logger:      logger,
		userService: userService,
	}
}

func (us *UserServiceServer) SignUp(context context.Context, userDetails *pb.UserDetails) (*pb.UserResult, error) {

	us.logger.LogTrace("SignUp")

	userModel := &commonAbstractions.UserModel{
		Username: userDetails.Username,
		Password: userDetails.Password,
	}

	result, err := us.userService.SignUp(userModel)
	if err != nil {
		return nil, attachErrorMetadata(err)
	}

	signUpResult := &pb.UserResult{
		Username: result.Username,
	}

	return signUpResult, nil
}

func (us *UserServiceServer) SignIn(context context.Context, userDetails *pb.UserDetails) (*pb.UserResult, error) {

	us.logger.LogTrace("SignIn")

	userModel := &commonAbstractions.UserModel{
		Username: userDetails.Username,
		Password: userDetails.Password,
	}

	result, err := us.userService.SignIn(userModel)
	if err != nil {
		return nil, attachErrorMetadata(err)
	}

	signInResult := &pb.UserResult{
		Username: result.Username,
	}

	return signInResult, nil
}

func attachErrorMetadata(err error) error {

	if dropperErr, ok := err.(customErrors.DropperError); ok {
		details := customErrors.GetErrorDetails(dropperErr)
		statusCode := customErrors.GetType(dropperErr)

		grpcErr := gRpc.AttachProtobufMetadata(int(statusCode), "user", details.Message, details.Description)

		return grpcErr

	}

	return err
}
