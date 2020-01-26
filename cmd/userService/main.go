package main

import (
	"net"

	pb "github.com/LDugdale/Dropper/proto"
	"github.com/ldugdale/dropper/pkg/log"

	"github.com/ldugdale/dropper/pkg/cryptography"
	"github.com/ldugdale/dropper/pkg/database"
	"github.com/ldugdale/dropper/pkg/services/userService/controller"
	"github.com/ldugdale/dropper/pkg/services/userService/data"
	"github.com/ldugdale/dropper/pkg/services/userService/services"
	"google.golang.org/grpc"
)

var port = "localhost:7101"
var dataSourceName = "root:password@(localhost)/dropper"

func main() {

	logger := log.NewLogger()

	service := initializeUserService(logger)

	//server := gRpc.SetUpServer(port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.LogError("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, service)

	logger.LogTrace("Running on port:", port)
	if err := server.Serve(lis); err != nil {
		logger.LogError("failed to serve: %v", err)
	}

}

func initializeUserService(logger log.Logger) *controller.UserServiceServer {

	database, err := database.NewDB(dataSourceName)
	if err != nil {
		logger.LogError(err)
	}

	passwordHasher := cryptography.NewPasswordHasher()

	userRepository := *data.NewUserRepository(logger, *database)
	userService := services.NewUserService(logger, &userRepository, passwordHasher)
	userServiceServer := controller.NewUserServiceServer(logger, userService)

	return userServiceServer
}
