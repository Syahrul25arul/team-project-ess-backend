package domain

type Client struct {
	IdClient     int32   `gorm:"PrimaryKey" json:"id_client"`
	NamaClient   string  `db:"nama_client" json:"nama_client"`
	Lattitude    float32 `db:"lattitude"`
	Longitude    float32 `db:"longitude"`
	AlamatClient string  `db:"alamat_client" json:"alamat_client"`
}

type ClientWithProject struct {
	IdClient     int32     `gorm:"PrimaryKey" json:"id_client"`
	NamaClient   string    `db:"nama_client" json:"nama_client"`
	Lattitude    float32   `db:"lattitude"`
	Longitude    float32   `db:"longitude"`
	AlamatClient string    `db:"alamat_client" json:"alamat_client"`
	Projects     []Project `gorm:"foreignKey:IdClient"`
}

func (Client) TableName() string {
	return "client"
}

func (ClientWithProject) TableName() string {
	return "client"
}
