package repositoryRegisterImpl

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

func GetRepository() (*gorm.DB, RepositoryRegisterImpl) {
	db := database.GetClientDb()
	return db, NewRepositoryRegisterImpl(db)
}

func TestNewRepositoryRegisterImpl(t *testing.T) {
	SetupTest()
	_, getRepository := GetRepository()

	reflection := reflect.TypeOf(getRepository)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "RepositoryRegisterImpl")
}

func TestRepositoryRegisterImpl_Register(t *testing.T) {
	// setup
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"employee", "users"})

	testCase := []struct {
		name     string
		want1    *domain.User
		want2    *domain.Employee
		expected *errs.AppErr
	}{
		{
			name: "Register success",
			want1: &domain.User{
				Email:          "test@gmail.com",
				Password:       "29385789sdljkgndsjkh",
				StatusVerified: "true",
				UserRole:       "employee",
			},
			want2: &domain.Employee{
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
			result := repository.Register(testTable.want1, testTable.want2)
			assert.Equal(t, testTable.expected, result)
		})
	}
}
