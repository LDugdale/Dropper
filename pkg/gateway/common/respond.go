package common

import (
	"encoding/json"
	"net/http"

	"github.com/ldugdale/dropper/pkg/log"
)

type Response struct {
	logger log.Logger
}

func NewResponse(logger log.Logger) *Response {
	return &Response{
		logger: logger,
	}
}

func (r *Response) Respond(response http.ResponseWriter, request *http.Request, status int, data interface{}) {

	response.WriteHeader(status)

	if data != nil {

		if err := json.NewEncoder(response).Encode(data); err != nil {
			r.logger.LogError("Error ocurred")
		}
	}
}
