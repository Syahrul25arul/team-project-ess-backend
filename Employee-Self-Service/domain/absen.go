package domain

import (
	"database/sql"
)

type Absen struct {
	IdAbsensi      int64         `gorm:"primaryKey"`
	IdEmploye      int64         `db:"id_employe"`
	Id_Cuti        sql.NullInt64 `db:"id_cuti"`
	TglAbsen       string        `db:"tgl_absen"`
	Masuk          string        `db:"masuk"`
	Pulang         string        `db:"pulang"`
	LokasiAbsen    string        `db:"lokasi_absen"`
	TerlambatKerja int           `db:"terlambat_kerja"`
	Bobot          int           `db:"bobot"`
	OverTime       int           `db:"over_time"`
	StatusOvertime string        `db:"status_overtime"`
	Keterangan     string        `db:"keterangan"`
}

func (Absen) TableName() string {
	return "absensi"
}
