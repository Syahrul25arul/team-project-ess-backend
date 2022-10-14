package repositoryPosition

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoryPosition interface {
	Save(position *domain.Position) *errs.AppErr
	FindById(id int64) (*domain.Position, *errs.AppErr)
	Delete(id int64) (*domain.Position, *errs.AppErr)
	Update(position domain.Position) *errs.AppErr
}
