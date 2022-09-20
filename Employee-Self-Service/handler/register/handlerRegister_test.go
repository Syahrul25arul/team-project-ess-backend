package registerHandler

import (
	"bytes"
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/helper"
	repositoryRegisterImpl "employeeSelfService/repository/register/impl"
	requestRegister "employeeSelfService/request/register"
	serviceRegisterImpl "employeeSelfService/service/register/impl"
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

func getHandler() (*gorm.DB, HandlerRegister) {
	db := database.GetClientDb()
	registerRepository := repositoryRegisterImpl.NewRepositoryRegisterImpl(db)
	registerService := serviceRegisterImpl.NewCustomerService(registerRepository)
	return db, HandlerRegister{&registerService}
}

func TestHandlerRegister_RegisterHandler(t *testing.T) {
	// setup gin
	r := gin.Default()
	SetupTest()

	// set handler product
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"employee", "users"})

	// set end point to testing
	r.POST("/register", handler.RegisterHandler)

	tests := []struct {
		name            string
		request         *requestRegister.RegisterRequest
		expectedCode    int
		expectedMessage string
	}{
		// TODO: Add test cases.
		{
			name: "register handler save data user and employee",
			request: &requestRegister.RegisterRequest{
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
			expectedCode:    http.StatusCreated,
			expectedMessage: "{code:201,message:Your Account have been created,status:ok}",
		},
	}
	for _, testTable := range tests {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.request)
			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))

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
