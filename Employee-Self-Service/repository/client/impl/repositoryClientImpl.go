package repositoryClientImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	"errors"

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
func (r repositoryClientImpl) GetById(id int) (*domain.Client, *errs.AppErr) {
	// create variable for domain client
	var client *domain.Client

	// get data client by id and check there error or not
	if tx := r.db.First(&client, id); tx.Error != nil {

		// check if data client by id not found
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

			// create logger error for debugging
			logger.Error("error get data client by id : " + tx.Error.Error())
			return nil, errs.NewNotFoundError("data client not found")
		} else {

			// this block for handle error unexpected
			// create logger error for debugging
			logger.Error("error get data client by id : " + tx.Error.Error())
			return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
		}
	}
	return client, nil
}
func (r repositoryClientImpl) Delete(id int) *errs.AppErr {
	// create variable for domain client
	var client *domain.Client = &domain.Client{IdClient: int32(id)}

	// delete client from database where id from request client
	if tx := r.db.Delete(client); tx.RowsAffected < int64(1) {

		// create looger error for debuggin delete failed because data client by id nof found
		logger.Error("error delete client, id not found ")
		return errs.NewNotFoundError("delete failed, client not found")
	} else if tx.Error != nil {

		// create logger error unexpected for debugging and return
		logger.Error("error delete client unexpected " + tx.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}

	return nil
}

func (r repositoryClientImpl) GetAllWithProject() ([]domain.ClientWithProject, *errs.AppErr) {
	// get data client and check there error or not
	var clients []domain.ClientWithProject
	if tx := r.db.Preload("Projects").Find(&clients); tx.Error != nil {
		logger.Error("error get all data client " + tx.Error.Error())
		return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return clients, nil
}
