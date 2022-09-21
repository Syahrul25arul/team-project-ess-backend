package repositoryUser

import (
	domainUser "employeeSelfService/domain/user"
	"employeeSelfService/errs"
)

type RepositoryUse interface {
	FindByEmail(email string) (*domainUser.User, *errs.AppErr)
}
