package registerHandler

import (
	repositoryRegisterImpl "employeeSelfService/repository/register/impl"
	"employeeSelfService/request"
	"employeeSelfService/response"
	serviceRegister "employeeSelfService/service/register"
	serviceRegisterImpl "employeeSelfService/service/register/impl"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerRegister struct {
	service serviceRegister.ServiceRegister
}

func NewHandlerRegister(db *gorm.DB) HandlerRegister {
	registerRepository := repositoryRegisterImpl.NewRepositoryRegisterImpl(db)
	registerService := serviceRegisterImpl.NewCustomerService(registerRepository)
	return HandlerRegister{&registerService}
}

func (h HandlerRegister) RegisterHandler(ctx *gin.Context) {
	// tangkap request body dari client
	var register *request.RegisterRequest
	ctx.ShouldBindJSON(&register)

	if err := h.service.Register(register); err != nil {
		// jika terjdi error tampilkan error
		ctx.JSON(err.Code, err.Message)
	} else {
		// response success
		response := response.NewReponseRegisterSuccess()
		// jika tidak error, berikan response ke client
		ctx.JSON(http.StatusCreated, response)
	}
}
