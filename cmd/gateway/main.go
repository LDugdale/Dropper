
package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"google.golang.org/grpc"
	pb "github.com/LDugdale/Dropper/proto"
	"github.com/ldugdale/dropper/pkg/gateway"
	"github.com/ldugdale/dropper/pkg/logger"
)


var address = ":7100"


func main() {

	log.Println("Main")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error { 	

	c := initializeGateway()

	log.Fatal(http.ListenAndServe(":8080", c.Router))

	return nil
}


func initializeGateway() gateway.Controller {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	
	//defer conn.Close()

	geoPostServiceClient := pb.NewGeoPostServiceClient(conn)
	authenticationServiceClient := pb.NewAuthenticationServiceClient(conn)
	userServiceClient := pb.NewUserServiceClient(conn)

	logger := logger.NewLogger()
	
	controller := gateway.NewController(logger, geoPostServiceClient, authenticationServiceClient, userServiceClient)

	return controller
}