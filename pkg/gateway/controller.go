package gateway

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ldugdale/dropper/pkg/log"
	pb "github.com/LDugdale/Dropper/proto"
)

type Controller struct {
	Router *mux.Router
	logger log.Logger
	geoPostServiceClient pb.GeoPostServiceClient
	authenticationServiceClient pb.AuthenticationServiceClient
	userServiceClient pb.UserServiceClient
}

func NewController(logger log.Logger, geoPostServiceClient pb.GeoPostServiceClient, authenticationServiceClient pb.AuthenticationServiceClient, userServiceClient pb.UserServiceClient) Controller {
	
	c := Controller{
		Router: mux.NewRouter().StrictSlash(true),
		logger: logger,
		geoPostServiceClient: geoPostServiceClient,
		authenticationServiceClient: authenticationServiceClient,
		userServiceClient: userServiceClient,		
	}

	c.routes()
	c.userRoutes()
	return c
}

func (c *Controller) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int){

	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			c.logger.LogError("Error ocurred")
		}
	}
}