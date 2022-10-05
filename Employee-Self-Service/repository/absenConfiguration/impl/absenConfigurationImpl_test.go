package repositoryAbsenConfigurationImpl

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

func GetRepository() (*gorm.DB, RepositoryAbsenConfigurationImpl) {
	db := database.GetClientDb()
	return db, NewRepositoryAbsenConfigurationImpl(db)
}

func TestNewRepositoryAbsenConfigurationImpl(t *testing.T) {
	SetupTest()
	_, getRepository := GetRepository()

	reflection := reflect.TypeOf(getRepository)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "RepositoryAbsenConfigurationImpl")
}

func TestRepositoryAbsenConfigurationImpl_Save(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"absen_configuration"})

	testCase := []struct {
		name     string
		want     *domain.AbsenConfiguration
		expected *errs.AppErr
	}{
		{
			name: "Absen Configuration Success",
			want: &domain.AbsenConfiguration{
				DurasiJamKerja:             9,
				IntervalKeterlambatan:      15,
				BobotKeterlambatan:         0.25,
				MaksimalBobotKeterlambatan: 1,
				IdPosition:                 1,
				MinimalMasukJamKerja:       "08:00",
				MaksimalMasukJamKerja:      "10:00",
			},
			expected: nil,
		},
		{
			name: "Absen Configuration Success 2",
			want: &domain.AbsenConfiguration{
				IdAbsenConfiguration:       1,
				DurasiJamKerja:             20,
				IntervalKeterlambatan:      125,
				BobotKeterlambatan:         0.25,
				MaksimalBobotKeterlambatan: 1,
				IdPosition:                 1,
				MinimalMasukJamKerja:       "08:00",
				MaksimalMasukJamKerja:      "10:00",
			},
			expected: nil,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := repository.Save(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}

func TestRepositoryAbsenConfigurationImpl_GetData(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"absen_configuration"})
	database.SetupDataAbsenConfiguration(db)

	result, err := repository.GetData()
	expected := &domain.AbsenConfiguration{
		IdAbsenConfiguration:       int64(1),
		DurasiJamKerja:             int8(8),
		IntervalKeterlambatan:      int8(15),
		BobotKeterlambatan:         float32(0.25),
		MaksimalBobotKeterlambatan: float32(1),
		IdPosition:                 int64(1),
		MinimalMasukJamKerja:       "08:00:00",
		MaksimalMasukJamKerja:      "10:00:00",
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, result)

	// testCase := []struct {
	// 	name     string
	// 	want     *domain.AbsenConfiguration
	// 	expected *errs.AppErr
	// }{
	// 	{
	// 		name: "Absen Configuration Success",
	// 		want: &domain.AbsenConfiguration{
	// 			DurasiJamKerja:             9,
	// 			IntervalKeterlambatan:      15,
	// 			BobotKeterlambatan:         0.25,
	// 			MaksimalBobotKeterlambatan: 1,
	// 			IdPosition:                 1,
	// 			MinimalMasukJamKerja:       "08:00",
	// 			MaksimalMasukJamKerja:      "10:00",
	// 		},
	// 		expected: nil,
	// 	},
	// }
	// for _, testTable := range testCase {
	// 	t.Run(testTable.name, func(t *testing.T) {
	// 		result := repository.Save(testTable.want)
	// 		assert.Equal(t, testTable.expected, result)
	// 	})
	// }
}
