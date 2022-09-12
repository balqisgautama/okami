package dao

import (
	"database/sql"
	serverconfig "okami.auth.backend/config/server"
	"okami.auth.backend/constanta"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	errorModel "okami.auth.backend/model/error"
)

type logAuditSystemDAO struct {
	AbstractDAO
}

var LogAuditSystemDAO = logAuditSystemDAO{}.New()

func (input logAuditSystemDAO) New() (output logAuditSystemDAO) {
	output.FileName = "LogActivitiesDAO.go"
	return
}

// created at 08-22-2022
func (input logAuditSystemDAO) InsertLogAuditSystem(data model.LogAuditSystem) (result model.LogAuditSystem, output res.Payload) {
	funcName := "InsertUser"
	result, err := input.insertLogAuditSystem(serverconfig.ServerAttribute.DBConnection, data)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-22-2022
func (input logAuditSystemDAO) insertLogAuditSystem(db *sql.DB, data model.LogAuditSystem) (result model.LogAuditSystem, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `INSERT INTO log_audit_system (audit_time, audit_detail, ` +
		`resource_client_id, data_old, data_new, action) ` +
		`VALUES ($1, $2, $3, $4, $5, $6) ` +
		`RETURNING *`

	row := db.QueryRow(sqlStatement, data.AuditTime.Time, data.AuditDetail.String, data.ResourceClientID.String,
		data.DataOld.String, data.DataNew.String, data.Action.Int16)

	err.Error = row.Scan(&result.AuditID, &result.AuditTime, &result.AuditDetail, &result.ResourceClientID,
		&result.DataOld, &result.DataNew, &result.Action)

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
