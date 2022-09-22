package loginHandler

import (
	"bytes"
	"database/sql"
	"employeeSelfService/config"
	"employeeSelfService/database"
	domainEmployee "employeeSelfService/domain/employee"
	domainUser "employeeSelfService/domain/user"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	reposiotryAuthImpl "employeeSelfService/repository/auth/impl"
	loginRequest "employeeSelfService/request/login"
	loginResponse "employeeSelfService/response/login"
	serviceLoginImpl "employeeSelfService/service/login/impl"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTest() {
	config.SetupEnv("../../.env")
	config.SanityCheck()
}
func SetupDataForAuth(db *gorm.DB) {
	tx := db.Begin()

	// create user
	userTest := &domainUser.User{
		Email:          "test@gmail.com",
		Password:       helper.BcryptPassword(config.SECRET_KEY + "password"),
		UserRole:       "employee",
		StatusVerified: "true",
	}
	tx.Create(userTest)

	// create employee
	employeeTest := &domainEmployee.Employee{
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

func getHandler() (*gorm.DB, HandlerLogin) {
	db := database.GetClientDb()
	repo := reposiotryAuthImpl.NewRepositoryAuthImpl(db)
	registerService := serviceLoginImpl.NewLoginService(repo)
	return db, HandlerLogin{&registerService}
}

func TestHandlerLogin_LoginHandler(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"employee", "users"})
	SetupDataForAuth(db)

	// set end point to testing
	r.POST("/login", handler.LoginHandler)

	testCase := []struct {
		name      string
		want      *loginRequest.LoginRequest
		expected1 *loginResponse.ResponseLogin
		expected2 *errs.AppErr
	}{
		{
			name:      "auth login success",
			want:      &loginRequest.LoginRequest{Email: "test@gmail.com", Password: "password"},
			expected1: &loginResponse.ResponseLogin{Message: "Your Login Success", Code: http.StatusOK},
			expected2: nil,
		},
	}

	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.want)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			fmt.Println("=== status code ====", w.Code)

			if w.Code == http.StatusOK {
				// get response body from handler
				var response *loginResponse.ResponseLogin
				body := w.Body.String()
				json.Unmarshal([]byte(body), &response)

				assert.Equal(t, testTable.expected1.Code, response.Code)
				assert.Equal(t, testTable.expected1.Message, response.Message)

				// cek jwt valid atau tiidak
				token, err := jwt.Parse(response.Token, func(token *jwt.Token) (interface{}, error) {
					return []byte(config.SECRET_KEY), nil
				})

				assert.True(t, token.Valid)
				assert.Nil(t, err)
			} else {
				response := w.Body.String()

				// clear double code
				response = helper.ClearDoubleCode(response)

				assert.Equal(t, testTable.expected1, w.Code)
				assert.Equal(t, testTable.expected2, response)
			}
		})
	}
}
