package main

import (
	"flag"
	pb "github.com/LDugdale/Dropper/proto"
	"github.com/ldugdale/dropper/pkg/authenticationService"
	"github.com/LDugdale/Dropper/pkg/gRpc"
)

var port = *flag.String("l", ":7100", "Specify the port that the server will listen on")

func main() {

	service := initializeAuthenticationService()

	server := gRpc.SetUpServer(port)

	pb.RegisterAuthenticationServiceServer(server, service)
}

func initializeAuthenticationService() *authenticationService.AuthenticationService {

	authenticationService := authenticationService.NewAuthenticationService()

	return authenticationService
}