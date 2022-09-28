package loginHandler

import (
	repositoryAuth "employeeSelfService/repository/auth/impl"
	"employeeSelfService/request"
	"employeeSelfService/response"
	serviceLogin "employeeSelfService/service/login"
	serviceLoginImpl "employeeSelfService/service/login/impl"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerLogin struct {
	service serviceLogin.ServiceLogin
}

func NewHandlerLogin(db *gorm.DB) HandlerLogin {
	repo := repositoryAuth.NewRepositoryAuthImpl(db)
	loginService := serviceLoginImpl.NewLoginService(repo)
	return HandlerLogin{&loginService}
}

func (h HandlerLogin) LoginHandler(ctx *gin.Context) {
	// tangkap request dari login
	var login *request.LoginRequest
	ctx.ShouldBindJSON(&login)

	fmt.Println("==== login request====", login)

	if resp, err := h.service.Login(login); err != nil {
		// jika terjdi error tampilkan error
		ctx.JSON(err.Code, response.NewLoginFailed(err.Code, err.Message))
	} else {
		// jika tidak error, berikan response ke client
		ctx.JSON(http.StatusOK, resp)
	}
}
