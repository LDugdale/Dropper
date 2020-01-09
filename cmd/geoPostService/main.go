package main

import (
	"flag"
	pb "github.com/LDugdale/Dropper/proto"
	"github.com/ldugdale/dropper/pkg/geoPostService"
	"github.com/ldugdale/dropper/pkg/logger"
	"github.com/LDugdale/Dropper/pkg/gRpc"
)

var port = *flag.String("l", ":7100", "Specify the port that the server will listen on")

func main() {

	service := initializeGeoPostService()

	server := gRpc.SetUpServer(port)

	pb.RegisterGeoPostServiceServer(server, service)
}

func initializeGeoPostService() geoPost.GeoPostService {

	logger := logger.NewLogger()

	geoPostService := geoPost.NewGeoPostService(logger)

	return geoPostService
}