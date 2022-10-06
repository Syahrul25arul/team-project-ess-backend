package repositoryClientImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryClient "employeeSelfService/repository/client"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func SetupTest() {
	config.SetupEnv("../../../.env")
	config.SanityCheck()
}

func GetRepository() (*gorm.DB, repositoryClient.RepositoryClient) {
	db := database.GetClientDb()
	return db, NewRepostioryClient(db)
}

func TestNewRepostioryClient(t *testing.T) {
	SetupTest()
	_, getRepository := GetRepository()

	reflection := reflect.TypeOf(getRepository)

	assert.NotNil(t, reflection.Name())
	assert.Equal(t, reflection.Name(), "repositoryClientImpl")
}

func Test_repositoryClientImpl_Save(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"client"})

	testCase := []struct {
		name     string
		want     *domain.Client
		expected *errs.AppErr
	}{
		{
			name: "Save Client Success",
			want: &domain.Client{
				NamaClient:   "Indo Maret",
				Lattitude:    -6.288405,
				Longitude:    106.812327,
				AlamatClient: "Jl. Al Maruf No.58, RT.10/RW.3, Cilandak Tim., Kec. Ps. Minggu, KOTA ADM, Daerah Khusus Ibukota Jakarta 12140",
			},
			expected: nil,
		},
	}
	for _, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			result := repository.Save(testTable.want)
			assert.Equal(t, testTable.expected, result)
		})
	}
}

func Test_repositoryClientImpl_GetAll(t *testing.T) {
	SetupTest()
	db, repository := GetRepository()
	helper.TruncateTable(db, []string{"client"})
	database.SetupDataClientDummy(db)

	testCase := []struct {
		name      string
		expected1 []domain.Client
		expected2 *errs.AppErr
	}{
		{
			name: "Get All Client Success",
			expected1: []domain.Client{
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
			expected2: nil,
		},
		{
			name:      "Get All Client Failed",
			expected1: []domain.Client{},
			expected2: nil,
		},
	}
	for i, testTable := range testCase {
		t.Run(testTable.name, func(t *testing.T) {
			if i == 1 {
				helper.TruncateTable(db, []string{"client"})
			}
			result, err := repository.GetAll()
			assert.Equal(t, testTable.expected1, result)
			assert.Equal(t, testTable.expected2, err)
		})
	}
}
