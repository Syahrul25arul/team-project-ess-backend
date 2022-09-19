package repositorylogin

import (
	"employeeSelfService/errs"
	loginRequest "employeeSelfService/request/login"
)

type RepositoryLogin interface {
	Login(loginRequest *loginRequest.LoginRequest) *errs.AppErr
}
