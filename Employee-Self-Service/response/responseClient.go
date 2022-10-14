package response

import (
	"net/http"
)

type ResponseClient struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    interface{}
}

func NewResponseClientSuccess(data interface{}) *ResponseClient {
	return &ResponseClient{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "success get data client",
		Data:    data,
	}
}
