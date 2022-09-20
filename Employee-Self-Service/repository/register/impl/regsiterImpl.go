package repositoryRegisterImpl

import (
	"database/sql"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	registerRequest "employeeSelfService/request/register"
	"fmt"

	"gorm.io/gorm"
)

type RepositoryRegisterImpl struct {
	db *gorm.DB
}

func NewRepositoryRegisterImpl(client *gorm.DB) RepositoryRegisterImpl {
	return RepositoryRegisterImpl{client}
}

func (repo RepositoryRegisterImpl) Register(registerRequest *registerRequest.RegisterRequest) *errs.AppErr {
	// begin transaction
	tx := repo.db.Begin()

	// convert user
	user := registerRequest.ToUser()
	user.UserRole = "employee"
	user.StatusVerified = "true"

	//convert employee
	employee := registerRequest.ToEmployee()

	// // save data employee
	if result := tx.Create(user); result.Error != nil {
		// if error rollback
		tx.Rollback()
		logger.Error("error insert data user : " + result.Error.Error())
		return errs.NewUnexpectedError("Sorry, Internal Server ERror")
	}

	// get id user in employee
	employee.IdUser = sql.NullInt64{Int64: int64(user.IdUser), Valid: true}
	fmt.Println("==== ID USER ====", employee.IdUser)

	if result := tx.Create(employee); result.Error != nil {
		// if error rollback
		logger.Error("error insert data user : " + result.Error.Error())
		tx.Rollback()
		return errs.NewUnexpectedError("Sorry, Internal Server ERror")
	}

	tx.Commit()
	return nil
}
