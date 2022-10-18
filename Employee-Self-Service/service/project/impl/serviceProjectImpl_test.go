package serviceProjectImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryProject "employeeSelfService/repository/project/impl"
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

func GetService() (*gorm.DB, serviceProjectImpl) {
	db := database.GetClientDb()
	repositiory := repositoryProject.NewRepositoryProject(db)
	return db, NewServiceProject(repositiory)
}

func TestNewServiceProject(t *testing.T) {
	// setup test
	SetupTest()
	_, service := GetService()

	reflection := reflect.TypeOf(service)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "serviceProjectImpl")
}

func Test_serviceProjectImpl_SaveProject(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataClientDummy(db)

	testCase := []struct {
		name      string
		want1     *domain.Project
		expected  *errs.AppErr
		expected2 *helper.SuccessResponseMessage
	}{
		{
			name: "Save Project Success",
			want1: &domain.Project{
				ProjectName: "Blue Bird Group #1",
				IdClient:    2,
			},
			expected2: &helper.SuccessResponseMessage{
				Code:    http.StatusCreated,
				Status:  "Ok",
				Message: "Data project has been created",
			},
			expected: nil,
		},
		{
			name:      "Save Project Failed unexepected",
			want1:     &domain.Project{},
			expected2: nil,
			expected:  errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result, err := service.SaveProject(testTable.want1)
			assert.Equal(t, testTable.expected, err)
			assert.Equal(t, testTable.expected2, result)
		})
	}
}

func Test_serviceProjectImpl_GetAllProject(t *testing.T) {
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	testCase := []struct {
		name      string
		expected1 *response.ResponseProject
		expected2 *errs.AppErr
	}{
		{
			name: "Success Response Get All Project",
			expected1: &response.ResponseProject{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "Get All Data Project",
			},
			expected2: nil,
		},
		{
			name: "Success Response Get All Project Data nil",
			expected1: &response.ResponseProject{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "Get All Data Project",
				Data:    nil,
			},
			expected2: nil,
		},
		{
			name:      "Failed Get All Project",
			expected1: nil,
			expected2: errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
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
				result, err := service.GetAllProject()
				assert.Equal(t, testTable.expected1, result)
				assert.Equal(t, testTable.expected2, err)
				return
			}
			result, err := service.GetAllProject()
			testTable.expected1.Data = result.Data
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func Test_serviceProjectImpl_GetById(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	testCase := []struct {
		name      string
		want      int32
		expected1 *response.ResponseProject
		expected2 *errs.AppErr
	}{
		{
			name: "Get Project By Id Success",
			want: 1,
			expected1: &response.ResponseProject{
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
			expected2: nil,
		},
		{
			name: "Get Project By Id Success2",
			want: 2,
			expected1: &response.ResponseProject{
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
			expected2: nil,
		},
		{
			name:      "Get Project By Id Failed Not Found",
			want:      20,
			expected1: nil,
			expected2: errs.NewNotFoundError("data project not found"),
		},
		{
			name:      "Get Project By Id Failed Unexpected Error",
			want:      20,
			expected1: nil,
			expected2: errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 3 {
				sql, _ := db.DB()
				sql.Close()
			}
			result, err := service.GetById(testTable.want)
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func Test_serviceProjectImpl_Update(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	testCase := []struct {
		name      string
		want1     *domain.Project
		expected  *errs.AppErr
		expected2 *helper.SuccessResponseMessage
	}{
		{
			name: "Update Project Success",
			want1: &domain.Project{
				IdProject:   3,
				ProjectName: "Blue Bird Group #1 Update",
				IdClient:    2,
			},
			expected2: &helper.SuccessResponseMessage{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "Data project has been updated",
			},
			expected: nil,
		},
		{
			name: "Update Project Failed Not Found",
			want1: &domain.Project{
				ProjectName: "Blue Bird Group #1 Update",
				IdClient:    2,
			},
			expected2: nil,
			expected:  errs.NewNotFoundError("data project not found"),
		},
		{
			name:      "Update Project Failed Unexpected",
			want1:     &domain.Project{},
			expected2: nil,
			expected:  errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 2 {
				sql, _ := db.DB()
				sql.Close()
			}
			result, err := service.Update(testTable.want1)
			assert.Equal(t, testTable.expected, err)
			assert.Equal(t, testTable.expected2, result)
		})
	}
}

func Test_serviceProjectImpl_Delete(t *testing.T) {
	// setup test
	SetupTest()
	db, service := GetService()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	testCase := []struct {
		name      string
		want1     int32
		expected  *errs.AppErr
		expected2 *helper.SuccessResponseMessage
	}{
		{
			name:  "Delete Project Success",
			want1: 1,
			expected2: &helper.SuccessResponseMessage{
				Code:    http.StatusOK,
				Status:  "Ok",
				Message: "Data project has been deleted",
			},
			expected: nil,
		},
		{
			name:      "Delete Project Failed Not Found",
			want1:     10,
			expected2: nil,
			expected:  errs.NewNotFoundError("delete failed, project not found"),
		},
		{
			name:      "Delete Project Failed Unexpected",
			want1:     2,
			expected2: nil,
			expected:  errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!"),
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 2 {
				sql, _ := db.DB()
				sql.Close()
			}
			result, err := service.Delete(testTable.want1)
			assert.Equal(t, testTable.expected, err)
			assert.Equal(t, testTable.expected2, result)
		})
	}
}
