package repositoryUserImpl

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

func GetRepository() (*gorm.DB, RepositoryUserImpl) {
	db := database.GetClientDb()
	return db, NewRepositoryUserImpl(db)
}

func TestNewRepositoryUserImpl(t *testing.T) {
	SetupTest()
	_, getRepository := GetRepository()

	reflection := reflect.TypeOf(getRepository)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "RepositoryUserImpl")
}

func TestRepositoryUserImpl_FindByEmail(t *testing.T) {
	// setup
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"employee", "users"})
	userTest := &domain.User{
		Email:          "test@gmail.com",
		Password:       "29385789sdljkgndsjkh",
		UserRole:       "employee",
		StatusVerified: "true",
	}
	db.Create(userTest)

	testCase := []struct {
		name      string
		want      string
		expected  *errs.AppErr
		expected2 *domain.User
	}{
		{
			name:      "Register success",
			want:      "test@gmail.com",
			expected:  nil,
			expected2: userTest,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			user, errors := repository.FindByEmail(testTable.want)
			assert.Equal(t, testTable.expected, errors)
			assert.Equal(t, testTable.expected2, user)
		})
	}
}
