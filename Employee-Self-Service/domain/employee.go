package domain

import (
	"database/sql"
)

type Employee struct {
	IdEmploye                 int64         `gorm:"primaryKey"`
	IdUser                    sql.NullInt64 `db:"id_user"`
	Nik                       int64         `db:"nik"`
	NamaLengkap               string        `db:"nama_lengkap"`
	TempatLahir               string        `db:"tempat_lahir"`
	TanggalLahir              string        `db:"tanggal_lahir"`
	AlamatKtp                 string        `db:"alamat_ktp"`
	PendidikanTerakhir        string        `db:"pendidikan_terakhir"`
	NamaPendidikanTerakhir    string        `db:"nama_pendidikan_terakhir"`
	JurusanPendidikanTerakhir string        `db:"jurusan_pendidikan_terkhir"`
	AlamatEmailAktif          string        `db:"alamat_email_aktif"`
	NoTlpAktif                string        `db:"no_tlp_aktif"`
	KontakDarurat             string        `db:"kontak_darurat"`
	NoTlpKontakDarurat        string        `db:"no_tlp_kontak_darurat"`
	StatusEmployee            string        `db:"status_employee"`
	PhotoEmployee             string        `db:"photo_employee"`
}

func (Employee) TableName() string {
	return "employee"
}
