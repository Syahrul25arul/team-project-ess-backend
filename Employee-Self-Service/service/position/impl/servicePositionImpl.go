package servicePositionImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	repositoryPosition "employeeSelfService/repository/position"
	"employeeSelfService/request"
	"employeeSelfService/response"
)

type ServicePositionImpl struct {
	repo repositoryPosition.RepositoryPosition
}

func NewRepositoryPositionImpl(repo repositoryPosition.RepositoryPosition) ServicePositionImpl {
	return ServicePositionImpl{repo: repo}
}

func (s ServicePositionImpl) Save(position *request.Position) (response.ResponsePosition, *errs.AppErr) {

	testDomain := domain.Position{}
	testDomain.PositionName = position.PositionName
	err := s.repo.Save(&testDomain)
	if err != nil {
		return response.NewResponsePositionFailed(500, "TEST error !!"), err
	}
	return *response.NewReponsePositionSuccess(), nil
}

func (s ServicePositionImpl) FindById(id int64) (*domain.Position, *errs.AppErr) {

	positions, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (s ServicePositionImpl) Delete(id int64) (*domain.Position, *errs.AppErr) {
	positions, err := s.repo.Delete(id)
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (s ServicePositionImpl) Update(position domain.Position) *errs.AppErr {
	if err := s.repo.Update(position); err != nil {
		logger.Error("Error Update Data Position")
		return err
	}
	return s.repo.Update(position)
}
