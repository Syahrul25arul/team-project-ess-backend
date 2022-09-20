package repositoryRegister

import (
	employeDomain "employeeSelfService/domain/employee"
	userDomain "employeeSelfService/domain/user"
	"employeeSelfService/errs"
)

type RepositoryInterface interface {
	Register(user *userDomain.User, employee *employeDomain.Employee) *errs.AppErr
}
