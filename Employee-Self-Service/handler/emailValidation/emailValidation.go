package emailValidationHandler

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	repositoryEmailValidation "employeeSelfService/repository/emailValidation/impl"
	repositoryUser "employeeSelfService/repository/user/impl"
	serviceEmailValidaton "employeeSelfService/service/emailValidation"
	serviceEmailValidationImpl "employeeSelfService/service/emailValidation/impl"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerEmailValidation struct {
	service serviceEmailValidaton.ServiceEmailValidation
}

func NewHandlerEmailValidation(db *gorm.DB) HandlerEmailValidation {
	repo := repositoryEmailValidation.NewRepositoryEmailValidationImpl(db)
	repoUser := repositoryUser.NewRepositoryUserImpl(db)
	emailValidationService := serviceEmailValidationImpl.NewServiceEmailValidationImpl(repo, repoUser)
	return HandlerEmailValidation{&emailValidationService}
}

func (handler HandlerEmailValidation) SaveEmailValidation(ctx *gin.Context) {
	// tangkap request dari client
	var emailValidation *domain.EmailValidation
	ctx.ShouldBindJSON(&emailValidation)

	// tangkap parameter userId
	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		errorResponse := errs.NewUnexpectedError("terjadi kesalahan, pastikan url anda sudah benar")
		ctx.JSON(errorResponse.Code, errorResponse)
		return
	}

	// kirim data emailValidation ke service, dan tankap response dari service
	resp := handler.service.Save(emailValidation, int64(userId))

	ctx.JSON(resp.Code, resp)

}
