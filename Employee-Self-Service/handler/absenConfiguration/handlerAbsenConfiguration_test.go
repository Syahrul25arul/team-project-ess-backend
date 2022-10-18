package handlerAbsenConfiguration

import (
	"bytes"
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/helper"
	repoAbsenConfiguration "employeeSelfService/repository/absenConfiguration/impl"
	repoUser "employeeSelfService/repository/user/impl"
	"employeeSelfService/request"
	serviceAbsenConfigurationImpl "employeeSelfService/service/absenConfiguration/impl"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTest() {
	config.SetupEnv("../../.env")
	config.SanityCheck()
}

func getHandler() (*gorm.DB, *HandlerAbsenConfiguration) {
	db := database.GetClientDb()
	repoAbsenConfiguration := repoAbsenConfiguration.NewRepositoryAbsenConfigurationImpl(db)
	repositoryUser := repoUser.NewRepositoryUserImpl(db)
	service := serviceAbsenConfigurationImpl.NewServiceEmailValidationImpl(repoAbsenConfiguration, repositoryUser)
	return db, &HandlerAbsenConfiguration{service: service}
}

func TestHandlerAbsenConfiguration_SaveAbsenConfiguration(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()

	helper.TruncateTable(db, []string{"absen_configuration", "users"})
	database.SetupDataUserDummy(db)

	// set end point to testing
	r.POST("/konfigurasi/:user_id/kehadiran", handler.SaveAbsenConfiguration)

	testCase := []struct {
		name            string
		requestBody     *request.AbsensiConfiguration
		requestUrl      string
		expectedMessage string
		expectedCode    int
	}{
		{
			name:            "Absen Configuration Failed insert",
			requestBody:     &request.AbsensiConfiguration{},
			requestUrl:      "/konfigurasi/1/kehadiran",
			expectedMessage: "{code:500,message:Sorry, an error has occurred on our system due to an internal server error. please try again!,status:error}",
			expectedCode:    http.StatusInternalServerError,
		},
		{
			name: "Absen Configuration Success insert",
			requestBody: &request.AbsensiConfiguration{
				DurasiJamKerja:             8,
				IntervalKeterlambatan:      15,
				BobotKeterlambatan:         0.25,
				MaksimalBobotKeterlambatan: 1,
				IdPosition:                 1,
				MinimalMasukJamKerja:       "08:00",
				MaksimalMasukJamKerja:      "10:00",
			},
			requestUrl:      "/konfigurasi/1/kehadiran",
			expectedMessage: "{code:201,message:Absen configuration has been created,status:ok}",
			expectedCode:    http.StatusCreated,
		},
		{
			name:            "Absen Configuration Update Failed",
			requestBody:     &request.AbsensiConfiguration{},
			requestUrl:      "/konfigurasi/1/kehadiran",
			expectedMessage: "{code:500,message:Sorry, an error has occurred on our system due to an internal server error. please try again!,status:error}",
			expectedCode:    http.StatusInternalServerError,
		},
		{
			name: "Absen Configuration Update Success",
			requestBody: &request.AbsensiConfiguration{
				DurasiJamKerja:             8,
				IntervalKeterlambatan:      15,
				BobotKeterlambatan:         0.25,
				MaksimalBobotKeterlambatan: 1,
				IdPosition:                 1,
				MinimalMasukJamKerja:       "08:00",
				MaksimalMasukJamKerja:      "10:00",
			},
			requestUrl:      "/konfigurasi/1/kehadiran",
			expectedMessage: "{code:201,message:Absen configuration has been created,status:ok}",
			expectedCode:    http.StatusCreated,
		},
		{
			name: "Absen Configuration Failed unexpected get data",
			requestBody: &request.AbsensiConfiguration{
				DurasiJamKerja:             8,
				IntervalKeterlambatan:      15,
				BobotKeterlambatan:         0.25,
				MaksimalBobotKeterlambatan: 1,
				IdPosition:                 1,
				MinimalMasukJamKerja:       "08:00",
				MaksimalMasukJamKerja:      "10:00",
			},
			requestUrl:      "/konfigurasi/2/kehadiran",
			expectedMessage: "{code:500,message:Sorry, an error has occurred on our system due to an internal server error. please try again!,status:error}",
			expectedCode:    http.StatusInternalServerError,
		},
	}
	for i, testTable := range testCase {

		t.Run(testTable.name, func(t *testing.T) {
			if i == 4 {
				sql, _ := db.DB()
				sql.Close()
			}
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.requestBody)
			req, _ := http.NewRequest("POST", testTable.requestUrl, bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, testTable.expectedCode, w.Code)

			// get response body from handler
			response := w.Body.String()

			// clear double quote from response body
			response = helper.ClearDoubleCode(response)

			assert.Equal(t, testTable.expectedMessage, response)
		})
	}
}
