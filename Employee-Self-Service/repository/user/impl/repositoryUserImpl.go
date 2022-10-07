package repositoryUserImpl

import (
	"context"
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	"employeeSelfService/response"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type RepositoryUserImpl struct {
	db *gorm.DB
}

func NewRepositoryUserImpl(client *gorm.DB) RepositoryUserImpl {
	return RepositoryUserImpl{client}
}

func (repo RepositoryUserImpl) FindByEmail(email string) (*domain.User, *errs.AppErr) {
	var user domain.User
	if result := repo.db.Where("email = ?", email).Find(&user); result.RowsAffected == 0 {
		logger.Error("error get data user by email not found")
		return nil, errs.NewAuthenticationError("Your Login Failed! Invalid Credential")
	}

	return &user, nil
}

func (repo RepositoryUserImpl) FindById(id int64) (*domain.User, *errs.AppErr) {

	// buat variable untuk menanpung data user dari database
	var user *domain.User = &domain.User{IdUser: id}
	fmt.Println("SEBELUM MASUK FIRST", user)
	if result := repo.db.First(user); result.RowsAffected == 0 {
		fmt.Print("MASUK FIRST", result)
		logger.Error("error get data user by id not found")
		return nil, errs.NewNotFoundError("user not found")
	} else if result.Error != nil {
		fmt.Print("MASUK ELSE IF", result.Error.Error())
		logger.Error("error get data user by id not found")
		return nil, errs.NewNotFoundError("user not found")
	}
	fmt.Println("SETELAH FIRST", user)
	return user, nil
}

func (repo RepositoryUserImpl) GetDataDashboard(id string) (*response.ResponseDashboard, *errs.AppErr) {
	// buat variabel untuk menapung data user dari database
	var dashboard = response.ResponseDashboardFromDatabase{}
	ctx := context.Background()
	dateNow := time.Now().Format("2006-01-02")

	script := `select e.id_employe, pic_absensi , es.id_employe_secondary, count(a.id_absensi) as kehadiran
	, approval
	, sudah_absen
	from employee e 
	left join employeesecondarydata es ON e.id_employe = es.id_employe
	left join employee_position ep on e.id_employe = ep.id_employe 
	left join absensi a on e.id_employe = a.id_employe
	cross join(
		select count(ac.id_absen_configuration) as pic_absensi from employee e2 
		left join employee_position ep2 on e2.id_employe = ep2.id_employe 
		left join "position" p on ep2.id_position = p.id_position 
		left join absen_configuration ac on p.id_position = ac.id_absen_configuration 
		where e2.id_employe = ` + id + `
	) pbs 
	cross join(
		select count(a.id_approval) as approval from employee e 
		left join approval a on e.id_employe = a.whos_approv  
		where a.feature_type = 'absen' and e.id_employe= ` + id + `
	) ap
	cross join(
		select count(a2.id_absensi) as sudah_absen from absensi a2 where a2.tgl_absen = '` + dateNow + `' and a2.id_employe = ` + id + `
	)sa
	group by e.id_employe ,es.id_employe_secondary, ap.approval, pbs.pic_absensi, sa.sudah_absen
	having e.id_employe = ` + id + ` 
	limit 1 ;`

	rows, err := repo.db.ConnPool.QueryContext(ctx, script)

	if err != nil {
		logger.Error("error get data from " + err.Error())
		return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}

	defer rows.Close()

	for rows.Next() {
		err := repo.db.ScanRows(rows, &dashboard)
		// rows.Scan(&dashboard.IdEmploye, &dashboard.PicAbsensi, &dashboard.IdeEmployeSecondary, &dashboard.Kehadiran, &dashboard.Approval, &dashboard.SudahAbsen)

		if err != nil {
			logger.Error("error scan data dashboard " + err.Error())
			return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
		}
	}

	reponseDashboard := dashboard.ToResponseDashboard()
	return &reponseDashboard, nil
}
