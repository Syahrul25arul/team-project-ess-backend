package handlerAbsenConfiguration

import (
	"employeeSelfService/errs"
	repoAbsenConfiguration "employeeSelfService/repository/absenConfiguration/impl"
	repoUser "employeeSelfService/repository/user/impl"
	"employeeSelfService/request"
	serviceAbsenConfiguration "employeeSelfService/service/absenConfiguration"
	serviceAbsenConfigurationImpl "employeeSelfService/service/absenConfiguration/impl"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerAbsenConfiguration struct {
	service serviceAbsenConfiguration.ServiceAbsenConfiguration
}

func NewHandlerAbsenConfiguration(db *gorm.DB) HandlerAbsenConfiguration {
	repoAbsenConfiguration := repoAbsenConfiguration.NewRepositoryAbsenConfigurationImpl(db)
	repositoryUser := repoUser.NewRepositoryUserImpl(db)
	service := serviceAbsenConfigurationImpl.NewServiceEmailValidationImpl(repoAbsenConfiguration, repositoryUser)
	return HandlerAbsenConfiguration{service: service}
}

func (handler HandlerAbsenConfiguration) SaveAbsenConfiguration(ctx *gin.Context) {
	// tangkap request dari client
	var absenConfiguration *request.AbsensiConfiguration
	ctx.ShouldBindJSON(&absenConfiguration)

	// tangkap parameter userId
	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		errorResponse := errs.NewUnexpectedError("something wrong")
		ctx.JSON(errorResponse.Code, errorResponse)
		return
	}

	// kirim data emailValidation ke service, dan tankap response dari service
	resp := handler.service.Save(absenConfiguration, int64(userId))

	ctx.JSON(resp.Code, resp)

}
