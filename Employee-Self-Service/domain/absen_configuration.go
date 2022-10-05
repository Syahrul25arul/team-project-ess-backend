package domain

type AbsenConfiguration struct {
	IdAbsenConfiguration       int64   `gorm:"primaryKey"`
	DurasiJamKerja             int8    `db:"durasi_jam_kerja"`
	IntervalKeterlambatan      int8    `db:"interval_keterlambatan"`
	BobotKeterlambatan         float32 `db:"bobot_keterlambatan"`
	MaksimalBobotKeterlambatan float32 `db:"maksimal_bobot_keterlambatan"`
	IdPosition                 int64   `db:"id_position"`
	MinimalMasukJamKerja       string  `db:"minimal_masuk_jam_kerja"`
	MaksimalMasukJamKerja      string  `db:"maksimal_masuk_jam_kerja"`
}

func (AbsenConfiguration) TableName() string {
	return "absen_configuration"
}
