package serviceDashboard

import (
	"employeeSelfService/errs"
	"employeeSelfService/response"
)

type ServiceDashboard interface {
	GetDashboard(id string) (*response.ResponseDashboard, *errs.AppErr)
}
