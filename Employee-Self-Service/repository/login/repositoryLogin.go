package repositorylogin

import (
	domainUser "employeeSelfService/domain/user"
	"employeeSelfService/errs"
)

type RepositoryLogin interface {
	Login(user *domainUser.User) *errs.AppErr
}
