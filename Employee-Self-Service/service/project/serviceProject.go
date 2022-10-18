package serviceProject

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	"employeeSelfService/response"
)

type ServiceProject interface {
	SaveProject(project *domain.Project) (*helper.SuccessResponseMessage, *errs.AppErr)
	GetAllProject() (*response.ResponseProject, *errs.AppErr)
	GetById(id int32) (*response.ResponseProject, *errs.AppErr)
	Update(project *domain.Project) (*helper.SuccessResponseMessage, *errs.AppErr)
	Delete(id int32) (*helper.SuccessResponseMessage, *errs.AppErr)
}
