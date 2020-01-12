
package main

import (
	"fmt"
	"os"
	"net/http"
	"google.golang.org/grpc"
	"github.com/gorilla/mux"
	pb "github.com/LDugdale/Dropper/proto"
	"github.com/ldugdale/dropper/pkg/gateway"
	"github.com/ldugdale/dropper/pkg/log"
	"github.com/ldugdale/dropper/pkg/gateway/controllers"
	"github.com/ldugdale/dropper/pkg/gateway/common"
	"github.com/ldugdale/dropper/pkg/gateway/services"
)


var address = ":7100"


func main() {


	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error { 


	logger, controller, router := initializeGateway()

	controller.Start()

	logger.LogError(http.ListenAndServe(":8080", router))

	return nil
}


func initializeGateway() (log.Logger, *gateway.Controller, *mux.Router) {

	logger := log.NewLogger()
	router := mux.NewRouter().StrictSlash(true)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.LogError("Did not connect: %v", err)
	}
	
	//defer conn.Close()

	//geoPostServiceClient := pb.NewGeoPostServiceClient(conn)
	//authenticationServiceClient := pb.NewAuthenticationServiceClient(conn)
	userServiceClient := pb.NewUserServiceClient(conn)

	response := *common.NewResponse(logger)

	userService := services.NewUserService(logger, userServiceClient)
	userController := *controllers.NewUserController(logger, *router, response, userService)	
	
	controller := gateway.NewController(logger, userController)

	return logger, controller, router
}