package serviceAbsenConfigurationImpl

import (
	"employeeSelfService/errs"
	repositoryAbsenConfiguration "employeeSelfService/repository/absenConfiguration"
	repositoryUser "employeeSelfService/repository/user"
	"employeeSelfService/request"
	"employeeSelfService/response"
)

type ServiceAbsenConfigurationImpl struct {
	repo     repositoryAbsenConfiguration.RepositoryAbsenConfiguration
	repoUser repositoryUser.RepositoryUser
}

func NewServiceEmailValidationImpl(repo repositoryAbsenConfiguration.RepositoryAbsenConfiguration, repoUser repositoryUser.RepositoryUser) ServiceAbsenConfigurationImpl {
	return ServiceAbsenConfigurationImpl{repo, repoUser}
}

func (service ServiceAbsenConfigurationImpl) Save(absenConfiguration *request.AbsensiConfiguration, idUser int64) response.ResponseAbsenConfiguration {
	// cek apokah user adalah admin
	user, err := service.repoUser.FindById(idUser)

	// cek apakah ada error
	if err != nil {
		newError := errs.NewNotFoundError(err.Message)
		return response.NewResponseAbsenConfigurationFailed(newError.Code, newError.Message)
	}

	if user.UserRole != "admin" {
		newError := errs.NewForbiddenError("you dont have credential")
		return response.NewResponseAbsenConfigurationFailed(newError.Code, newError.Message)
	}

	// cek apakah data absen configuration sudah ada
	domainAbsenConfiguration, err := service.repo.GetData()
	if err != nil {
		return response.NewResponseAbsenConfigurationFailed(err.Code, err.Message)
	}
	if domainAbsenConfiguration == nil {
		domainAbsenConfiguration = absenConfiguration.ToDomainAbsen()

		if errSave := service.repo.Save(domainAbsenConfiguration); errSave != nil {
			return response.NewResponseAbsenConfigurationFailed(errSave.Code, errSave.Message)
		} else {
			return response.NewResponseAbsenConfiguration()
		}
	} else {
		idAbsen := domainAbsenConfiguration.IdAbsenConfiguration
		domainAbsenConfiguration = absenConfiguration.ToDomainAbsen()
		domainAbsenConfiguration.IdAbsenConfiguration = idAbsen

		if errSave := service.repo.Save(domainAbsenConfiguration); errSave != nil {
			return response.NewResponseAbsenConfigurationFailed(errSave.Code, errSave.Message)
		} else {
			return response.NewResponseAbsenConfiguration()
		}
	}
}
