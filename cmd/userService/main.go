package main

import (
	"net"
	"fmt"
	"log"
	pb "github.com/LDugdale/Dropper/proto"
	"github.com/ldugdale/dropper/pkg/logger"
	"github.com/LDugdale/Dropper/pkg/gRpc"
	"github.com/ldugdale/dropper/pkg/userService"
	"github.com/ldugdale/dropper/pkg/database"

)

var port = "localhost:7100"
var dataSourceName = "root:password@(localhost)/dropper"


func main() {

	service := initializeUserService()

	server := gRpc.SetUpServer(port)


	netListener := getNetListener(7100)
	    // start the server
    if err := server.Serve(netListener); err != nil {
        log.Fatalf("failed to serve: %s", err)
    }

	pb.RegisterUserServiceServer(server, service)
}

func initializeUserService() *userService.UserService {

	database, err := database.NewDB(dataSourceName)
	if err != nil {
        log.Panic(err)
	}
	logger := logger.NewLogger()
	userRepository := *userService.NewUserRepository(*database)
	userService := userService.NewUserService(&userRepository, logger)
	
	return userService
}

func getNetListener(port uint) net.Listener {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
        panic(fmt.Sprintf("failed to listen: %v", err))
    }

    return lis
}
