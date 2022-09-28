package serviceLogin

import (
	"employeeSelfService/errs"
	"employeeSelfService/request"
	"employeeSelfService/response"
)

type ServiceLogin interface {
	Login(loginRequest *request.LoginRequest) (*response.ResponseLogin, *errs.AppErr)
}
