package handlerProject

import (
	"bytes"
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryProjectImpl "employeeSelfService/repository/project/impl"
	"employeeSelfService/response"
	serviceProjectImpl "employeeSelfService/service/project/impl"
	"encoding/json"
	"fmt"
	"io"
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

func getHandler() (*gorm.DB, handlerProject) {
	db := database.GetClientDb()
	repositoryProject := repositoryProjectImpl.NewRepositoryProject(db)
	service := serviceProjectImpl.NewServiceProject(repositoryProject)
	return db, handlerProject{service: service}
}

func Test_handlerProject_SaveProject(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataClientDummy(db)

	// set end point to testing
	r.POST("/project/:user_id", handler.SaveProject)

	testCase := []struct {
		name             string
		requestBody      *domain.Project
		expectedResponse interface{}
	}{
		{
			name: "Save Project Success",
			requestBody: &domain.Project{
				ProjectName: "Blue Bird Group #1",
				IdClient:    2,
			},
			expectedResponse: &helper.SuccessResponseMessage{
				Code:    http.StatusCreated,
				Status:  "Ok",
				Message: "Data project has been created",
			},
		},

		{
			name:             "Save Project Failed unexepected",
			requestBody:      &domain.Project{},
			expectedResponse: errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.requestBody)
			req, _ := http.NewRequest("POST", "/project/:user_id", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// get response body from handler
			response := w.Result()

			// read result from response body
			responseRead, _ := io.ReadAll(response.Body)

			// create varaibel struct to get response body
			// check response code >= 400 or not
			var resultResponse interface{}
			if w.Code >= 400 {
				resultResponse = &errs.AppErr{}
			} else {
				resultResponse = &helper.SuccessResponseMessage{}
			}
			json.Unmarshal(responseRead, resultResponse)

			assert.Equal(t, testTable.expectedResponse, resultResponse)
		})
	}
}

func Test_handlerProject_UpdateProject(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	// set end point to testing
	r.PUT("/project/:user_id", handler.UpdateProject)

	testCase := []struct {
		name             string
		requestBody      *domain.Project
		expectedResponse interface{}
	}{
		{
			name: "Update Project Success",
			requestBody: &domain.Project{
				IdProject:   3,
				ProjectName: "Blue Bird Group #1 Update",
				IdClient:    2,
			},
			expectedResponse: &helper.SuccessResponseMessage{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "Data project has been updated",
			},
		},
		{
			name: "Update Project Failed Not Found",
			requestBody: &domain.Project{
				ProjectName: "Blue Bird Group #1 Update",
				IdClient:    2,
			},
			expectedResponse: errs.NewNotFoundError("data project not found"),
		},
		{
			name:             "Update Project Failed Unexpected",
			requestBody:      &domain.Project{},
			expectedResponse: errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 2 {
				sql, _ := db.DB()
				sql.Close()
			}
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.requestBody)
			req, _ := http.NewRequest("PUT", "/project/:user_id", bytes.NewBuffer(jsonValue))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// get response body from handler
			response := w.Result()

			// read result from response body
			responseRead, _ := io.ReadAll(response.Body)

			// create varaibel struct to get response body
			// check response code >= 400 or not
			var resultResponse interface{}
			if w.Code >= 400 {
				resultResponse = &errs.AppErr{}
			} else {
				resultResponse = &helper.SuccessResponseMessage{}
			}
			json.Unmarshal(responseRead, resultResponse)

			assert.Equal(t, testTable.expectedResponse, resultResponse)
		})
	}
}

func Test_handlerProject_GetAllProject(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	// set end point to testing
	r.GET("/project", handler.GetAllProject)

	testCase := []struct {
		name             string
		expectedResponse interface{}
	}{
		{
			name: "Success Response Get All Project",
			expectedResponse: &response.ResponseProject{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "Get All Data Project",
				Data: []domain.ProjectWithClient{
					{
						IdProject:   1,
						IdClient:    1,
						ProjectName: "Indo Maret #1",
						Client: domain.Client{
							IdClient:     1,
							NamaClient:   "Indo Maret",
							Lattitude:    -6.288405,
							Longitude:    106.812327,
							AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
						},
					},
					{
						IdProject:   2,
						IdClient:    1,
						ProjectName: "Indo Maret #2",
						Client: domain.Client{
							IdClient:     1,
							NamaClient:   "Indo Maret",
							Lattitude:    -6.288405,
							Longitude:    106.812327,
							AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
						},
					},
					{
						IdProject:   3,
						IdClient:    2,
						ProjectName: "Blue Bird Group #1",
						Client: domain.Client{
							IdClient:     2,
							NamaClient:   "Blue Bird Group",
							Lattitude:    -6.255734,
							Longitude:    106.776826,
							AlamatClient: "Jl. Mampang Prpt. Raya No.60, RT.9/RW.3, Tegal Parang, Kec. Mampang Prpt., Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12790",
						},
					},
					{
						IdProject:   4,
						IdClient:    2,
						ProjectName: "Blue Bird Group #2",
						Client: domain.Client{
							IdClient:     2,
							NamaClient:   "Blue Bird Group",
							Lattitude:    -6.255734,
							Longitude:    106.776826,
							AlamatClient: "Jl. Mampang Prpt. Raya No.60, RT.9/RW.3, Tegal Parang, Kec. Mampang Prpt., Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12790",
						},
					},
				},
			},
		},
		{
			name: "Success Response Get All Project Data nil",
			expectedResponse: &response.ResponseProject{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "Get All Data Project",
				Data:    []domain.ProjectWithClient(nil),
			},
		},
		{
			name:             "Failed Get All Project",
			expectedResponse: errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 1 {
				helper.TruncateTable(db, []string{"client", "project"})
			}
			if i == 2 {
				sql, _ := db.DB()
				sql.Close()
			}
			// set data request to bytes and put to NewRequest
			req, _ := http.NewRequest("GET", "/project", nil)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// get response body from handler
			resp := w.Result()
			type Data struct {
				Data []domain.ProjectWithClient
			}

			// read response
			responseByte, _ := io.ReadAll(resp.Body)

			if w.Code >= 400 {
				var responseProject *errs.AppErr
				json.Unmarshal(responseByte, &responseProject)
				assert.Equal(t, testTable.expectedResponse, responseProject)
			} else {
				var responseProject *response.ResponseProject
				var GetData Data

				json.Unmarshal(responseByte, &responseProject)
				json.Unmarshal(responseByte, &GetData)
				responseProject.Data = GetData.Data
				fmt.Println("======== DATA ==========")
				fmt.Println(responseProject)
				assert.Equal(t, testTable.expectedResponse, responseProject)
			}

		})
	}
}

func Test_handlerProject_GetProjectById(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	// set end point to testing
	r.GET("/project/:project_id", handler.GetProjectById)

	testCase := []struct {
		name             string
		url              string
		expectedResponse interface{}
	}{
		{
			name: "Get Project By Id Success",
			url:  "/project/1",
			expectedResponse: &response.ResponseProject{
				Code:    http.StatusOK,
				Message: "Get Data Project By Id",
				Status:  "Ok",
				Data: &domain.ProjectWithClient{
					IdProject:   1,
					IdClient:    1,
					ProjectName: "Indo Maret #1",
					Client: domain.Client{
						IdClient:     1,
						NamaClient:   "Indo Maret",
						Lattitude:    -6.288405,
						Longitude:    106.812327,
						AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
					},
				},
			},
		},
		{
			name: "Get Project By Id Success2",
			url:  "/project/2",
			expectedResponse: &response.ResponseProject{
				Code:    http.StatusOK,
				Message: "Get Data Project By Id",
				Status:  "Ok",
				Data: &domain.ProjectWithClient{
					IdProject:   2,
					IdClient:    1,
					ProjectName: "Indo Maret #2",
					Client: domain.Client{
						IdClient:     1,
						NamaClient:   "Indo Maret",
						Lattitude:    -6.288405,
						Longitude:    106.812327,
						AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
					},
				},
			},
		},
		{
			name:             "Get Project By Id Failed Not Found",
			url:              "/project/20",
			expectedResponse: errs.NewNotFoundError("data project not found"),
		},
		{
			name:             "Get Project By Id Failed Unexpected Error",
			url:              "/project/20",
			expectedResponse: errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 2 {
				helper.TruncateTable(db, []string{"client", "project"})
			}
			if i == 3 {
				sql, _ := db.DB()
				sql.Close()
			}
			// set data request to bytes and put to NewRequest
			req, _ := http.NewRequest("GET", testTable.url, nil)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// get response body from handler
			resp := w.Result()
			type Data struct {
				Data *domain.ProjectWithClient
			}

			// read response
			responseByte, _ := io.ReadAll(resp.Body)

			if w.Code >= 400 {
				var responseProject *errs.AppErr
				json.Unmarshal(responseByte, &responseProject)
				assert.Equal(t, testTable.expectedResponse, responseProject)
			} else {
				var responseProject *response.ResponseProject
				var GetData Data

				json.Unmarshal(responseByte, &responseProject)
				json.Unmarshal(responseByte, &GetData)
				responseProject.Data = GetData.Data

				assert.Equal(t, testTable.expectedResponse, responseProject)
			}

		})
	}
}

func Test_handlerProject_DeleteProject(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	// set end point to testing
	r.DELETE("/project/:project_id", handler.DeleteProject)

	testCase := []struct {
		name             string
		requestUrl       string
		expectedResponse interface{}
	}{
		{
			name:       "Delete Project Success",
			requestUrl: "/project/1",
			expectedResponse: &helper.SuccessResponseMessage{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "Data project has been deleted",
			},
		},
		{
			name:             "Delete Project Failed Not Found",
			requestUrl:       "/project/10",
			expectedResponse: errs.NewNotFoundError("delete failed, project not found"),
		},
		{
			name:             "Delete Project Failed Unexpected",
			requestUrl:       "/project/2",
			expectedResponse: errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 1 {
				helper.TruncateTable(db, []string{"client", "project"})
			}
			if i == 2 {
				sql, _ := db.DB()
				sql.Close()
			}
			// set data request to bytes and put to NewRequest
			req, _ := http.NewRequest("DELETE", testTable.requestUrl, nil)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// get response body from handler
			resp := w.Result()

			// read response
			responseByte, _ := io.ReadAll(resp.Body)

			if w.Code >= 400 {
				var responseProject *errs.AppErr
				json.Unmarshal(responseByte, &responseProject)
				assert.Equal(t, testTable.expectedResponse, responseProject)
			} else {
				var responseProject *helper.SuccessResponseMessage
				json.Unmarshal(responseByte, &responseProject)
				assert.Equal(t, testTable.expectedResponse, responseProject)
			}

		})
	}
}
