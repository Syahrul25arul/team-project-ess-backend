package repositoryEmailValidation

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryEmailValidation interface {
	Save(emailValidation *domain.EmailValidation) *errs.AppErr
}
