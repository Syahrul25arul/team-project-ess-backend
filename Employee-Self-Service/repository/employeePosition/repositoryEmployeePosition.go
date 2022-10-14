package repositoryEmployeePosition

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryEmployeePosition interface {
	Save(employeePosition *domain.EmployeePosition) *errs.AppErr
	Delete(id int64) *errs.AppErr
	FindById(id int64) (*domain.EmployeePosition, *errs.AppErr)
}
