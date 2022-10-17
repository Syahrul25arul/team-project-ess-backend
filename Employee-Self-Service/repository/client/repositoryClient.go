package repositoryClient

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryClient interface {
	Save(client *domain.Client) *errs.AppErr
	GetAll() ([]domain.Client, *errs.AppErr)
	GetById(id int) (*domain.Client, *errs.AppErr)
	Delete(id int) *errs.AppErr
	GetAllWithProject() ([]domain.ClientWithProject, *errs.AppErr)
}
