package response

import "database/sql"

type ResponseDashboardFromDatabase struct {
	Status              string `db:"status"`
	Code                string `db:"code"`
	PicAbsensi          int64  `db:"pic_absensi"`
	IdeEmployeSecondary sql.NullInt64
	KehadiranCreated    string `db:"kehadiran_created"`
	KehadiranDenied     string `db:"kehadiran_denied"`
	KehadiranReview     string `db:"kehadiran_review"`
	KehadiranApprove    string `db:"kehadiran_approve"`
	Approval            int32
}
