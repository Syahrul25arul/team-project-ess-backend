package response

import "database/sql"

type ResponseDashboardFromDatabase struct {
	IdEmploye           int64         `db:"id_employee"`
	PicAbsensi          int           `db:"pic_absensi"`
	IdeEmployeSecondary sql.NullInt64 `db:"id_employee_secondary"`
	Kehadiran           int32         `db:"kehadiran"`
	Approval            int32         `db:"approval"`
	SudahAbsen          int           `db:"sudah_absen"`
}

type ResponseDashboard struct {
	Status              string `json:"status"`
	Code                uint   `json:"code"`
	PicAbsensi          bool   `json:"pic_absensi"`
	IdeEmployeSecondary bool   `json:"id_employee_secondary"`
	Kehadiran           int32  `json:"kehadiran"`
	Approval            int32  `json:"approval"`
	SudahAbsen          bool   `json:"sudha_absen"`
}

func (r ResponseDashboardFromDatabase) ToResponseDashboard() ResponseDashboard {
	responseDashboard := ResponseDashboard{
		Status:    "ok",
		Code:      200,
		Kehadiran: r.Kehadiran,
		Approval:  r.Approval,
	}
	if r.PicAbsensi == 0 {
		responseDashboard.PicAbsensi = false
	} else {
		responseDashboard.PicAbsensi = true
	}

	if !r.IdeEmployeSecondary.Valid {
		responseDashboard.IdeEmployeSecondary = false
	} else {
		responseDashboard.IdeEmployeSecondary = true
	}

	if r.SudahAbsen == 0 {
		responseDashboard.SudahAbsen = false
	} else {
		responseDashboard.SudahAbsen = true
	}

	return responseDashboard
}
