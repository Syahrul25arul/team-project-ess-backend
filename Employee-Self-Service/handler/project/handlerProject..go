package handlerProject

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	repositoryProjectImpl "employeeSelfService/repository/project/impl"
	serviceProject "employeeSelfService/service/project"
	serviceProjectImpl "employeeSelfService/service/project/impl"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handlerProject struct {
	service serviceProject.ServiceProject
}

func NewHandlerProject(db *gorm.DB) handlerProject {
	repositoryProject := repositoryProjectImpl.NewRepositoryProject(db)
	serviceProject := serviceProjectImpl.NewServiceProject(repositoryProject)
	return handlerProject{service: serviceProject}
}

func (h handlerProject) SaveProject(ctx *gin.Context) {
	// create variabel client and catch data client from request
	var project *domain.Project
	err := ctx.ShouldBindJSON(&project)

	fmt.Println("======== REQUEST BODY ==========")
	fmt.Println(project)

	if err != nil {
		logger.Error("error scan data project")
		errResponse := errs.NewValidationError("invalid data project")
		ctx.JSON(errResponse.Code, errResponse)
		return
	}

	if response, err := h.service.SaveProject(project); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}

func (h handlerProject) UpdateProject(ctx *gin.Context) {
	// create variabel project and catch data project from request
	var project *domain.Project
	err := ctx.ShouldBindJSON(&project)

	if err != nil {
		logger.Error("error scan data project")
		errResponse := errs.NewValidationError("invalid data project")
		ctx.JSON(errResponse.Code, errResponse)
		return
	}

	if response, err := h.service.Update(project); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}

func (h handlerProject) GetAllProject(ctx *gin.Context) {
	// get all project from database
	if response, err := h.service.GetAllProject(); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}

func (h handlerProject) GetProjectById(ctx *gin.Context) {
	// get id project for parameter
	id_client, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		errorResponse := errs.NewBadRequestError("id project not valid")
		ctx.JSON(errorResponse.Code, errorResponse)
		return
	}

	// get project from database
	if response, err := h.service.GetById(int32(id_client)); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}

func (h handlerProject) DeleteProject(ctx *gin.Context) {
	// get id project for parameter
	id_client, err := strconv.Atoi(ctx.Param("project_id"))
	if err != nil {
		errorResponse := errs.NewBadRequestError("id project not valid")
		ctx.JSON(errorResponse.Code, errorResponse)
		return
	}

	// delete project from database
	if response, err := h.service.Delete(int32(id_client)); err != nil {
		ctx.JSON(err.Code, err)
	} else {
		ctx.JSON(response.Code, response)
	}
}
