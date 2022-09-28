package repositorylogin

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryLogin interface {
	Login(user *domain.User) *errs.AppErr
}
