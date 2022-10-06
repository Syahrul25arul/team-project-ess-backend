package repositoryClient

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryClient interface {
	Save(client *domain.Client) *errs.AppErr
	GetAll() ([]domain.Client, *errs.AppErr)
}
