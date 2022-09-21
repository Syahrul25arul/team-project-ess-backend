package loginImpl

import (
	domainUser "employeeSelfService/domain/user"
	"employeeSelfService/errs"

	"gorm.io/gorm"
)

type RepositoryLoginImpl struct {
	DB *gorm.DB
}

func NewRepositoryLoginImpl(client *gorm.DB) RepositoryLoginImpl {
	return RepositoryLoginImpl{client}
}

func (repo RepositoryLoginImpl) Login(user *domainUser.User) *errs.AppErr {
	return nil
}
