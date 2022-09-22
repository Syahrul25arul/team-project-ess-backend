package responseLogin

import "net/http"

type ResponseLogin struct {
	Code    int
	Message string
	Status  string
	Token   string
}

func NewLoginSucess(token string) *ResponseLogin {
	return &ResponseLogin{
		Token:   token,
		Message: "Your Login Success",
		Code:    http.StatusOK,
		Status:  "ok",
	}
}

func NewLoginFailed(code int, message string) *ResponseLogin {
	return &ResponseLogin{
		Code:    code,
		Message: message,
		Status:  "error",
	}
}
