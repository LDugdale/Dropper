package main

import (
	"net"
	"log"
	pb "github.com/LDugdale/Dropper/proto"
	"github.com/ldugdale/dropper/pkg/logger"
	//"github.com/LDugdale/Dropper/pkg/gRpc"
	"google.golang.org/grpc"
	"github.com/ldugdale/dropper/pkg/database"
	"github.com/ldugdale/dropper/pkg/services/userService/controller"
	"github.com/ldugdale/dropper/pkg/services/userService/services"
	"github.com/ldugdale/dropper/pkg/services/userService/data"
	"github.com/ldugdale/dropper/pkg/cryptography"
)

var port = "localhost:7100"
var dataSourceName = "root:password@(localhost)/dropper"


func main() {

	service := initializeUserService()

	//server := gRpc.SetUpServer(port)


	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, service)

	log.Println("Running on port:", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func initializeUserService() *controller.UserServiceServer {

	database, err := database.NewDB(dataSourceName)
	if err != nil {
        log.Panic(err)
	}
	logger := logger.NewLogger()
	passwordHasher := cryptography.NewPasswordHasher()

	userRepository := *data.NewUserRepository(*database)
	userService := services.NewUserService(logger, &userRepository, passwordHasher)
	userServiceServer := controller.NewUserServiceServer(logger, userService)
	
	return userServiceServer
}