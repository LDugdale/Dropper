package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/LDugdale/Dropper/pkg/gRpc"
	"github.com/gorilla/mux"
	"github.com/ldugdale/dropper/pkg/commonAbstractions"
	"github.com/ldugdale/dropper/pkg/gateway/abstractions"
	"github.com/ldugdale/dropper/pkg/gateway/common"
	"github.com/ldugdale/dropper/pkg/log"
)

type UserController struct {
	router      *mux.Router
	logger      log.Logger
	response    common.Response
	userService abstractions.UserService
}

func NewUserController(logger log.Logger, router *mux.Router, response common.Response, userService abstractions.UserService) *UserController {

	userController := &UserController{
		logger:      logger,
		router:      router,
		response:    response,
		userService: userService,
	}

	return userController
}

func (c *UserController) Routes() {
	c.router.HandleFunc("/user/signup", c.handleSignUp()).Methods("POST")
	c.router.HandleFunc("/user/signin", c.handleSignIn()).Methods("POST")

	c.logger.LogTrace("User controller routes started")

}

func (c *UserController) handleSignUp() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		c.logger.LogTrace("handleSignUp")

		var user commonAbstractions.UserModel

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {

		}

		result, err := c.userService.SignUp(user)
		if err != nil {
			c.response.Respond(w, r, http.StatusBadRequest, gRpc.ExtractProtobufMetadata(err))
		} else {
			c.response.Respond(w, r, http.StatusOK, result)
		}

	}
}

func (c *UserController) handleSignIn() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		c.logger.LogTrace("handleSignIn")

		var user commonAbstractions.UserModel

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {

		}

		result, err := c.userService.SignIn(user)
		if err != nil {
			c.response.Respond(w, r, http.StatusBadRequest, gRpc.ExtractProtobufMetadata(err))
		} else {
			c.response.Respond(w, r, http.StatusOK, result)
		}

	}
}
