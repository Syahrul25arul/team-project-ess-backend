package serviceProjectImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryProject "employeeSelfService/repository/project"
	"employeeSelfService/response"
	"net/http"
)

type serviceProjectImpl struct {
	repo repositoryProject.RepositoryProject
}

func NewServiceProject(repo repositoryProject.RepositoryProject) serviceProjectImpl {
	return serviceProjectImpl{repo: repo}
}
func (s serviceProjectImpl) SaveProject(project *domain.Project) (*helper.SuccessResponseMessage, *errs.AppErr) {
	if err := s.repo.SaveProject(project); err != nil {
		return nil, err
	} else {
		return helper.NewSuccessResponseMessage(http.StatusCreated, "project", "created"), nil
	}
}

func (s serviceProjectImpl) GetAllProject() (*response.ResponseProject, *errs.AppErr) {
	if result, err := s.repo.GetAllProject(); err != nil {
		return nil, err
	} else {
		resp := response.NewResponseProject(http.StatusOK, "Get All Data Project", "Ok", result)
		return resp, nil
	}
}

func (s serviceProjectImpl) GetById(id int32) (*response.ResponseProject, *errs.AppErr) {
	if result, err := s.repo.GetById(id); err != nil {
		return nil, err
	} else {
		resp := response.NewResponseProject(http.StatusOK, "Get Data Project By Id", "Ok", result)
		return resp, nil
	}
}

func (s serviceProjectImpl) Update(project *domain.Project) (*helper.SuccessResponseMessage, *errs.AppErr) {
	if _, err := s.repo.GetById(project.IdProject); err != nil {
		return nil, err
	} else {
		resp := helper.NewSuccessResponseMessage(http.StatusOK, "project", "updated")
		return resp, nil
	}
}

func (s serviceProjectImpl) Delete(id int32) (*helper.SuccessResponseMessage, *errs.AppErr) {
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	} else {
		resp := helper.NewSuccessResponseMessage(http.StatusOK, "project", "deleted")
		return resp, nil
	}
}
