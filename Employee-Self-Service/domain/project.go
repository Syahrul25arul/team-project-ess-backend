package domain

type Project struct {
	IdProject   int32  `gorm:"PrimaryKey"`
	IdClient    int32  `db:"id_client"`
	ProjectName string `db:"project_name"`
}
type ProjectWithClient struct {
	IdProject   int32  `gorm:"PrimaryKey"`
	ProjectName string `db:"project_name"`
	IdClient    int32  `db:"id_client"`
	Client      Client `gorm:"foreignKey:IdClient"`
}

func (Project) TableName() string {
	return "project"
}

func (ProjectWithClient) TableName() string {
	return "project"
}
