package serviceClient

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	"employeeSelfService/response"
)

type ServiceClient interface {
	SaveClient(client *domain.Client) (*helper.SuccessResponseMessage, *errs.AppErr)
	GetAllClient() (*response.ResponseClient, *errs.AppErr)
	GetClientById(id int) (*response.ResponseClient, *errs.AppErr)
	DeleteClient(id int) (*helper.SuccessResponseMessage, *errs.AppErr)
	Update(client *domain.Client) (*helper.SuccessResponseMessage, *errs.AppErr)
}
