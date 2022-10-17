package repositoryProjectImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryProject "employeeSelfService/repository/project"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTest() {
	config.SetupEnv("../../../.env")
	config.SanityCheck()
}

func GetRepository() (*gorm.DB, repositoryProject.RepositoryProject) {
	db := database.GetClientDb()
	return db, NewRepositoryProject(db)
}

func TestNewRepositoryProject(t *testing.T) {
	SetupTest()
	_, getRepository := GetRepository()

	reflection := reflect.TypeOf(getRepository)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "repositoryProjectImpl")
}

func Test_repositoryProjectImpl_SaveProject(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"project"})

	testCase := []struct {
		name     string
		want     *domain.Project
		expected *errs.AppErr
	}{
		{
			name: "Save Project Success",
			want: &domain.Project{
				ProjectName: "Blue Bird Group #1",
				IdClient:    2,
			},
			expected: nil,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := repository.SaveProject(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}

func Test_repositoryProjectImpl_GetAllProject(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	testCase := []struct {
		name      string
		expected1 []domain.ProjectWithClient
		expected2 *errs.AppErr
	}{
		{
			name: "Get All Project Success",
			expected1: []domain.ProjectWithClient{
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
			expected2: nil,
		},
		{
			name:      "Get All Client Failed",
			expected1: []domain.ProjectWithClient(nil),
			expected2: nil,
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 1 {
				helper.TruncateTable(db, []string{"client"})
			}
			result, err := repository.GetAllProject()
			fmt.Println(result)
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func Test_repositoryProjectImpl_GetById(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	testCase := []struct {
		name      string
		want      int32
		expected1 *domain.ProjectWithClient
		expected2 *errs.AppErr
	}{
		{
			name: "Get Project By Id Success",
			want: 1,
			expected1: &domain.ProjectWithClient{
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
			expected2: nil,
		},
		{
			name: "Get Client By Id Success2",
			want: 2,
			expected1: &domain.ProjectWithClient{
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
			expected2: nil,
		},
		{
			name:      "Get Client By Id Failed Not Found",
			want:      20,
			expected1: nil,
			expected2: errs.NewNotFoundError("data client not found"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {

			result, err := repository.GetById(testTable.want)
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}

func Test_repositoryProjectImpl_Delete(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"client", "project"})
	database.SetupDataProjectDummy(db)

	testCase := []struct {
		name     string
		want     int32
		expected *errs.AppErr
	}{
		{
			name:     "Delete Delete Project Success",
			want:     1,
			expected: nil,
		},
		{
			name:     "Delete Project Success2",
			want:     2,
			expected: nil,
		},
		{
			name:     "Delete Project Failed",
			want:     20,
			expected: errs.NewNotFoundError("delete failed, project not found"),
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			err := repository.Delete(testTable.want)
			assert.Equal(t, testTable.expected, err)
		})
	}
}
