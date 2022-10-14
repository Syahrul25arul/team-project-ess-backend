package handlerClient

import (
	"bytes"
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryClient "employeeSelfService/repository/client/impl"
	"employeeSelfService/response"
	serviceClient "employeeSelfService/service/client/impl"
	"encoding/json"
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

func getHandler() (*gorm.DB, handlerClient) {
	db := database.GetClientDb()
	repositoryClient := repositoryClient.NewRepostioryClient(db)
	service := serviceClient.NewServiceClient(repositoryClient)
	return db, handlerClient{service: service}
}

func Test_handlerClient_SaveClient(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client"})

	// set end point to testing
	r.POST("/client/:user_id", handler.SaveClient)

	testCase := []struct {
		name            string
		requestBody     *domain.Client
		expectedMessage string
		expectedCode    int
	}{
		{
			name: "Save Client Success",
			requestBody: &domain.Client{
				NamaClient:   "Indo Maret",
				Lattitude:    -6.288405,
				Longitude:    106.812327,
				AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
			},
			expectedMessage: "{Code:201,Status:Ok,Message:Data client has been created}",
			expectedCode:    http.StatusCreated,
		},

		{
			name: "Save Client Success",
			requestBody: &domain.Client{
				IdClient:     1,
				NamaClient:   "Indo Maret update",
				Lattitude:    -6.288405,
				Longitude:    106.812327,
				AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
			},
			expectedMessage: "{Code:200,Status:Ok,Message:Data client has been updated}",
			expectedCode:    http.StatusOK,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.requestBody)
			req, _ := http.NewRequest("POST", "/client/:user_id", bytes.NewBuffer(jsonValue))

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

func Test_handlerClient_GetAllClient(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client"})
	database.SetupDataClientDummy(db)

	// set end point to testing
	r.GET("/client/:user_id", handler.GetAllClient)

	testCase := []struct {
		name             string
		expectedResponse *response.ResponseClient
	}{
		{
			name: "Get All Client Success",
			expectedResponse: &response.ResponseClient{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "success get data client",
				Data: []domain.Client{
					{
						IdClient:     1,
						NamaClient:   "Indo Maret",
						Lattitude:    -6.288405,
						Longitude:    106.812327,
						AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
					},
					{
						IdClient:     2,
						NamaClient:   "Blue Bird Group",
						Lattitude:    -6.255734,
						Longitude:    106.776826,
						AlamatClient: "Jl. Mampang Prpt. Raya No.60, RT.9/RW.3, Tegal Parang, Kec. Mampang Prpt., Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12790",
					},
				},
			},
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			req, _ := http.NewRequest("GET", "/client/:user_id", nil)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// get response body from handler
			resp := w.Result()
			type Data struct {
				Data []domain.Client
			}

			var GetData Data
			// read response
			responseByte, _ := io.ReadAll(resp.Body)
			var responseClient *response.ResponseClient
			json.Unmarshal(responseByte, &responseClient)
			json.Unmarshal(responseByte, &GetData)

			responseClient.Data = GetData.Data

			assert.Equal(t, testTable.expectedResponse, responseClient)
		})
	}
}

func Test_handlerClient_GetClientById(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client"})
	database.SetupDataClientDummy(db)

	// set end point to testing
	r.GET("/client/:user_id/:client_id", handler.GetClientById)

	testCase := []struct {
		name             string
		requestUrl       string
		expectedResponse interface{}
	}{
		{
			name:       "Get ClientById Success",
			requestUrl: "/client/1/1",
			expectedResponse: &response.ResponseClient{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "success get data client",
				Data: &domain.Client{
					IdClient:     1,
					NamaClient:   "Indo Maret",
					Lattitude:    -6.288405,
					Longitude:    106.812327,
					AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
				},
			},
		},
		{
			name:       "Get ClientById Success",
			requestUrl: "/client/1/2",
			expectedResponse: &response.ResponseClient{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "success get data client",
				Data: &domain.Client{
					IdClient:     2,
					NamaClient:   "Blue Bird Group",
					Lattitude:    -6.255734,
					Longitude:    106.776826,
					AlamatClient: "Jl. Mampang Prpt. Raya No.60, RT.9/RW.3, Tegal Parang, Kec. Mampang Prpt., Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12790",
				},
			},
		},
		{
			name:             "Get ClientById Not Found",
			requestUrl:       "/client/1/3",
			expectedResponse: errs.NewNotFoundError("data client not found"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			req, _ := http.NewRequest("GET", testTable.requestUrl, nil)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// get response body from handler
			resp := w.Result()
			type Data struct {
				Data *domain.Client
			}

			var GetData Data
			// read response
			responseByte, _ := io.ReadAll(resp.Body)

			if w.Code == http.StatusNotFound {
				var responseErr *errs.AppErr
				json.Unmarshal(responseByte, &responseErr)
				assert.Equal(t, testTable.expectedResponse, responseErr)
			} else {
				var responseClient *response.ResponseClient
				json.Unmarshal(responseByte, &responseClient)
				json.Unmarshal(responseByte, &GetData)

				responseClient.Data = GetData.Data

				assert.Equal(t, testTable.expectedResponse, responseClient)
			}

		})
	}
}

func Test_handlerClient_DeleteClient(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client"})
	database.SetupDataClientDummy(db)

	// set end point to testing
	r.DELETE("/client/:user_id/:client_id", handler.DeleteClient)

	testCase := []struct {
		name             string
		requestUrl       string
		expectedResponse interface{}
	}{
		{
			name:             "Get ClientById Success",
			requestUrl:       "/client/1/1",
			expectedResponse: helper.NewSuccessResponseMessage(http.StatusOK, "client", "deleted"),
		},
		{
			name:             "Get ClientById Success",
			requestUrl:       "/client/1/2",
			expectedResponse: helper.NewSuccessResponseMessage(http.StatusOK, "client", "deleted"),
		},
		{
			name:             "Get ClientById Not Found",
			requestUrl:       "/client/1/3",
			expectedResponse: errs.NewNotFoundError("delete failed, client not found"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			req, _ := http.NewRequest("DELETE", testTable.requestUrl, nil)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// get response body from handler
			resp := w.Result()

			// read response
			responseByte, _ := io.ReadAll(resp.Body)

			if w.Code == http.StatusNotFound {
				var responseErr *errs.AppErr
				json.Unmarshal(responseByte, &responseErr)
				assert.Equal(t, testTable.expectedResponse, responseErr)
			} else {
				var responseClient *helper.SuccessResponseMessage
				json.Unmarshal(responseByte, &responseClient)
				assert.Equal(t, testTable.expectedResponse, responseClient)
			}

		})
	}
}

func Test_handlerClient_UpdateClient(t *testing.T) {
	// setup
	r := gin.Default()
	SetupTest()

	// setup handler and data dummy
	db, handler := getHandler()
	helper.TruncateTable(db, []string{"client"})
	database.SetupDataClientDummy(db)

	// set end point to testing
	r.PUT("/client/:user_id", handler.UpdateClient)

	testCase := []struct {
		name            string
		requestBody     *domain.Client
		expectedMessage string
		expectedCode    int
	}{

		{
			name: "Update Client Success",
			requestBody: &domain.Client{
				IdClient:     1,
				NamaClient:   "Indo Maret update",
				Lattitude:    -6.288405,
				Longitude:    106.812327,
				AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
			},
			expectedMessage: "{Code:201,Status:Ok,Message:Data client has been updated}",
			expectedCode:    http.StatusCreated,
		},
		{
			name: "Update Not Found",
			requestBody: &domain.Client{
				IdClient:     125,
				NamaClient:   "Indo Maret update",
				Lattitude:    -6.288405,
				Longitude:    106.812327,
				AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
			},
			expectedMessage: "{code:404,message:data client not found,data:null}",
			expectedCode:    http.StatusNotFound,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			// set data request to bytes and put to NewRequest
			jsonValue, _ := json.Marshal(testTable.requestBody)
			req, _ := http.NewRequest("PUT", "/client/:user_id", bytes.NewBuffer(jsonValue))

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
