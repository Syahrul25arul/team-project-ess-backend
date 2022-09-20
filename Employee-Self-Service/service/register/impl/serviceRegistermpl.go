package serviceRegisterImpl

import (
	"employeeSelfService/config"
	"employeeSelfService/errs"
	"employeeSelfService/helper"
	repositoryRegister "employeeSelfService/repository/register"
	registerRequest "employeeSelfService/request/register"
)

type ServiceRegisterImpl struct {
	repo repositoryRegister.RepositoryInterface
}

func NewCustomerService(repo repositoryRegister.RepositoryInterface) ServiceRegisterImpl {
	return ServiceRegisterImpl{repo}
}

func (s *ServiceRegisterImpl) Register(registerRequest *registerRequest.RegisterRequest) *errs.AppErr {
	// convert registerRequest to user
	user := registerRequest.ToUser()
	user.StatusVerified = "true"
	user.UserRole = "employee"

	// hash password
	user.Password = helper.BcryptPassword(config.SECRET_KEY + user.Password)

	// convert registerRequest to employee
	employee := registerRequest.ToEmployee()

	// send to repo register for save data
	return s.repo.Register(user, employee)

}
