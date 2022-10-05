package serviceEmailValidation

import (
	"employeeSelfService/domain"
	"employeeSelfService/response"
)

type ServiceEmailValidation interface {
	Save(emailValidation *domain.EmailValidation, id int64) response.ReponseEmailValidation
}
