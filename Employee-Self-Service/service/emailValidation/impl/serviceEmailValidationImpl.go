package serviceEmailValidationImpl

import (
	"employeeSelfService/domain"
	repositoryEmailValidation "employeeSelfService/repository/emailValidation"
	repositoryUser "employeeSelfService/repository/user"
	"employeeSelfService/response"
)

type ServiceEmaiValidationImpl struct {
	repo repositoryEmailValidation.RepositoryEmailValidation
	user repositoryUser.RepositoryUser
}

func NewServiceEmailValidationImpl(repo repositoryEmailValidation.RepositoryEmailValidation, repoUser repositoryUser.RepositoryUser) ServiceEmaiValidationImpl {
	return ServiceEmaiValidationImpl{repo, repoUser}
}

func (service ServiceEmaiValidationImpl) Save(emailValidation *domain.EmailValidation, id int64) response.ReponseEmailValidation {

	// save email validation ke database
	resp := service.repo.Save(emailValidation)

	// cek apakah ada error atau tidak
	if resp != nil {
		return response.NewResponseEmailValidationFailed(resp.Code, resp.Message)
	}
	return response.NewResponseEmailValidationSuccess()

}
