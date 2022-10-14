package repositoryEmployeePositionImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"

	"gorm.io/gorm"
)

type RepositoryEmployeePositionImpl struct {
	db *gorm.DB
}

func NewRepositoryEmployeePositionImpl(client *gorm.DB) RepositoryEmployeePositionImpl {
	return RepositoryEmployeePositionImpl{client}
}

func (repo RepositoryEmployeePositionImpl) Save(employeePosition *domain.EmployeePosition) *errs.AppErr {
	if result := repo.db.Create(employeePosition); result.Error != nil {
		logger.Error("error insert data employee position" + result.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}

func (repo RepositoryEmployeePositionImpl) FindById(id int64) (*domain.EmployeePosition, *errs.AppErr) {
	var employeeposition *domain.EmployeePosition = &domain.EmployeePosition{IdEmployeePosition: id}
	if result := repo.db.First(employeeposition); result.RowsAffected == 0 {
		logger.Error("error get data user by id not found")
		return nil, errs.NewNotFoundError("user not found")
	}
	return employeeposition, nil
}

func (repo RepositoryEmployeePositionImpl) Delete(id int64) *errs.AppErr {
	if result, err := repo.FindById(id); err != nil {
		return err
	} else {
		if resultDelete := repo.db.Delete(result); resultDelete.Error != nil {
			logger.Error("delete employee position failed" + resultDelete.Error.Error())
			return errs.NewUnexpectedError("delete employee position failed")

		} else {

			return nil
		}
	}
}
