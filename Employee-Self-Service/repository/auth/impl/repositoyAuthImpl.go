package repositoryUserImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	"errors"

	"gorm.io/gorm"
)

type RepositoryAuthImpl struct {
	db *gorm.DB
}

func NewRepositoryAuthImpl(client *gorm.DB) RepositoryAuthImpl {
	return RepositoryAuthImpl{client}
}

func (repo RepositoryAuthImpl) FindByEmail(email string) (*domain.Auth, *errs.AppErr) {
	var auth domain.Auth

	// get data user for payload token by email
	// result := repo.db.Table("users").Select("users.id_user,users.email,users.user_role,employee.id_employe,employee.nama_lengkap,position.id_position,position.position_name").Joins("left join employee on users.id_user = employee.id_user").Joins("left join employee_position on employee.id_employe = employee_position.id_employe").Joins("left join position on employee_position.id_position = position.id_position").Scan(&auth)
	result := repo.db.Table("users").Select("users.id_user,users.email,users.user_role,users.password,employee.id_employe,employee.nama_lengkap").Joins("left join employee on users.id_user = employee.id_user").Where("users.email = ?", email).Scan(&auth)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error("error get data user by email not found : " + err.Error())
			return nil, errs.NewAuthenticationError("Your Login Failed! Invalid Credential")
		} else {
			logger.Error("error get data user by email not found : " + err.Error())
			return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again! " + err.Error())
		}
	}

	return &auth, nil
}
