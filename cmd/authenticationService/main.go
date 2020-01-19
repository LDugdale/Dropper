package main

import (
	"flag"

	"github.com/LDugdale/Dropper/pkg/gRpc"
	"github.com/LDugdale/Dropper/pkg/services/authenticationService/controller"
	"github.com/LDugdale/Dropper/pkg/services/authenticationService/services"
	pb "github.com/LDugdale/Dropper/proto"
)

var port = *flag.String("l", ":7100", "Specify the port that the server will listen on")

func main() {

	service := initializeAuthenticationService()

	server := gRpc.SetUpServer(port)

	pb.RegisterAuthenticationServiceServer(server, service)
}

func initializeAuthenticationService() *controller.AuthenticationServiceServer {

	authenticationService := services.NewAuthenticationService()
	authenticationServiceServer := controller.NewAuthenticationServiceServer(authenticationService)

	return authenticationServiceServer
}
