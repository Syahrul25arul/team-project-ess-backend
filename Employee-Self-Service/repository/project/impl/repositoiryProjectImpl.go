package repositoryProjectImpl

import (
	"employeeSelfService/domain"
	"employeeSelfService/errs"
	"employeeSelfService/logger"
	"errors"

	"gorm.io/gorm"
)

type repositoryProjectImpl struct {
	db *gorm.DB
}

func NewRepositoryProject(db *gorm.DB) repositoryProjectImpl {
	return repositoryProjectImpl{db: db}
}

func (r repositoryProjectImpl) SaveProject(project *domain.Project) *errs.AppErr {
	if tx := r.db.Save(project); tx.Error != nil {
		logger.Error("error save project : " + tx.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}

func (r repositoryProjectImpl) GetAllProject() ([]domain.ProjectWithClient, *errs.AppErr) {
	// create variable for get all data project
	var projects []domain.ProjectWithClient

	// get rows project from database
	if sql, err := r.db.Table("project").Joins("inner join client on project.id_client = client.id_client ").Rows(); err != nil {
		// create logger error and return error
		logger.Error("error get all data project " + err.Error())
		return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	} else {
		// close rows at the end execution program
		defer sql.Close()

		// loop rows project from database
		for sql.Next() {
			// create variable domain project and client
			var project *domain.ProjectWithClient
			var client *domain.Client

			// scan data project and save in variable
			err := r.db.ScanRows(sql, &project)

			// if any error, create logger error and return error
			if err != nil {
				logger.Error("error scan data project with client " + err.Error())
				return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
			}

			// scan data client from rows and save in variable
			err = r.db.ScanRows(sql, &client)

			// if any error, create logger error and return error
			if err != nil {
				logger.Error("error scan data client in project " + err.Error())
				return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
			}

			// paste data client from rows in domain project field client
			project.Client = *client

			// append data project to variable array
			projects = append(projects, *project)
		}
		return projects, nil
	}

}

func (r repositoryProjectImpl) GetById(id int32) (*domain.ProjectWithClient, *errs.AppErr) {
	// create variable for domain project
	var project *domain.ProjectWithClient

	// get data project by id and check there error or not
	if rows, err := r.db.Table("project").Joins("inner join client on project.id_client = client.id_client ").Where("project.id_project = ?", id).Rows(); err != nil {

		// check if data project by id not found
		if errors.Is(err, gorm.ErrRecordNotFound) {

			// create logger error for debugging
			logger.Error("error get data project by id : " + err.Error())
			return nil, errs.NewNotFoundError("data project not found")
		} else {

			// this block for handle error unexpected
			// create logger error for debugging
			logger.Error("error get data project by id : " + err.Error())
			return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
		}
	} else {
		// close rows at the end execution program
		defer rows.Close()

		// loop rows project from database
		for rows.Next() {
			// create variable domain project and client
			var client *domain.Client

			// scan data project and save in variable
			err := r.db.ScanRows(rows, &project)

			// if any error, create logger error and return error
			if err != nil {
				logger.Error("error scan data project with client " + err.Error())
				return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
			}

			// scan data client from rows and save in variable
			err = r.db.ScanRows(rows, &client)

			// if any error, create logger error and return error
			if err != nil {
				logger.Error("error scan data client in project " + err.Error())
				return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
			}

			// paste data client from rows in domain project field client
			project.Client = *client

		}
		return project, nil
	}
}

func (r repositoryProjectImpl) Delete(id int32) *errs.AppErr {
	// create variable for domain project
	var project *domain.Project = &domain.Project{IdProject: int32(id)}

	// delete project from database where id from request project browser
	if tx := r.db.Delete(project); tx.RowsAffected < int64(1) {

		// create looger error for debuggin delete failed because data project by id nof found
		logger.Error("error delete project, id not found ")
		return errs.NewNotFoundError("delete failed, project not found")
	} else if tx.Error != nil {

		// create logger error unexpected for debugging and return
		logger.Error("error delete project unexpected " + tx.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}

	return nil
}
