package repositoryUserImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
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

func TestRepositoryUserImpl_FindById(t *testing.T) {
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
		want      int64
		expected  *errs.AppErr
		expected2 *domain.User
	}{
		{
			name:      "Register success",
			want:      int64(1),
			expected:  nil,
			expected2: userTest,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			user, errors := repository.FindById(testTable.want)
			assert.Equal(t, testTable.expected, errors)
			assert.Equal(t, testTable.expected2, user)
		})
	}
}

func TestRepositoryUserImpl_GetDataDashboard(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()

	helper.TruncateTable(db, []string{"employee", "users", "approval", "employee_position", "position", "absensi", "absen_configuration"})
	database.SetupDataDummy(db)

	testCase := []struct {
		name      string
		want      string
		expected  *errs.AppErr
		expected2 *response.ResponseDashboard
	}{
		{
			name:     "get dashboard success employee no approval",
			want:     "1",
			expected: nil,
			expected2: &response.ResponseDashboard{
				PicAbsensi:          false,
				IdeEmployeSecondary: false,
				Kehadiran:           int32(2),
				Approval:            int32(0),
				SudahAbsen:          true,
				Status:              "ok",
				Code:                http.StatusOK,
			},
		},
		{
			name:     "get dashboard success employee with approval",
			want:     "2",
			expected: nil,
			expected2: &response.ResponseDashboard{
				PicAbsensi:          true,
				IdeEmployeSecondary: false,
				Kehadiran:           int32(0),
				Approval:            int32(4),
				SudahAbsen:          false,
				Status:              "ok",
				Code:                http.StatusOK,
			},
		},
		{
			name:     "get dashboard success employee with no data",
			want:     "3",
			expected: nil,
			expected2: &response.ResponseDashboard{
				PicAbsensi:          false,
				IdeEmployeSecondary: false,
				Kehadiran:           int32(0),
				Approval:            int32(0),
				SudahAbsen:          false,
				Status:              "ok",
				Code:                200,
			},
		},
		{
			name:      "get dashboard failed",
			want:      "asdfadsfh",
			expected:  errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
			expected2: nil,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			user, errors := repository.GetDataDashboard(testTable.want)
			assert.Equal(t, testTable.expected, errors)
			assert.Equal(t, testTable.expected2, user)
		})
	}
}
