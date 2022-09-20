package repositoryRegister

import (
	"employeeSelfService/errs"
	registerRequest "employeeSelfService/request/register"
)

type RepositoryInterface interface {
	Register(registerRequest *registerRequest.RegisterRequest) *errs.AppErr
}
