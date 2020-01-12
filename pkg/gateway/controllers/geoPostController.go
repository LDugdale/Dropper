package controllers

// import (
//     "encoding/json"
// 	"net/http"
// 	"github.com/gorilla/mux"
//     "golang.org/x/net/context"
//     "github.com/ldugdale/dropper/pkg/log"
//     pb "github.com/LDugdale/Dropper/proto"
// )

// type GeoPostController struct {
//     Router *mux.Router
//     logger log.Logger
//     userServiceClient pb.UserServiceClient
// }

// func NewGeoPostController(logger log.Logger) *GeoPostController{
//     geoPostController = &GeoPostController {
//         Router: mux.NewRouter().StrictSlash(true),
//         logger: logger,
//     }

//     return geoPostController
// }

// func (c *GeoPostController) routes() {
// 	c.Router.HandleFunc("/user/signup", c.handleSignUp()).Methods("POST")
// }

// func (c *GeoPostController) handleSignUp() http.HandlerFunc {
    
//     return func(w http.ResponseWriter, r *http.Request) {
        
//         c.logger.LogTrace("handleSignUp")

//         var user User

//         err := json.NewDecoder(r.Body).Decode(&user)
//         if err != nil {
            
//         }

//         parameters := new(pb.UserDetails)
//         result, err := c.userServiceClient.SignUp(context.Background(), parameters)

//         if err != nil {
//             c.logger.LogTrace("ServiceClient: ", err)
//             c.respond(w, r, nil, http.StatusBadRequest)
//         }

//         c.respond(w, r, result, http.StatusOK)
//     }
// }