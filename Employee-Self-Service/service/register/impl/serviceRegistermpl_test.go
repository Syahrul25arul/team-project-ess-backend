package serviceRegisterImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryRegister "employeeSelfService/repository/register/impl"
	"employeeSelfService/request"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTest() {
	config.SetupEnv("../../../.env")
	config.SanityCheck()
}

func GetService() (*gorm.DB, ServiceRegisterImpl) {
	db := database.GetClientDb()
	registerRepository := repositoryRegister.NewRepositoryRegisterImpl(db)
	return db, NewCustomerService(registerRepository)
}

func TestNewCustomerService(t *testing.T) {
	// setup test
	SetupTest()
	_, service := GetService()

	reflection := reflect.TypeOf(service)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "ServiceRegisterImpl")
}

func TestServiceRegisterImpl_Register(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"employee", "users"})

	testCase := []struct {
		name     string
		want     *request.RegisterRequest
		expected *errs.AppErr
	}{
		{
			name: "Register success",
			want: &request.RegisterRequest{
				Email:                     "test@gmail.com",
				Password:                  "29385789sdljkgndsjkh",
				NamaLengkap:               "Teddy",
				TempatLahir:               "Jakarta",
				TanggalLahir:              "13-09-1992",
				Nik:                       2389235897352,
				AlamatKtp:                 "Cilandak timur, jeruk purut",
				PendidikanTerakhir:        "Sarjana",
				NamaPendidikanTerakhir:    "USTJ",
				JurusanPendidikanTerakhir: "Teknik Informatika",
				AlamatEmailAktif:          "teddythebear@gmail.com",
				NoTlpAktif:                "967826342389",
				KontakDarurat:             "motherBear",
				NoTlpKontakDarurat:        "2938789",
				StatusEmployee:            "aktif",
				PhotoEmployee:             "teddyBear.jpg",
			},
			expected: nil,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := service.Register(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
