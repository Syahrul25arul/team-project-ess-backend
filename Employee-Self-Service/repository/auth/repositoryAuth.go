package repositoryAuth

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryAuth interface {
	FindByEmail(email string) (*domain.Auth, *errs.AppErr)
}
