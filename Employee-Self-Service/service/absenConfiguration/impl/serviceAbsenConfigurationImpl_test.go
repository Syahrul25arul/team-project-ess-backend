package serviceAbsenConfigurationImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/helper"
	repositoryAbsenConfiguration "employeeSelfService/repository/absenConfiguration/impl"
	repositoryUser "employeeSelfService/repository/user/impl"
	"employeeSelfService/request"
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

func GetService() (*gorm.DB, ServiceAbsenConfigurationImpl) {
	db := database.GetClientDb()
	repositiory := repositoryAbsenConfiguration.NewRepositoryAbsenConfigurationImpl(db)
	repositioryUser := repositoryUser.NewRepositoryUserImpl(db)
	return db, NewServiceEmailValidationImpl(repositiory, repositioryUser)
}

func TestNewServiceEmailValidationImpl(t *testing.T) {
	// setup test
	SetupTest()
	_, service := GetService()

	reflection := reflect.TypeOf(service)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "ServiceAbsenConfigurationImpl")
}

func TestServiceAbsenConfigurationImpl_Save(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"absen_configuration", "users"})
	database.SetupDataUserDummy(db)

	testCase := []struct {
		name     string
		want1    *request.AbsensiConfiguration
		want2    int
		expected response.ResponseAbsenConfiguration
	}{
		{
			name: "Email Validation Success",
			want1: &request.AbsensiConfiguration{
				DurasiJamKerja:             8,
				IntervalKeterlambatan:      15,
				BobotKeterlambatan:         0.25,
				MaksimalBobotKeterlambatan: 1,
				IdPosition:                 1,
				MinimalMasukJamKerja:       "08:00",
				MaksimalMasukJamKerja:      "10:00",
			},
			want2:    1,
			expected: response.NewResponseAbsenConfiguration(),
		},

		{
			name: "Email Validation Failed forbidden",
			want1: &request.AbsensiConfiguration{
				DurasiJamKerja:             8,
				IntervalKeterlambatan:      15,
				BobotKeterlambatan:         0.25,
				MaksimalBobotKeterlambatan: 1,
				IdPosition:                 1,
				MinimalMasukJamKerja:       "08:00",
				MaksimalMasukJamKerja:      "10:00",
			},
			want2:    2,
			expected: response.NewResponseAbsenConfigurationFailed(http.StatusForbidden, "you dont have credential"),
		},
		{
			name: "Email Validation Failed user not foyund",
			want1: &request.AbsensiConfiguration{
				DurasiJamKerja:             8,
				IntervalKeterlambatan:      15,
				BobotKeterlambatan:         0.25,
				MaksimalBobotKeterlambatan: 1,
				IdPosition:                 1,
				MinimalMasukJamKerja:       "08:00",
				MaksimalMasukJamKerja:      "10:00",
			},
			want2:    5,
			expected: response.NewResponseAbsenConfigurationFailed(http.StatusNotFound, "user not found"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := service.Save(testTable.want1, int64(testTable.want2))
			assert.Equal(t, testTable.expected, result)
		})
	}
}
