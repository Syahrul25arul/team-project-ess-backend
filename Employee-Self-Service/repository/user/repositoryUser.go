package repositoryUser

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/response"
)

type RepositoryUser interface {
	FindByEmail(email string) (*domain.User, *errs.AppErr)
	FindById(id int64) (*domain.User, *errs.AppErr)
	GetDataDashboard(id string) (*response.ResponseDashboard, *errs.AppErr)
}
