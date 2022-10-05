package domain

type Approval struct {
	IdApproval     int64   `gorm:"primaryKey"`
	FeatureType    string  `db:"feature_type"`
	WhosApprov     int64   `db:"whos_approv"`
	StatusApprov   string  `db:"status_approv"`
	FeatureId      int64   `db:"feature_id"`
	TujuanApproval string  `db:"tujuan_approval"`
	CreatedAt      *string `db:"created_at"`
	UpdatedAt      *string `db:"updated_at"`
	Keterangan     string  `db:"keterangan"`
}

func (Approval) TableName() string {
	return "approval"
}
