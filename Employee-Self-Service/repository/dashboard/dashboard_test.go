package repositoryDashboard

import (
	"employeeSelfService/config"
	"employeeSelfService/database"
	"employeeSelfService/domain"
	"employeeSelfService/helper"
	"fmt"
	"testing"
)

func SetupTest() {
	config.SetupEnv("../../.env")
	config.SanityCheck()
}

func TestDashboard(t *testing.T) {
	SetupTest()
	db := database.GetClientDb()

	helper.TruncateTable(db, []string{"employee", "users", "approval", "employee_position", "position", "absensi", "absen_configuration"})
	database.SetupDataDummy(db)

	var employee *domain.Employee

	// type UserResponse struct {
	// 	IdUser   int64
	// 	Email    string
	// 	UserRole string
	// 	Employee *domain.Employee
	// }
	list := map[string]interface{}{}
	db.Table("users").Select("users.id_user,users.email,users.user_role,employee.id_employe,employee.nama_lengkap").Joins("left join employee on users.id_user = employee.id_user").Find(&list)

	fmt.Println(list)
	fmt.Println(employee)
}
