package handlerClient

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	repositoryClientImpl "employeeSelfService/repository/client/impl"
	serviceClient "employeeSelfService/service/client"
	serviceClientImpl "employeeSelfService/service/client/impl"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handlerClient struct {
	service serviceClient.ServiceClient
}

func NewHandlerClient(db *gorm.DB) handlerClient {
	repositoryClient := repositoryClientImpl.NewRepostioryClient(db)
	service := serviceClientImpl.NewServiceClient(repositoryClient)
	return handlerClient{service: service}
}

func (h handlerClient) SaveClient(ctx *gin.Context) {
	// create variabel client and catch data client from request
	var client *domain.Client
	err := ctx.ShouldBindJSON(&client)

	if err != nil {
		logger.Error("error scan data client")
		errResponse := errs.NewValidationError("invalid data client")
		ctx.JSON(errResponse.Code, errResponse)
		return
	}

	if response, err := h.service.SaveClient(client); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}

func (h handlerClient) GetAllClient(ctx *gin.Context) {
	// get all client from database
	if response, err := h.service.GetAllClient(); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}

func (h handlerClient) GetClientById(ctx *gin.Context) {
	// get id client for parameter
	id_client, err := strconv.Atoi(ctx.Param("client_id"))
	if err != nil {
		errorResponse := errs.NewBadRequestError("id client not valid")
		ctx.JSON(errorResponse.Code, errorResponse)
		return
	}

	// get client from database
	if response, err := h.service.GetClientById(id_client); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}

func (h handlerClient) DeleteClient(ctx *gin.Context) {
	// get id client for parameter
	id_client, err := strconv.Atoi(ctx.Param("client_id"))
	if err != nil {
		errorResponse := errs.NewBadRequestError("id client not valid")
		ctx.JSON(errorResponse.Code, errorResponse)
		return
	}

	// get client from database
	if response, err := h.service.DeleteClient(id_client); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}

func (h handlerClient) UpdateClient(ctx *gin.Context) {
	// create variabel client and catch data client from request
	var client *domain.Client
	err := ctx.ShouldBindJSON(&client)

	if err != nil {
		logger.Error("error scan data client")
		errResponse := errs.NewValidationError("invalid data client")
		ctx.JSON(errResponse.Code, errResponse)
		return
	}

	if response, err := h.service.Update(client); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}
