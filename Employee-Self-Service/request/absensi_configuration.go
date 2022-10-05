package request

import "employeeSelfService/domain"

type AbsensiConfiguration struct {
	DurasiJamKerja             int8    `json:"durasi_jam_kerja"`
	IntervalKeterlambatan      int8    `json:"interval_keterlambatan"`
	BobotKeterlambatan         float32 `json:"bobot_keterlambatan"`
	MaksimalBobotKeterlambatan float32 `json:"maksimal_bobot_keterlambatan"`
	IdPosition                 int64   `json:"id_position"`
	MinimalMasukJamKerja       string  `json:"minimal_masuk_jam_kerja"`
	MaksimalMasukJamKerja      string  `json:"maksimal_masuk_jam_kerja"`
}

func (ab AbsensiConfiguration) ToDomainAbsen() *domain.AbsenConfiguration {
	return &domain.AbsenConfiguration{
		DurasiJamKerja:             ab.DurasiJamKerja,
		IntervalKeterlambatan:      ab.IntervalKeterlambatan,
		BobotKeterlambatan:         ab.BobotKeterlambatan,
		MaksimalBobotKeterlambatan: ab.MaksimalBobotKeterlambatan,
		IdPosition:                 ab.IdPosition,
		MinimalMasukJamKerja:       ab.MinimalMasukJamKerja,
		MaksimalMasukJamKerja:      ab.MaksimalMasukJamKerja,
	}
}
