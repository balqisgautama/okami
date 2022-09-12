package dao

import (
	"database/sql"
	serverconfig "okami.auth.backend/config/server"
	"okami.auth.backend/constanta"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	errorModel "okami.auth.backend/model/error"
)

type logActivitiesDAO struct {
	AbstractDAO
}

var LogActivitiesDAO = logActivitiesDAO{}.New()

func (input logActivitiesDAO) New() (output logActivitiesDAO) {
	output.FileName = "LogActivitiesDAO.go"
	return
}

// created at 08-22-2022
func (input logActivitiesDAO) InsertLogActivity(data model.LogActivity) (result model.LogActivity, output res.Payload) {
	funcName := "InsertLogActivity"
	result, err := input.insertLogActivity(serverconfig.ServerAttribute.DBConnection, data)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-22-2022
func (input logActivitiesDAO) insertLogActivity(db *sql.DB, data model.LogActivity) (result model.LogActivity, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `INSERT INTO log_activities (activity_time, activity_detail, resource_client_id) ` +
		`VALUES ($1, $2, $3) ` +
		`RETURNING *`

	row := db.QueryRow(sqlStatement, data.ActivityTime.Time, data.ActivityDetail.String, data.ResourceClientID.String)

	err.Error = row.Scan(&result.ActivityID, &result.ActivityTime, &result.ActivityDetail, &result.ResourceClientID)

	switch err.Error {
	case sql.ErrNoRows:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.CausedBy.Error()
		return
	case nil:
		return
	default:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return
	}
	return
}
