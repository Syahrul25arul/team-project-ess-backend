package repositoryAbsenConfiguration

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
)

type RepositoiryAbsenConfiguration interface {
	GetData() (*domain.AbsenConfiguration, *errs.AppErr)
	Save(absenConfiguration *domain.AbsenConfiguration) *errs.AppErr
}
