package serviceDashboardImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryUser "employeeSelfService/repository/user/impl"
	"employeeSelfService/response"
	serviceDashboard "employeeSelfService/service/dashboard"
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

func GetService() (*gorm.DB, serviceDashboard.ServiceDashboard) {
	db := database.GetClientDb()
	repositioryUser := repositoryUser.NewRepositoryUserImpl(db)
	return db, NewServiceDashboard(repositioryUser)
}

func TestNewServiceDashboard(t *testing.T) {
	// setup test
	SetupTest()
	_, service := GetService()

	reflection := reflect.TypeOf(service)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "serviceDashboardImpl")
}

func Test_serviceDashboardImpl_GetDashboard(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"employee", "users", "approval", "employee_position", "position", "absensi", "absen_configuration"})
	database.SetupDataDummy(db)

	testCase := []struct {
		name      string
		want1     string
		expected  *errs.AppErr
		expected2 *response.ResponseDashboard
	}{
		{
			name:     "get dashboard success employee no approval",
			want1:    "1",
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
			want1:    "2",
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
			name:      "get dashboard success employee with no data",
			want1:     "4",
			expected:  errs.NewNotFoundError("user not found"),
			expected2: nil,
		},
		{
			name:      "get dashboard failed id not relevant",
			want1:     "asdfadsfh",
			expected:  errs.NewBadRequestError("id not relevant"),
			expected2: nil,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result, err := service.GetDashboard(testTable.want1)
			assert.Equal(t, testTable.expected, err)
			assert.Equal(t, testTable.expected2, result)
		})
	}
}
