package serviceEmailValidationImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/helper"
	repositoryEmailValidation "employeeSelfService/repository/emailValidation/impl"
	repositoryUser "employeeSelfService/repository/user/impl"
	"employeeSelfService/response"
	"net/http"
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
	repositioryUser := repositoryUser.NewRepositoryUserImpl(db)
	return db, NewServiceEmailValidationImpl(repositiory, repositioryUser)
}

func TestNewServiceEmailValidationImpl(t *testing.T) {
	// setup test
	SetupTest()
	_, service := GetService()

	reflection := reflect.TypeOf(service)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "ServiceEmaiValidationImpl")
}

func TestServiceEmaiValidationImpl_Save(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"email_validation", "users"})
	database.SetupDataEmailValidationDummy(db)
	database.SetupDataUserDummy(db)

	testCase := []struct {
		name     string
		want1    *domain.EmailValidation
		want2    int
		expected response.ReponseEmailValidation
	}{
		{
			name:     "Email Validation Success",
			want1:    &domain.EmailValidation{NamaEmailValidation: "@gmail.com"},
			want2:    1,
			expected: response.NewResponseEmailValidationSuccess(),
		},
		{
			name:     "Email Validation Failed email exist",
			want1:    &domain.EmailValidation{NamaEmailValidation: "@celerates.co.id"},
			want2:    1,
			expected: response.NewResponseEmailValidationFailed(http.StatusBadRequest, "email for @celerates.co.id exist"),
		},
		{
			name:     "Email Validation Failed forbidden",
			want1:    &domain.EmailValidation{NamaEmailValidation: "@tai.co.id"},
			want2:    2,
			expected: response.NewResponseEmailValidationFailed(http.StatusInternalServerError, "Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 2 {
				sql, _ := db.DB()
				sql.Close()
			}
			result := service.Save(testTable.want1, int64(testTable.want2))
			assert.Equal(t, testTable.expected, result)
		})
	}
}
