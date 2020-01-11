
package main

import (
	"net/http"
	"fmt"
	"os"
	"google.golang.org/grpc"
	pb "github.com/LDugdale/Dropper/proto"
	"github.com/ldugdale/dropper/pkg/gateway"
	"github.com/ldugdale/dropper/pkg/log"
)


var address = ":7100"


func main() {


	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error { 

	logger := log.NewLogger()

	c := initializeGateway(logger)

	logger.LogError(http.ListenAndServe(":8080", c.Router))

	return nil
}


func initializeGateway(logger log.Logger) gateway.Controller {


	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.LogError("Did not connect: %v", err)
	}
	
	//defer conn.Close()

	geoPostServiceClient := pb.NewGeoPostServiceClient(conn)
	authenticationServiceClient := pb.NewAuthenticationServiceClient(conn)
	userServiceClient := pb.NewUserServiceClient(conn)

	
	controller := gateway.NewController(logger, geoPostServiceClient, authenticationServiceClient, userServiceClient)

	return controller
}