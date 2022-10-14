package servicePosition

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/request"
	"employeeSelfService/response"
)

type ServicePosition interface {
	Save(position *request.Position) (response.ResponsePosition, *errs.AppErr)
	FindById(id int64) (*domain.Position, *errs.AppErr)
	Delete(id int64) (*domain.Position, *errs.AppErr)
	Update(position domain.Position) *errs.AppErr
}
