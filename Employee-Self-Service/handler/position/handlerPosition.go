package handlerPosition

import (
	"employeeSelfService/domain"
	repoPosition "employeeSelfService/repository/position/impl"
	"employeeSelfService/request"
	"employeeSelfService/response"
	servicePosition "employeeSelfService/service/position"
	servicePositionImpl "employeeSelfService/service/position/impl"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerPosition struct {
	service servicePosition.ServicePosition
}

func NewHandlerPosition(db *gorm.DB) HandlerPosition {
	repoPosition := repoPosition.NewRepositoryPositionImpl(db)
	service := servicePositionImpl.NewRepositoryPositionImpl(repoPosition)
	return HandlerPosition{service: service}
}

func (handler HandlerPosition) SavePosition(ctx *gin.Context) {
	var positionRequest *request.Position
	ctx.ShouldBindJSON(&positionRequest)

	if _, err := handler.service.Save(positionRequest); err != nil {
		ctx.JSON(err.Code, "Failed save data position!!")
	} else {
		// response success
		response := response.NewReponsePositionSuccess()
		// jika tidak error, berikan response ke client
		ctx.JSON(http.StatusCreated, response)
	}
}

func (handler HandlerPosition) GetPositionById(ctx *gin.Context) {
	idposition := ctx.Param("id_position")
	IdPosition, _ := strconv.Atoi(idposition)

	positions, err := handler.service.FindById(int64(IdPosition))
	if err != nil {
		errsMessage := fmt.Sprintf("Failed get data position by id %v error!", idposition)
		res := response.NewResponsePositionFailed(500, errsMessage)
		ctx.JSON(http.StatusBadRequest, res)
		return

	} else {
		response := response.NewReponsePositionSuccess()
		ctx.JSON(http.StatusOK, response)
	}

	ctx.JSON(http.StatusOK, positions)
}

func (handler HandlerPosition) DeletePosition(ctx *gin.Context) {
	idposition := ctx.Param("id_position")
	IdPosition, _ := strconv.Atoi(idposition)

	_, err := handler.service.Delete(int64(IdPosition))
	if err != nil {
		errsMessage := fmt.Sprintf("Failed delete data position %v error!", idposition)
		res := response.NewResponsePositionFailed(500, errsMessage)
		ctx.JSON(http.StatusBadRequest, res)
		return
	} else {
		response := response.NewReponsePositionDeleteSuccess()
		ctx.JSON(http.StatusOK, response)
	}
}

func (handler HandlerPosition) UpdatePosition(ctx *gin.Context) {
	var position domain.Position
	ctx.ShouldBindJSON(&position)

	idposition := ctx.Param("id_position")
	IdPositions, _ := strconv.Atoi(idposition)

	position.IdPosition = int64(IdPositions)

	if err := handler.service.Update(position); err != nil {
		errsMessage := fmt.Sprintf("Failed update data position %v error!", idposition)
		res := response.NewResponsePositionFailed(500, errsMessage)
		ctx.JSON(http.StatusBadRequest, res)
		return
	} else {
		response := response.NewReponsePositionUpdateSuccess()
		ctx.JSON(http.StatusOK, response)
	}

}
