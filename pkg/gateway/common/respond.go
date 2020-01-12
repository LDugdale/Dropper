package common

import (
	"net/http"
	"encoding/json"
	"github.com/ldugdale/dropper/pkg/log"
)

type Response struct {
	logger log.Logger
}

func NewResponse(logger log.Logger) *Response{
	return &Response {
		logger: logger,
	}
}

func (r *Response)Respond(response http.ResponseWriter, request *http.Request, data interface{}, status int){

	response.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(response).Encode(data); err != nil {
			r.logger.LogError("Error ocurred")
		}
	}
}