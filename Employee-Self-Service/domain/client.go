package domain

type Client struct {
	IdClient     int32   `gorm:"PrimaryKey" json:"id_client"`
	NamaClient   string  `db:"nama_client" json:"nama_client"`
	Lattitude    float32 `db:"lattitude"`
	Longitude    float32 `db:"longitude"`
	AlamatClient string  `db:"alamat_client" json:"alamat_client"`
}

func (Client) TableName() string {
	return "client"
}
