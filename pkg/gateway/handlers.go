package gateway

import (
    "net/http"
    "golang.org/x/net/context"
    pb "github.com/LDugdale/Dropper/proto"
)

func (c *Controller) handleGetEvents() http.HandlerFunc {
    
    return func(w http.ResponseWriter, r *http.Request) {
        
        c.logger.LogTrace("handleGetEvents")

        parameters := new(pb.AddGeoPostParameters)
        result, err := c.geoPostServiceClient.AddGeoPost(context.Background(), parameters)

        if err != nil {
            c.logger.LogError("AddGeoPost err: %v")
            c.respond(w, r, nil, http.StatusBadRequest)
        }

        c.respond(w, r, result, http.StatusOK)
    }
}