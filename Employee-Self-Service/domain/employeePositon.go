package domain

type EmployeePosition struct {
	IdEmployeePosition int64 `gorm:"primaryKey"`
	IdPosition         int64 `db:"id_position"`
	IdEmploye          int64 `db:"id_employe"`
}

func (EmployeePosition) TableName() string {
	return "employee_position"
}
