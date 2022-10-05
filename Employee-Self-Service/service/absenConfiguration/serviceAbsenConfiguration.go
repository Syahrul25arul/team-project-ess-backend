package serviceAbsenConfiguration

import (
	"employeeSelfService/request"
	"employeeSelfService/response"
)

type ServiceAbsenConfiguration interface {
	Save(absenConfiguration *request.AbsensiConfiguration, idUser int64) response.ResponseAbsenConfiguration
}
