package repositoryAuth

import (
	domainAuth "employeeSelfService/domain/auth"
	"employeeSelfService/errs"
)

type RepositoryAuth interface {
	FindByEmail(email string) (*domainAuth.Auth, *errs.AppErr)
}
