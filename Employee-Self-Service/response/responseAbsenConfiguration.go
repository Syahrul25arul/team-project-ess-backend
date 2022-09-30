package response

import "net/http"

type ResponseAbsenConfiguration struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewResponseAbsenConfiguration() ResponseAbsenConfiguration {
	return ResponseAbsenConfiguration{
		Code:    http.StatusCreated,
		Status:  "ok",
		Message: "Absen configuration has been created",
	}
}

func NewResponseAbsenConfigurationFailed(code int, message string) ResponseAbsenConfiguration {
	return ResponseAbsenConfiguration{
		Code:    code,
		Status:  "error",
		Message: message,
	}
}
