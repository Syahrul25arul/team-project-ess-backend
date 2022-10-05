package response

import "net/http"

type ReponseEmailValidation struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewResponseEmailValidationSuccess() ReponseEmailValidation {
	return ReponseEmailValidation{
		Code:    http.StatusCreated,
		Status:  "ok",
		Message: "Email for validation register has been created",
	}
}

func NewResponseEmailValidationFailed(code int, message string) ReponseEmailValidation {
	return ReponseEmailValidation{
		Code:    code,
		Status:  "error",
		Message: message,
	}
}
