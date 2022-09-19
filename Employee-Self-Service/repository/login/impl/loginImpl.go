package loginImpl

import (
	"gorm.io/gorm"
)

type RepositoryLoginImpl struct {
	DB *gorm.DB
}

func NewRepositoryLoginImpl(client *gorm.DB) RepositoryLoginImpl {
	return RepositoryLoginImpl{client}
}
