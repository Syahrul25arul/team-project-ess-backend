package repositoryRegister

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryInterface interface {
	Register(user *domain.User, employee *domain.Employee) *errs.AppErr
}
