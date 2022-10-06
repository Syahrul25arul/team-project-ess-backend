package repositoryClientImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"

	"gorm.io/gorm"
)

type repositoryClientImpl struct {
	db *gorm.DB
}

func NewRepostioryClient(db *gorm.DB) repositoryClientImpl {
	return repositoryClientImpl{db: db}
}

func (r repositoryClientImpl) Save(client *domain.Client) *errs.AppErr {
	if tx := r.db.Save(client); tx.Error != nil {
		logger.Error("error save data client: " + tx.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}

func (r repositoryClientImpl) GetAll() ([]domain.Client, *errs.AppErr) {
	// get data client and check there error or not
	var client []domain.Client
	if tx := r.db.Find(&client); tx.Error != nil {
		logger.Error("error get all data client " + tx.Error.Error())
		return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}

	return client, nil
}
