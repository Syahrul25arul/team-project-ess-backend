package handlerDashboard

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/helper"
	repoistoryUser "employeeSelfService/repository/user/impl"
	"employeeSelfService/response"
	serviceDashboardImpl "employeeSelfService/service/dashboard/impl"
	"encoding/json"
	"fmt"
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

func getHandler() (*gorm.DB, handlerDashboard) {
	db := database.GetClientDb()
	repositoryUser := repoistoryUser.NewRepositoryUserImpl(db)
	service := serviceDashboardImpl.NewServiceDashboard(repositoryUser)
	return db, handlerDashboard{service: service}
}

func Test_handlerDashboard_GetDashboardHandler(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"employee", "users", "approval", "employee_position", "position", "absensi", "absen_configuration"})
	database.SetupDataDummy(db)

	// set end point to testing
	r.GET("/dashboard/:user_id", handler.GetDashboardHandler)

	testCase := []struct {
		name             string
		requestUrl       string
		expectedCode     int
		expectedResponse *response.ResponseDashboard
	}{
		{
			name:         "get dashboard success employee no approval",
			requestUrl:   "/dashboard/1",
			expectedCode: http.StatusOK,
			expectedResponse: &response.ResponseDashboard{
				PicAbsensi:          false,
				IdeEmployeSecondary: false,
				Kehadiran:           int32(2),
				Approval:            int32(0),
				SudahAbsen:          true,
				Status:              "ok",
				Code:                http.StatusOK,
			},
		},

		{
			name:         "get dashboard success employee with approval",
			requestUrl:   "/dashboard/2",
			expectedCode: http.StatusOK,
			expectedResponse: &response.ResponseDashboard{
				PicAbsensi:          true,
				IdeEmployeSecondary: false,
				Kehadiran:           int32(0),
				Approval:            int32(4),
				SudahAbsen:          false,
				Status:              "ok",
				Code:                http.StatusOK,
			},
		},
		{
			name:         "get dashboard success employee with no data",
			requestUrl:   "/dashboard/4",
			expectedCode: http.StatusNotFound,
			expectedResponse: &response.ResponseDashboard{
				PicAbsensi:          false,
				IdeEmployeSecondary: false,
				Kehadiran:           int32(0),
				Approval:            int32(0),
				SudahAbsen:          false,
				Status:              "",
				Code:                http.StatusNotFound,
			},
		},
		{
			name:         "get dashboard failed id not relevant",
			requestUrl:   "/dashboard/asdfadsfh",
			expectedCode: http.StatusBadRequest,
			expectedResponse: &response.ResponseDashboard{
				PicAbsensi:          false,
				IdeEmployeSecondary: false,
				Kehadiran:           int32(0),
				Approval:            int32(0),
				SudahAbsen:          false,
				Status:              "",
				Code:                http.StatusBadRequest,
			},
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {

			req, _ := http.NewRequest("GET", testTable.requestUrl, nil)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, testTable.expectedCode, w.Code)

			// get response body from handler
			// response := w.Body.String()

			// clear double quote from response body
			// response = helper.ClearDoubleCode(response)

			var responseDashboard *response.ResponseDashboard
			json.Unmarshal(w.Body.Bytes(), &responseDashboard)
			fmt.Println("=========== RESPONSE DASHBOARD ===========")
			fmt.Println("")
			fmt.Println(responseDashboard)
			fmt.Println("")
			fmt.Println("=========== RESPONSE DASHBOARD ===========")
			// assert.Equal(t, testTable.expectedMessage, response)
		})
	}
}
