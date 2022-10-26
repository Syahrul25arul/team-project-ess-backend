package domain

type Project struct {
	IdProject   int32  `gorm:"PrimaryKey" json:"id_project"`
	IdClient    int32  `db:"id_client" json:"id_client"`
	ProjectName string `db:"project_name" json:"project_name"`
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
