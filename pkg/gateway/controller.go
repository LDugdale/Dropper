package gateway

import (
	"github.com/ldugdale/dropper/pkg/log"
	"github.com/ldugdale/dropper/pkg/gateway/controllers"
)

type Controller struct {
	logger log.Logger
	userController controllers.UserController
}

func NewController(logger log.Logger, userController controllers.UserController) *Controller {
	
	controller := &Controller{
		logger: logger,
		userController: userController,
	}

	return controller
}

func (c *Controller) Start(){
	c.userController.Routes()
}