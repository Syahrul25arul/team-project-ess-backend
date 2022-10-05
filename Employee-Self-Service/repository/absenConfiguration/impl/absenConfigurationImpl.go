package repositoryAbsenConfigurationImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	"errors"

	"gorm.io/gorm"
)

type RepositoryAbsenConfigurationImpl struct {
	db *gorm.DB
}

func NewRepositoryAbsenConfigurationImpl(client *gorm.DB) RepositoryAbsenConfigurationImpl {
	return RepositoryAbsenConfigurationImpl{client}
}

func (repo RepositoryAbsenConfigurationImpl) Save(absenConfiguration *domain.AbsenConfiguration) *errs.AppErr {
	if result := repo.db.Save(absenConfiguration); result.Error != nil {
		logger.Error("error insert data absen configuration " + result.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}

func (repo RepositoryAbsenConfigurationImpl) GetData() (*domain.AbsenConfiguration, *errs.AppErr) {
	// variable untuk menampung data absen configuration
	var absenConfiguration *domain.AbsenConfiguration

	if result := repo.db.First(&absenConfiguration); result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logger.Error("error get data absen configuration " + result.Error.Error())
		return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	} else if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else {
		return absenConfiguration, nil
	}

}
