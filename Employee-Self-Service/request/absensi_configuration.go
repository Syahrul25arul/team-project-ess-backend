package request

type AbsensiConfiguration struct {
	DurasiJamKerja             int8    `json:"durasi_jam_kerja"`
	IntervalKeterlambatan      int8    `json:"interval_keterlambatan"`
	BobotKeterlambatan         float32 `json:"bobot_keterlambatan"`
	MaksimalBobotKeterlambatan float32 `json:"maksimal_bobot_keterlambatan"`
	IdPosition                 int64   `json:"id_position"`
	MinimalMasukJamKerja       string  `json:"minimal_masuk_jam_kerja"`
	MaksimalMasukJamKerja      string  `json:"maksimal_masuk_jam_kerja"`
}
