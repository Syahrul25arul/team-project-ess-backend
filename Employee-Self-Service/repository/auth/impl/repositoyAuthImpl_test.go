package repositoryUserImpl

import (
	"database/sql"
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

func GetRepository() (*gorm.DB, RepositoryAuthImpl) {
	db := database.GetClientDb()
	return db, NewRepositoryAuthImpl(db)
}

func TestNewRepositoryAuthImpl(t *testing.T) {
	SetupTest()
	_, getRepository := GetRepository()

	reflection := reflect.TypeOf(getRepository)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "RepositoryAuthImpl")
}

func SetupDataForAuth(db *gorm.DB) {
	tx := db.Begin()

	// create user
	userTest := &domain.User{
		Email:          "test@gmail.com",
		Password:       "29385789sdljkgndsjkh",
		UserRole:       "employee",
		StatusVerified: "true",
	}
	tx.Create(userTest)

	// create employee
	employeeTest := &domain.Employee{
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
	}
	employeeTest.IdUser = sql.NullInt64{Int64: int64(userTest.IdUser), Valid: true}
	tx.Create(employeeTest)

	// create position
	positionTest := &domain.Position{
		PositionName: "HRD",
	}
	tx.Create(positionTest)

	// create employeePosition
	employeePositionTest := &domain.EmployeePosition{
		IdPosition: positionTest.IdPosition,
		IdEmploye:  employeeTest.IdEmploye,
	}

	tx.Create(employeePositionTest)

	tx.Commit()

}

func TestRepositoryAuthImpl_FindByEmail(t *testing.T) {
	// setup
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"employee", "users", "position", "employee_position"})
	SetupDataForAuth(db)

	testCase := []struct {
		name      string
		want      string
		expected  *errs.AppErr
		expected2 *domain.Auth
	}{
		{
			name:     "Register success",
			want:     "test@gmail.com",
			expected: nil,
			expected2: &domain.Auth{
				Email:       "test@gmail.com",
				IdUser:      sql.NullInt64{Int64: int64(1), Valid: true},
				IdEmploye:   sql.NullInt64{Int64: int64(1), Valid: true},
				UserRole:    "employee",
				NamaLengkap: "Teddy",
			},
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			user, errors := repository.FindByEmail(testTable.want)
			assert.Equal(t, testTable.expected, errors)
			testTable.expected2.Password = user.Password
			assert.Equal(t, testTable.expected2, user)
		})
	}
}
