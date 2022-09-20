package registerHandler

import (
	registerRequest "employeeSelfService/request/register"
	responseRegister "employeeSelfService/response/register"
	serviceRegister "employeeSelfService/service/register"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerRegister struct {
	service serviceRegister.ServiceRegister
}

func (h HandlerRegister) RegisterHandler(ctx *gin.Context) {
	// tangkap request body dari client
	var register *registerRequest.RegisterRequest
	ctx.ShouldBindJSON(&register)

	if err := h.service.Register(register); err != nil {
		// jika terjdi error tampilkan error
		ctx.JSON(err.Code, err.Message)
	} else {
		// response success
		response := responseRegister.NewReponseRegisterSuccess()
		// jika tidak error, berikan response ke client
		ctx.JSON(http.StatusCreated, response)
	}
}
