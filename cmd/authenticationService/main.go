package main

import (
	"flag"
	"net"

	"github.com/LDugdale/Dropper/pkg/log"
	"github.com/LDugdale/Dropper/pkg/services/authenticationService/controller"
	"github.com/LDugdale/Dropper/pkg/services/authenticationService/services"
	pb "github.com/LDugdale/Dropper/proto"
	"google.golang.org/grpc"
)

var port = *flag.String("l", ":7100", "Specify the port that the server will listen on")

func main() {

	logger := log.NewLogger()

	service := initializeAuthenticationService(logger)

	//server := gRpc.SetUpServer(port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.LogError("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(server, service)

	logger.LogTrace("Running on port:", port)
	if err := server.Serve(lis); err != nil {
		logger.LogError("failed to serve: %v", err)
	}

	pb.RegisterAuthenticationServiceServer(server, service)
}

func initializeAuthenticationService(logger log.Logger) *controller.AuthenticationServiceServer {

	authenticationService := services.NewAuthenticationService(logger)
	authenticationServiceServer := controller.NewAuthenticationServiceServer(logger, authenticationService)

	return authenticationServiceServer
}

// func main() {

// 	logger := log.NewLogger()

// 	service := initializeUserService(logger)

// 	//server := gRpc.SetUpServer(port)

// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		logger.LogError("failed to listen: %v", err)
// 	}

// 	server := grpc.NewServer()
// 	pb.RegisterUserServiceServer(server, service)

// 	logger.LogTrace("Running on port:", port)
// 	if err := server.Serve(lis); err != nil {
// 		logger.LogError("failed to serve: %v", err)
// 	}

// }

// func initializeUserService(logger log.Logger) *controller.UserServiceServer {

// 	database, err := database.NewDB(dataSourceName)
// 	if err != nil {
// 		logger.LogError(err)
// 	}

// 	passwordHasher := cryptography.NewPasswordHasher()

// 	userRepository := *data.NewUserRepository(logger, *database)
// 	userService := services.NewUserService(logger, &userRepository, passwordHasher)
// 	userServiceServer := controller.NewUserServiceServer(logger, userService)

// 	return userServiceServer
// }
