package serviceDashboardImpl

import (
	"employeeSelfService/errs"
	repositoryUser "employeeSelfService/repository/user"
	"employeeSelfService/response"
	"strconv"
)

type serviceDashboardImpl struct {
	repo repositoryUser.RepositoryUser
}

func NewServiceDashboard(repo repositoryUser.RepositoryUser) serviceDashboardImpl {
	return serviceDashboardImpl{repo: repo}
}

func (s serviceDashboardImpl) GetDashboard(id string) (*response.ResponseDashboard, *errs.AppErr) {

	// convert to int for findById
	idInt, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return nil, errs.NewBadRequestError("id not relevant")
	}

	// get user by id
	if _, errs := s.repo.FindById(idInt); errs != nil {
		return nil, errs
	} else {
		return s.repo.GetDataDashboard(id)
	}
}
