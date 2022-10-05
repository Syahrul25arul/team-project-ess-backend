package repositoryRegisterImpl

import (
	"database/sql"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"

	"gorm.io/gorm"
)

type RepositoryRegisterImpl struct {
	db *gorm.DB
}

func NewRepositoryRegisterImpl(client *gorm.DB) RepositoryRegisterImpl {
	return RepositoryRegisterImpl{client}
}

func (repo RepositoryRegisterImpl) Register(user *domain.User, employee *domain.Employee) *errs.AppErr {
	// begin transaction
	tx := repo.db.Begin()

	// save data employee
	if result := tx.Create(user); result.Error != nil {
		// if error rollback
		tx.Rollback()
		logger.Error("error insert data user : " + result.Error.Error())
		return errs.NewUnexpectedError("Sorry, Internal Server ERror")
	}

	// get id user in employee
	employee.IdUser = sql.NullInt64{Int64: int64(user.IdUser), Valid: true}

	if result := tx.Create(employee); result.Error != nil {
		// if error rollback
		logger.Error("error insert data user : " + result.Error.Error())
		tx.Rollback()
		return errs.NewUnexpectedError("Sorry, Internal Server ERror")
	}

	tx.Commit()
	return nil
}
