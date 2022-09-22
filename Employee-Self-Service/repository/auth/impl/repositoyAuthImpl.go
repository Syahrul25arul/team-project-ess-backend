package repositoryUserImpl

import (
	domainAuth "employeeSelfService/domain/auth"
	"employeeSelfService/errs"
	"employeeSelfService/logger"

	"gorm.io/gorm"
)

type RepositoryAuthImpl struct {
	db *gorm.DB
}

func NewRepositoryAuthImpl(client *gorm.DB) RepositoryAuthImpl {
	return RepositoryAuthImpl{client}
}

func (repo RepositoryAuthImpl) FindByEmail(email string) (*domainAuth.Auth, *errs.AppErr) {
	var auth domainAuth.Auth

	// get data user for payload token by email
	// result := repo.db.Table("users").Select("users.id_user,users.email,users.user_role,employee.id_employe,employee.nama_lengkap,position.id_position,position.position_name").Joins("left join employee on users.id_user = employee.id_user").Joins("left join employee_position on employee.id_employe = employee_position.id_employe").Joins("left join position on employee_position.id_position = position.id_position").Scan(&auth)
	result := repo.db.Table("users").Select("users.id_user,users.email,users.user_role,users.password,employee.id_employe,employee.nama_lengkap").Joins("left join employee on users.id_user = employee.id_user").Scan(&auth)

	if result.RowsAffected == 0 {
		logger.Error("error get data user by email not found")
		return nil, errs.NewAuthenticationError("Your Login Failed! Invalid Credential")
	}

	return &auth, nil
}
