package loginImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"

	"gorm.io/gorm"
)

type RepositoryLoginImpl struct {
	DB *gorm.DB
}

func NewRepositoryLoginImpl(client *gorm.DB) RepositoryLoginImpl {
	return RepositoryLoginImpl{client}
}

func (repo RepositoryLoginImpl) Login(user *domain.User) *errs.AppErr {
	return nil
}
