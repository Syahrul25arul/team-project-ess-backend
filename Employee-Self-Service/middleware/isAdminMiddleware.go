package middleware

import (
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	repoUser "employeeSelfService/repository/user/impl"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IsAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		repoUser := repoUser.NewRepositoryUserImpl(db)

		// get User id
		userId, err := strconv.Atoi(ctx.Param("user_id"))

		if err != nil {
			errResponse := errs.NewBadRequestError("id not relevant from middleware")
			logger.Error("error convert iduser " + err.Error())
			ctx.AbortWithStatusJSON(errResponse.Code, errResponse)
		}

		// get data user by param user id
		if user, errF := repoUser.FindById(int64(userId)); errF != nil {
			ctx.AbortWithStatusJSON(errF.Code, errF)
		} else {
			// cek user is not admin
			if user.UserRole != "admin" {
				errResponse := errs.NewForbiddenError("Forbidden, you dont have credential")
				logger.Error("error error is admin " + errResponse.Message)
				ctx.AbortWithStatusJSON(errResponse.Code, errResponse)
			}
		}
	}
}
