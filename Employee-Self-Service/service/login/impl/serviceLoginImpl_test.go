package serviceLoginImpl

import (
	"database/sql"
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryAuth "employeeSelfService/repository/auth/impl"
	"employeeSelfService/request"
	"employeeSelfService/response"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupDataForAuth(db *gorm.DB) {
	tx := db.Begin()

	// create user
	userTest := &domain.User{
		Email:          "test@gmail.com",
		Password:       helper.BcryptPassword(config.SECRET_KEY + "password"),
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

	tx.Commit()

}

func SetupTest() {
	config.SetupEnv("../../../.env")
	config.SanityCheck()
}

func GetService() (*gorm.DB, ServiceLoginImpl) {
	db := database.GetClientDb()
	registerRepository := repositoryAuth.NewRepositoryAuthImpl(db)
	return db, NewLoginService(registerRepository)
}

func TestNewLoginService(t *testing.T) {
	// setup test
	SetupTest()
	_, service := GetService()

	reflection := reflect.TypeOf(service)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "ServiceLoginImpl")
}

func TestServiceLoginImpl_Login(t *testing.T) {
	// setup db and dummy data
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"employee", "users"})
	SetupDataForAuth(db)

	testCase := []struct {
		name      string
		want      *request.LoginRequest
		expected1 *response.ResponseLogin
		expected2 *errs.AppErr
	}{
		{
			name:      "auth login success",
			want:      &request.LoginRequest{Email: "test@gmail.com", Password: "password"},
			expected1: &response.ResponseLogin{Message: "Your Login Success", Code: http.StatusOK},
			expected2: nil,
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			resultResponse, resultErr := service.Login(testTable.want)

			// cek jika data response ada
			if resultResponse != nil {
				assert.Equal(t, testTable.expected1.Code, resultResponse.Code)
				assert.Equal(t, testTable.expected1.Message, resultResponse.Message)

				// cek jwt valid atau tiidak
				token, err := jwt.Parse(resultResponse.Token, func(token *jwt.Token) (interface{}, error) {
					return []byte(config.SECRET_KEY), nil
				})

				assert.True(t, token.Valid)
				assert.Nil(t, err)

			}

			assert.Equal(t, testTable.expected2, resultErr)

		})
	}
}
