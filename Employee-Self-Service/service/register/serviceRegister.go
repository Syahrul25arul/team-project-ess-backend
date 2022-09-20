package serviceRegister

import (
	"employeeSelfService/errs"
	registerRequest "employeeSelfService/request/register"
)

type ServiceRegister interface {
	Register(registerRequest *registerRequest.RegisterRequest) *errs.AppErr
}
