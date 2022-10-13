package serviceClientImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryClientImpl "employeeSelfService/repository/client/impl"
	"employeeSelfService/response"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTest() {
	config.SetupEnv("../../../.env")
	config.SanityCheck()
}

func GetService() (*gorm.DB, serviceClientImpl) {
	db := database.GetClientDb()
	repositiory := repositoryClientImpl.NewRepostioryClient(db)
	return db, NewServiceClient(repositiory)
}

func TestNewServiceClient(t *testing.T) {
	// setup test
	SetupTest()
	_, service := GetService()

	reflection := reflect.TypeOf(service)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "serviceClientImpl")
}

func Test_serviceClientImpl_SaveClient(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"client"})

	testCase := []struct {
		name      string
		want1     *domain.Client
		expected  *errs.AppErr
		expected2 *helper.SuccessResponseMessage
	}{
		{
			name: "Save Client Success",
			want1: &domain.Client{
				NamaClient:   "Indo Maret",
				Lattitude:    -6.288405,
				Longitude:    106.812327,
				AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
			},
			expected2: &helper.SuccessResponseMessage{
				Code:    http.StatusCreated,
				Status:  "Ok",
				Message: "Data client has been created",
			},
			expected: nil,
		},

		{
			name: "Save update Client Success",
			want1: &domain.Client{
				IdClient:     1,
				NamaClient:   "Indo Maret update",
				Lattitude:    -6.288405,
				Longitude:    106.812327,
				AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
			},
			expected2: &helper.SuccessResponseMessage{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "Data client has been updated",
			},
			expected: nil,
		},
		{
			name: "Save update Client Success",
			want1: &domain.Client{
				NamaClient:   "Blue Bird Group",
				Lattitude:    -6.255734,
				Longitude:    106.776826,
				AlamatClient: "Jl. Mampang Prpt. Raya No.60, RT.9/RW.3, Tegal Parang, Kec. Mampang Prpt., Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12790",
			},
			expected2: &helper.SuccessResponseMessage{
				Code:    http.StatusCreated,
				Status:  "Ok",
				Message: "Data client has been created",
			},
			expected: nil,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result, err := service.SaveClient(testTable.want1)
			assert.Equal(t, testTable.expected, err)
			assert.Equal(t, testTable.expected2, result)
		})
	}
}

func Test_serviceClientImpl_GetAllClient(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"client"})
	database.SetupDataClientDummy(db)

	testCase := []struct {
		name      string
		expected1 *response.ResponseClient
		expected2 *errs.AppErr
	}{
		{
			name: "Success Response Get All Client",
			expected1: &response.ResponseClient{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "success get data client",
			},
			expected2: nil,
		},
		{
			name: "Success Response Get All Client nil",
			expected1: &response.ResponseClient{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "success get data client",
			},
			expected2: nil,
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 1 {
				helper.TruncateTable(db, []string{"client"})
			}
			result, err := service.GetAllClient()
			testTable.expected1.Data = result.Data
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func Test_serviceClientImpl_GetClientById(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"client"})
	database.SetupDataClientDummy(db)

	testCase := []struct {
		name      string
		want      int
		expected1 *response.ResponseClient
		expected2 *errs.AppErr
	}{
		{
			name: "Success Response Get Client By Id",
			want: 1,
			expected1: &response.ResponseClient{
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
			expected2: nil,
		},
		{
			name:      "Failed Response Get Client by id not found",
			want:      5,
			expected1: nil,
			expected2: errs.NewNotFoundError("data client not found"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result, err := service.GetClientById(testTable.want)
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func Test_serviceClientImpl_DeleteClient(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"client"})
	database.SetupDataClientDummy(db)

	testCase := []struct {
		name      string
		want      int
		expected1 *helper.SuccessResponseMessage
		expected2 *errs.AppErr
	}{
		{
			name:      "Success Delete Client",
			want:      1,
			expected1: helper.NewSuccessResponseMessage(http.StatusOK, "client", "deleted"),
			expected2: nil,
		},
		{
			name:      "Failed Delete Client",
			want:      5,
			expected1: nil,
			expected2: errs.NewNotFoundError("delete failed, client not found"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result, err := service.DeleteClient(testTable.want)
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}
