package serviceLogin

import (
	"employeeSelfService/errs"
	requestLogin "employeeSelfService/request/login"
	responseLogin "employeeSelfService/response/login"
)

type ServiceLogin interface {
	Login(loginRequest *requestLogin.LoginRequest) (*responseLogin.ResponseLogin, *errs.AppErr)
}
