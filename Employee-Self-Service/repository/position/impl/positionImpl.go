package repositoryPositionImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"

	"gorm.io/gorm"
)

type RepositoryPositionImpl struct {
	db *gorm.DB
}

func NewRepositoryPositionImpl(client *gorm.DB) RepositoryPositionImpl {
	return RepositoryPositionImpl{client}
}

func (repo RepositoryPositionImpl) Save(position *domain.Position) *errs.AppErr {
	if result := repo.db.Create(position); result.Error != nil {
		logger.Error("error insert data position" + result.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}

func (repo RepositoryPositionImpl) FindById(id int64) (*domain.Position, *errs.AppErr) {
	var position *domain.Position = &domain.Position{IdPosition: id}
	if result := repo.db.First(position); result.RowsAffected == 0 {
		logger.Error("error get data position by id not found")
		return nil, errs.NewNotFoundError("position not found")
	}
	return position, nil
}

func (repo RepositoryPositionImpl) Delete(id int64) (*domain.Position, *errs.AppErr) {
	var err error
	var position *domain.Position = &domain.Position{IdPosition: id}
	err = repo.db.Where("id_position = ?", id).Delete(&position).Error
	if err != nil {
		logger.Error("error fetch data to inventory table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	} else {
		return position, nil
	}
}

func (repo RepositoryPositionImpl) Update(position domain.Position) *errs.AppErr {

	if result := repo.db.Save(&position); result.Error != nil {
		logger.Error("error update data position " + result.Error.Error())
		return errs.NewUnexpectedError("error update data position")
	}
	return nil
}
