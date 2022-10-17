package repositoryProject

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryProject interface {
	SaveProject(project *domain.Project) *errs.AppErr
	GetAllProject() ([]domain.ProjectWithClient, *errs.AppErr)
	GetById(id int32) (*domain.ProjectWithClient, *errs.AppErr)
	Delete(id int32) *errs.AppErr
}
