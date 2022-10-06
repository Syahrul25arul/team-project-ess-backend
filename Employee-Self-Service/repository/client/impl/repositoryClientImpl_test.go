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
	tests := []struct {
		name  string
		r     repositoryClientImpl
		want  []domain.Client
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.GetAll()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repositoryClientImpl.GetAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("repositoryClientImpl.GetAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
