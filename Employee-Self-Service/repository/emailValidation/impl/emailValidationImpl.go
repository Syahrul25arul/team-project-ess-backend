package repositoryEmailValidatoinImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"

	"gorm.io/gorm"
)

type RepositoryEmailValidationImpl struct {
	db *gorm.DB
}

func NewRepositoryEmailValidationImpl(client *gorm.DB) RepositoryEmailValidationImpl {
	return RepositoryEmailValidationImpl{client}
}

func (repo RepositoryEmailValidationImpl) Save(emailValidation *domain.EmailValidation) *errs.AppErr {
	// cari email terlebih dahulu apakah ada atau tidak
	if result := repo.db.First(emailValidation, "nama_email_validation = ?", emailValidation.NamaEmailValidation); result.RowsAffected > 0 {
		return errs.NewBadRequestError("email for " + emailValidation.NamaEmailValidation + " exist")
	}

	if result := repo.db.Create(emailValidation); result.Error != nil {
		logger.Error("error insert data email validationt " + result.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}
