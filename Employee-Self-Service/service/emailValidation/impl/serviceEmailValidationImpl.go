package serviceEmailValidationImpl

import (
	"employeeSelfService/domain"
	repositoryEmailValidation "employeeSelfService/repository/emailValidation"
	"employeeSelfService/response"
)

type ServiceEmaiValidationImpl struct {
	repo repositoryEmailValidation.RepositoryEmailValidation
}

func NewServiceEmailValidationImpl(repo repositoryEmailValidation.RepositoryEmailValidation) ServiceEmaiValidationImpl {
	return ServiceEmaiValidationImpl{repo}
}

func (service ServiceEmaiValidationImpl) Save(emailValidation *domain.EmailValidation) response.ReponseEmailValidation {
	// save email validation ke database
	resp := service.repo.Save(emailValidation)

	// cek apakah ada error atau tidak
	if resp != nil {
		return response.NewResponseEmailValidationFailed(resp.Message)
	}
	return response.NewResponseEmailValidationSuccess()

}
