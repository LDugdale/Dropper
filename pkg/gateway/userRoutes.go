package gateway

func (c *Controller) userRoutes() {

	c.Router.HandleFunc("/user/signup", c.handleSignUp()).Methods("POST")
}
