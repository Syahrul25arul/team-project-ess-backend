package serviceRegister

import (
	"employeeSelfService/errs"
	"employeeSelfService/request"
)

type ServiceRegister interface {
	Register(registerRequest *request.RegisterRequest) *errs.AppErr
}
