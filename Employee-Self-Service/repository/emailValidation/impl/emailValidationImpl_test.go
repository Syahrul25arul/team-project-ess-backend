package repositoryEmailValidatoinImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTest() {
	config.SetupEnv("../../../.env")
	config.SanityCheck()
}

func GetRepository() (*gorm.DB, RepositoryEmailValidationImpl) {
	db := database.GetClientDb()
	return db, NewRepositoryEmailValidationImpl(db)
}

func TestNewRepositoryEmailValidationImpl(t *testing.T) {
	SetupTest()
	_, getRepository := GetRepository()

	reflection := reflect.TypeOf(getRepository)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "RepositoryEmailValidationImpl")
}

func SetupDataForEmailValidation(db *gorm.DB) {
	tx := db.Begin()
	email1 := &domain.EmailValidation{NamaEmailValidation: "@celerates.co.id"}
	email2 := &domain.EmailValidation{NamaEmailValidation: "@celerates.com"}

	tx.Create(email1)
	tx.Create(email2)
	tx.Commit()

}

func TestRepositoryEmailValidationImpl_Save(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"email_validation"})
	SetupDataForEmailValidation(db)

	testCase := []struct {
		name     string
		want     *domain.EmailValidation
		expected *errs.AppErr
	}{
		{
			name:     "Email Validation Success",
			want:     &domain.EmailValidation{NamaEmailValidation: "@gmail.com"},
			expected: nil,
		},
		{
			name:     "Email Validation Failed",
			want:     &domain.EmailValidation{NamaEmailValidation: "@celerates.co.id"},
			expected: errs.NewBadRequestError("email for @celerates.co.id exist"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := repository.Save(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
