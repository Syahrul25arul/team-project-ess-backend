package emailValidationHandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/helper"
	repositoryEmailValidation "employeeSelfService/repository/emailValidation/impl"
	repositoryUser "employeeSelfService/repository/user/impl"
	serviceEmailValidation "employeeSelfService/service/emailValidation/impl"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTest() {
	config.SetupEnv("../../.env")
	config.SanityCheck()
}

func getHandler() (*gorm.DB, HandlerEmailValidation) {
	db := database.GetClientDb()
	repo := repositoryEmailValidation.NewRepositoryEmailValidationImpl(db)
	repoUser := repositoryUser.NewRepositoryUserImpl(db)
	registerService := serviceEmailValidation.NewServiceEmailValidationImpl(repo, repoUser)
	return db, HandlerEmailValidation{&registerService}
}

func TestHandlerEmailValidation_SaveEmailValidation(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"email_validation", "users"})
	database.SetupDataEmailValidationDummy(db)
	database.SetupDataUserDummy(db)

	// set end point to testing
	r.POST("/konfigurasi/:user_id/email", handler.SaveEmailValidation)

	tests := []struct {
		name            string
		requestBody     *domain.EmailValidation
		requestUrl      string
		expectedMessage string
		expectedCode    int
	}{
		// TODO: Add test cases.
		{
			name:            "register handler save data Email Validation Success",
			requestBody:     &domain.EmailValidation{NamaEmailValidation: "@gmail.com"},
			requestUrl:      "/konfigurasi/1/email",
			expectedMessage: "{code:201,message:Email for validation register has been created,status:ok}",
			expectedCode:    201,
		},
		{
			name:            "register handler save data Email Validation Failed email exist",
			requestBody:     &domain.EmailValidation{NamaEmailValidation: "@celerates.co.id"},
			requestUrl:      "/konfigurasi/1/email",
			expectedMessage: "{code:400,message:email for @celerates.co.id exist,status:error}",
			expectedCode:    400,
		},
		{
			name:            "register handler save data Email Validation Failed forbidden",
			requestBody:     &domain.EmailValidation{NamaEmailValidation: "@tai.co.id"},
			requestUrl:      "/konfigurasi/2/email",
			expectedMessage: "{code:403,message:you dont have credential,status:error}",
			expectedCode:    403,
		},
		{
			name:            "register handler save data Email Validation Failed user not foyund",
			requestBody:     &domain.EmailValidation{NamaEmailValidation: "@tai.co.id"},
			requestUrl:      "/konfigurasi/5/email",
			expectedMessage: "{code:404,message:user not found,status:error}",
			expectedCode:    404,
		},
	}
	for _, testTable := range tests {
		t.Run(testTable.name, func(t *testing.T) {
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
