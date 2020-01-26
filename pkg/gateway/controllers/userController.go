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
	router                *mux.Router
	logger                log.Logger
	response              common.Response
	userService           abstractions.UserService
	authenticationService abstractions.AuthenticationService
}

func NewUserController(logger log.Logger, router *mux.Router, response common.Response, userService abstractions.UserService, authenticationService abstractions.AuthenticationService) *UserController {

	userController := &UserController{
		logger:                logger,
		router:                router,
		response:              response,
		userService:           userService,
		authenticationService: authenticationService,
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

		signUpResult, err := c.userService.SignUp(user)
		if err != nil {
			c.response.Respond(w, r, http.StatusBadRequest, gRpc.ExtractProtobufMetadata(err))
			return
		}

		tokenResult, err := c.authenticationService.CreateToken(signUpResult.Username)
		if err != nil {
			c.response.Respond(w, r, http.StatusBadRequest, gRpc.ExtractProtobufMetadata(err))
			return
		}

		userResult := commonAbstractions.UserWithToken{
			User: commonAbstractions.User{
				Username: signUpResult.Username,
			},
			Token: *tokenResult,
		}

		if err != nil {
			c.response.Respond(w, r, http.StatusBadRequest, gRpc.ExtractProtobufMetadata(err))
		} else {
			c.response.Respond(w, r, http.StatusOK, userResult)
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

		signInResult, err := c.userService.SignIn(user)
		if err != nil {
			c.response.Respond(w, r, http.StatusBadRequest, gRpc.ExtractProtobufMetadata(err))
			return
		}
		tokenResult, err := c.authenticationService.CreateToken(signInResult.Username)
		if err != nil {
			c.response.Respond(w, r, http.StatusBadRequest, gRpc.ExtractProtobufMetadata(err))
			return
		}

		c.logger.LogTrace(tokenResult)
		userResult := commonAbstractions.UserWithToken{
			User: commonAbstractions.User{
				Username: signInResult.Username,
			},
			Token: *tokenResult,
		}

		if err != nil {
			c.response.Respond(w, r, http.StatusBadRequest, gRpc.ExtractProtobufMetadata(err))
		} else {
			c.response.Respond(w, r, http.StatusOK, userResult)
		}
	}
}
