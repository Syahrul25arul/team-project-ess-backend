package response

import "net/http"

type ResponsePosition struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewReponsePositionSuccess() *ResponsePosition {
	return &ResponsePosition{
		Code:    http.StatusCreated,
		Message: "Your Account have been created",
		Status:  "ok",
	}
}

func NewReponsePositionDeleteSuccess() *ResponsePosition {
	return &ResponsePosition{
		Code:    http.StatusCreated,
		Message: "Delete Data Successfully",
		Status:  "ok",
	}
}

func NewReponsePositionUpdateSuccess() *ResponsePosition {
	return &ResponsePosition{
		Code:    http.StatusCreated,
		Message: "Update Data Successfully",
		Status:  "ok",
	}
}

func NewResponsePositionFailed(code int, message string) ResponsePosition {
	return ResponsePosition{
		Code:    code,
		Status:  "error",
		Message: message,
	}
}
