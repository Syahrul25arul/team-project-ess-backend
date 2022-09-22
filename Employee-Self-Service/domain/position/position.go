package domainPosition

type Position struct {
	IdPosition   int64  `gorm:"primaryKey"`
	PositionName string `db:"position_name"`
}

func (Position) TableName() string {
	return "position"
}
