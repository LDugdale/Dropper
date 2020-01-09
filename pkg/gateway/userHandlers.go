package gateway

import (
    "net/http"
    "golang.org/x/net/context"
    pb "github.com/LDugdale/Dropper/proto"
)

func (c *Controller) handleSignUp() http.HandlerFunc {
    
    return func(w http.ResponseWriter, r *http.Request) {
        
        c.logger.LogTrace("handleSignUp")

        parameters := new(pb.UserDetails)
        result, err := c.userServiceClient.SignUp(context.Background(), parameters)

        if err != nil {
            c.logger.LogTrace("ServiceClient: ", err)
            c.respond(w, r, nil, http.StatusBadRequest)
        }

        c.respond(w, r, result, http.StatusOK)
    }
}