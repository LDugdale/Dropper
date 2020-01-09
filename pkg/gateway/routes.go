package gateway

func (c *Controller) routes() {

	c.Router.HandleFunc("/events", c.handleGetEvents()).Methods("GET")
}
