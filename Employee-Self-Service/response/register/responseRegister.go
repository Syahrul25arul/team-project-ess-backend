package responseRegister

import "net/http"

type ResponseRegister struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewReponseRegisterSuccess() *ResponseRegister {
	return &ResponseRegister{
		Code:    http.StatusCreated,
		Message: "Your Account have been created",
		Status:  "ok",
	}
}
