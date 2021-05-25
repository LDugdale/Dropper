package main

import (
	"fmt"
	"net/http"
	"os"

	pb "github.com/LDugdale/dropper/proto"
	"github.com/gorilla/mux"
	"github.com/ldugdale/dropper/pkg/gateway"
	"github.com/ldugdale/dropper/pkg/gateway/common"
	"github.com/ldugdale/dropper/pkg/gateway/controllers"
	"github.com/ldugdale/dropper/pkg/gateway/services"
	"github.com/ldugdale/dropper/pkg/log"
	"google.golang.org/grpc"
)

var address = ":7100"

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {

	// logger, controller, router := initializeGateway()
	logger := log.NewLogger()
	router := mux.NewRouter().StrictSlash(true)
	router.Use(mux.CORSMethodMiddleware(router))

	//defer conn.Close()

	//geoPostServiceClient := pb.NewGeoPostServiceClient(conn)
	authenticationServiceClient := createAuthenticationServiceClient(logger, ":7100")
	userServiceClient := createUserServiceClient(logger, ":7101")

	response := *common.NewResponse(logger)

	userService := services.NewUserService(logger, userServiceClient)
	authenticationService := services.NewAuthenticationService(logger, authenticationServiceClient)
	userController := *controllers.NewUserController(logger, router, response, userService, authenticationService)

	controller := gateway.NewController(logger, userController)
	controller.Start()

	logger.LogError(http.ListenAndServe(":8080", router))

	return nil
}

func createAuthenticationServiceClient(logger log.Logger, address string) pb.AuthenticationServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.LogError("Did not connect: %v", err)
	}

	authenticationServiceClient := pb.NewAuthenticationServiceClient(conn)

	return authenticationServiceClient
}

func createUserServiceClient(logger log.Logger, address string) pb.UserServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.LogError("Did not connect: %v", err)
	}

	userServiceClient := pb.NewUserServiceClient(conn)

	return userServiceClient
}

// func initializeGateway() (log.Logger, *gateway.Controller, *mux.Router) {

// 	// logger := log.NewLogger()
// 	// router := mux.NewRouter().StrictSlash(true)

// 	// conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	// if err != nil {
// 	// 	logger.LogError("Did not connect: %v", err)
// 	// }

// 	// //defer conn.Close()

// 	// //geoPostServiceClient := pb.NewGeoPostServiceClient(conn)
// 	// //authenticationServiceClient := pb.NewAuthenticationServiceClient(conn)
// 	// userServiceClient := pb.NewUserServiceClient(conn)

// 	// response := *common.NewResponse(logger)

// 	// userService := services.NewUserService(logger, userServiceClient)
// 	// userController := *controllers.NewUserController(logger, *router, response, userService)

// 	// controller := gateway.NewController(logger, userController)

// 	// return logger, controller, *router
// }
