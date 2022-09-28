package serviceEmailValidationImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/helper"
	repositoryEmailValidation "employeeSelfService/repository/emailValidation/impl"
	"employeeSelfService/response"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTest() {
	config.SetupEnv("../../../.env")
	config.SanityCheck()
}

func GetService() (*gorm.DB, ServiceEmaiValidationImpl) {
	db := database.GetClientDb()
	repositiory := repositoryEmailValidation.NewRepositoryEmailValidationImpl(db)
	return db, NewServiceEmailValidationImpl(repositiory)
}

func TestNewServiceEmailValidationImpl(t *testing.T) {
	// setup test
	SetupTest()
	_, service := GetService()

	reflection := reflect.TypeOf(service)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "ServiceEmaiValidationImpl")
}

func SetupDataForEmailValidation(db *gorm.DB) {
	tx := db.Begin()
	email1 := &domain.EmailValidation{NamaEmailValidation: "@celerates.co.id"}
	email2 := &domain.EmailValidation{NamaEmailValidation: "@celerates.com"}

	tx.Create(email1)
	tx.Create(email2)
	tx.Commit()

}

func TestServiceEmaiValidationImpl_Save(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"email_validation"})
	SetupDataForEmailValidation(db)

	testCase := []struct {
		name     string
		want     *domain.EmailValidation
		expected response.ReponseEmailValidation
	}{
		{
			name:     "Email Validation Success",
			want:     &domain.EmailValidation{NamaEmailValidation: "@gmail.com"},
			expected: response.NewResponseEmailValidationSuccess(),
		},
		{
			name:     "Email Validation Failed",
			want:     &domain.EmailValidation{NamaEmailValidation: "@celerates.co.id"},
			expected: response.NewResponseEmailValidationFailed("email for @celerates.co.id exist"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := service.Save(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
