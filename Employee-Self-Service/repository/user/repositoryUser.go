package repositoryUser

import (
	domainUser "employeeSelfService/domain/user"
	"employeeSelfService/errs"
)

type RepositoryUser interface {
	FindByEmail(email string) (*domainUser.User, *errs.AppErr)
}
