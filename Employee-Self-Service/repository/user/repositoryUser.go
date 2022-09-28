package repositoryUser

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryUser interface {
	FindByEmail(email string) (*domain.User, *errs.AppErr)
}
