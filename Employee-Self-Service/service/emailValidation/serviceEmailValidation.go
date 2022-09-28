package serviceEmailValidation

import (
	"employeeSelfService/domain"
	"employeeSelfService/response"
)

type ServiceEmailValidation interface {
	Save(emailValidation *domain.EmailValidation) response.ReponseEmailValidation
}
