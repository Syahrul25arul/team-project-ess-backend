package repositoryUserImpl

import (
	domainUser "employeeSelfService/domain/user"
	"employeeSelfService/errs"
	"employeeSelfService/logger"

	"gorm.io/gorm"
)

type RepositoryUserImpl struct {
	db *gorm.DB
}

func NewRepositoryUserImpl(client *gorm.DB) RepositoryUserImpl {
	return RepositoryUserImpl{client}
}

func (repo RepositoryUserImpl) FindByEmail(email string) (*domainUser.User, *errs.AppErr) {
	var user domainUser.User
	if result := repo.db.Where("email = ?", email).Find(&user); result.RowsAffected == 0 {
		logger.Error("error get data user by email not found")
		return nil, errs.NewAuthenticationError("Your Login Failed! Invalid Credential")
	}

	return &user, nil
}
