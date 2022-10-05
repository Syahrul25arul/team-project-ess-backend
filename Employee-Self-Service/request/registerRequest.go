package request

import (
	"employeeSelfService/domain"
)

type RegisterRequest struct {
	NamaLengkap               string `json:"nama_lengkap"`
	Email                     string `json:"email"`
	Password                  string `json:"password"`
	TempatLahir               string `json:"tempat_lahir"`
	TanggalLahir              string `json:"tanggal_lahir"`
	Nik                       int64  `json:"nik"`
	AlamatKtp                 string `json:"alamat_ktp"`
	PendidikanTerakhir        string `json:"pendidikan_terakhir"`
	NamaPendidikanTerakhir    string `json:"nama_pendidikan_terakhir"`
	JurusanPendidikanTerakhir string `json:"jurusan_pendidikan_terakhir"`
	AlamatEmailAktif          string `json:"alamat_email_aktif"`
	NoTlpAktif                string `json:"no_tlp_aktif"`
	KontakDarurat             string `json:"kontak_darurat"`
	NoTlpKontakDarurat        string `json:"no_tlp_kontak_darurat"`
	StatusEmployee            string `json:"status_employee"`
	PhotoEmployee             string `json:"photo_employee"`
}

func (u RegisterRequest) ToUser() *domain.User {
	return &domain.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u RegisterRequest) ToEmployee() *domain.Employee {
	return &domain.Employee{
		Nik:                       u.Nik,
		NamaLengkap:               u.NamaLengkap,
		TempatLahir:               u.TempatLahir,
		TanggalLahir:              u.TanggalLahir,
		AlamatKtp:                 u.AlamatKtp,
		PendidikanTerakhir:        u.PendidikanTerakhir,
		NamaPendidikanTerakhir:    u.NamaPendidikanTerakhir,
		JurusanPendidikanTerakhir: u.JurusanPendidikanTerakhir,
		AlamatEmailAktif:          u.AlamatEmailAktif,
		NoTlpAktif:                u.NoTlpAktif,
		KontakDarurat:             u.KontakDarurat,
		NoTlpKontakDarurat:        u.NoTlpKontakDarurat,
		StatusEmployee:            u.StatusEmployee,
		PhotoEmployee:             u.PhotoEmployee,
	}
}
