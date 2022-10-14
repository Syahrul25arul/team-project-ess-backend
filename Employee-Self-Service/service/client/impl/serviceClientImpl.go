package serviceClientImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryClient "employeeSelfService/repository/client"
	"employeeSelfService/response"
	"net/http"
)

type serviceClientImpl struct {
	repo repositoryClient.RepositoryClient
}

func NewServiceClient(repo repositoryClient.RepositoryClient) serviceClientImpl {
	return serviceClientImpl{repo: repo}
}

func (s serviceClientImpl) SaveClient(client *domain.Client) (*helper.SuccessResponseMessage, *errs.AppErr) {
	idClient := client.IdClient
	err := s.repo.Save(client)
	if err != nil {
		return nil, err
	} else {

		if idClient != 0 {
			return helper.NewSuccessResponseMessage(http.StatusOK, "client", "updated"), nil
		} else {
			return helper.NewSuccessResponseMessage(http.StatusCreated, "client", "created"), nil
		}
	}
}

func (s serviceClientImpl) GetAllClient() (*response.ResponseClient, *errs.AppErr) {
	result, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	} else {
		response := response.NewResponseClientSuccess(result)
		return response, nil
	}
}

func (s serviceClientImpl) GetClientById(id int) (*response.ResponseClient, *errs.AppErr) {
	result, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	} else {
		response := response.NewResponseClientSuccess(result)
		return response, nil
	}
}

func (s serviceClientImpl) DeleteClient(id int) (*helper.SuccessResponseMessage, *errs.AppErr) {
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	} else {
		response := helper.NewSuccessResponseMessage(http.StatusOK, "client", "deleted")
		return response, nil
	}
}
func (s serviceClientImpl) Update(client *domain.Client) (*helper.SuccessResponseMessage, *errs.AppErr) {
	if _, err := s.repo.GetById(int(client.IdClient)); err != nil {
		return nil, err
	} else {

		if err := s.repo.Save(client); err != nil {
			return nil, err
		} else {
			return helper.NewSuccessResponseMessage(http.StatusCreated, "client", "updated"), nil
		}
	}
}
