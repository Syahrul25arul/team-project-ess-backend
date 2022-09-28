package database

import (
	"database/sql"
	"employeeSelfService/config"
	"employeeSelfService/domain"
	"employeeSelfService/helper"
	"time"

	"gorm.io/gorm"
)

func SetupDataDummy(db *gorm.DB) {
	tx := db.Begin()

	// create user
	userTest := &domain.User{
		Email:          "test@gmail.com",
		Password:       helper.BcryptPassword(config.SECRET_KEY + "password"),
		UserRole:       "employee",
		StatusVerified: "true",
	}
	userHr := &domain.User{
		Email:          "kresna@gmail.com",
		Password:       helper.BcryptPassword(config.SECRET_KEY + "password"),
		UserRole:       "employee",
		StatusVerified: "true",
	}
	userAdmin := &domain.User{
		Email:          "admin@gmail.com",
		Password:       helper.BcryptPassword(config.SECRET_KEY + "password"),
		UserRole:       "admin",
		StatusVerified: "true",
	}
	tx.Create(userTest)
	tx.Create(userAdmin)
	tx.Create(userHr)

	// create employee
	employeeTest := &domain.Employee{
		NamaLengkap:               "Teddy",
		TempatLahir:               "Jakarta",
		TanggalLahir:              "13-09-1992",
		Nik:                       2389235897352,
		AlamatKtp:                 "Cilandak timur, jeruk purut",
		PendidikanTerakhir:        "Sarjana",
		NamaPendidikanTerakhir:    "USTJ",
		JurusanPendidikanTerakhir: "Teknik Informatika",
		AlamatEmailAktif:          "teddythebear@gmail.com",
		NoTlpAktif:                "967826342389",
		KontakDarurat:             "motherBear",
		NoTlpKontakDarurat:        "2938789",
		StatusEmployee:            "aktif",
		PhotoEmployee:             "teddyBear.jpg",
	}
	employeeTest.IdUser = sql.NullInt64{Int64: int64(userTest.IdUser), Valid: true}
	employeeHR := &domain.Employee{
		NamaLengkap:               "Kresna",
		TempatLahir:               "Jakarta",
		TanggalLahir:              "13-09-1992",
		Nik:                       2389235897352,
		AlamatKtp:                 "Cilandak timur, jeruk purut",
		PendidikanTerakhir:        "Sarjana",
		NamaPendidikanTerakhir:    "USTJ",
		JurusanPendidikanTerakhir: "Teknik Informatika",
		AlamatEmailAktif:          "teddythebear@gmail.com",
		NoTlpAktif:                "967826342389",
		KontakDarurat:             "motherBear",
		NoTlpKontakDarurat:        "2938789",
		StatusEmployee:            "aktif",
		PhotoEmployee:             "teddyBear.jpg",
	}
	employeeHR.IdUser = sql.NullInt64{Int64: int64(userHr.IdUser), Valid: true}
	tx.Create(employeeTest)
	tx.Create(employeeHR)

	// position
	positionHR := &domain.Position{
		PositionName: "HR",
	}
	positionTL := &domain.Position{
		PositionName: "Team Lead",
	}
	tx.Create(positionHR)
	tx.Create(positionTL)

	// employee position
	employeePosition := &domain.EmployeePosition{
		IdPosition: positionTL.IdPosition,
		IdEmploye:  employeeTest.IdEmploye,
	}

	hrPosition := &domain.EmployeePosition{
		IdPosition: positionHR.IdPosition,
		IdEmploye:  employeeHR.IdEmploye,
	}

	tx.Create(employeePosition)
	tx.Create(hrPosition)

	// create absen configuration
	absenConfiguration := &domain.AbsenConfiguration{
		DurasiJamKerja:             8,
		IntervalKeterlambatan:      15,
		BobotKeterlambatan:         0.25,
		MaksimalBobotKeterlambatan: 1,
		PicAbsensi:                 positionHR.IdPosition,
		MinimalMasukJamKerja:       "08:00",
		MaksimalMasukJamKerja:      "10:00",
	}
	tx.Create(absenConfiguration)

	// absen
	absen := &domain.Absen{
		IdEmploye:      employeeTest.IdEmploye,
		TglAbsen:       time.Now().Format("2006-01-02"),
		Masuk:          time.Now().Format("15:04"),
		Pulang:         "18:00",
		LokasiAbsen:    "main office",
		TerlambatKerja: 0,
		Bobot:          0,
		OverTime:       0,
		StatusOvertime: "false",
		Keterangan:     "tidak ada",
	}

	resultAbsen, _ := time.Parse("2006-01-02 15:04:05", "2022-09-24 08:00:00")
	absen2 := &domain.Absen{
		IdEmploye:      employeeTest.IdEmploye,
		TglAbsen:       resultAbsen.Format("2006-01-02"),
		Masuk:          resultAbsen.Format("15:04"),
		Pulang:         "18:00",
		LokasiAbsen:    "main office",
		TerlambatKerja: 0,
		Bobot:          0,
		OverTime:       0,
		StatusOvertime: "false",
		Keterangan:     "tidak ada",
	}
	tx.Create(absen)
	tx.Create(absen2)

	// approval
	pic_approval := employeeHR.IdEmploye
	waktuAbsenMasuk := time.Now().Format("2006-01-02 15:04:05")
	approval := &domain.Approval{
		FeatureType:    "absen",
		WhosApprov:     pic_approval,
		StatusApprov:   "diterima",
		FeatureId:      absen.IdAbsensi,
		TujuanApproval: "absen pagi",
		CreatedAt:      &waktuAbsenMasuk,
		Keterangan:     "approval untuk absen pagi",
	}

	result, _ := time.Parse("2006-01-02 15:04:05", "2022-09-23 17:00:00")
	waktuAbsenPulang := result.Format("2006-01-02 15:04:05")
	approvalPulang := &domain.Approval{
		FeatureType:    "absen",
		WhosApprov:     pic_approval,
		StatusApprov:   "diterima",
		FeatureId:      absen.IdAbsensi,
		TujuanApproval: "absen pagi",
		CreatedAt:      &waktuAbsenPulang,
		Keterangan:     "approval untuk absen pagi",
	}
	tx.Create(approval)
	tx.Create(approvalPulang)

	pic_approval2 := employeeHR.IdEmploye
	waktuAbsenMasuk2 := time.Now().Format("2006-01-02 15:04:05")
	approval2 := &domain.Approval{
		FeatureType:    "absen",
		WhosApprov:     pic_approval2,
		StatusApprov:   "ditolak",
		FeatureId:      absen2.IdAbsensi,
		TujuanApproval: "absen pagi",
		CreatedAt:      &waktuAbsenMasuk2,
		Keterangan:     "approval untuk absen pagi",
	}

	result2, _ := time.Parse("2006-01-02 15:04:05", "2022-09-23 17:00:00")
	waktuAbsenPulang2 := result2.Format("2006-01-02 15:04:05")
	approvalPulang2 := &domain.Approval{
		FeatureType:    "absen",
		WhosApprov:     pic_approval,
		StatusApprov:   "ditinjau",
		FeatureId:      absen2.IdAbsensi,
		TujuanApproval: "absen pagi",
		CreatedAt:      &waktuAbsenPulang2,
		Keterangan:     "approval untuk absen pagi",
	}
	tx.Create(approval2)
	tx.Create(approvalPulang2)

	tx.Commit()
}
