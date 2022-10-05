package handlerDashboard

import (
	repositoryUser "employeeSelfService/repository/user/impl"
	serviceDashboard "employeeSelfService/service/dashboard"
	serviceDashboardImpl "employeeSelfService/service/dashboard/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handlerDashboard struct {
	service serviceDashboard.ServiceDashboard
}

func NewHandlerDashboard(db *gorm.DB) handlerDashboard {
	return handlerDashboard{service: serviceDashboardImpl.NewServiceDashboard(repositoryUser.NewRepositoryUserImpl(db))}
}

func (h handlerDashboard) GetDashboardHandler(ctx *gin.Context) {
	// tangkap id parameter
	userId := ctx.Param("user_id")

	// get data from service
	if resp, err := h.service.GetDashboard(userId); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(int(resp.Code), resp)
	}
}
