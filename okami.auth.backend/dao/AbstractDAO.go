package dao

import (
	"database/sql"
	"okami.auth.backend/constanta"
	res "okami.auth.backend/dto/out"
)

type AbstractDAO struct {
	FileName           string
	TableName          string
	ElasticSearchIndex string
}

type FieldStatus struct {
	IsCheck   bool
	FieldName string
	Value     interface{}
}

type DefaultFieldMustCheck struct {
	ID        FieldStatus
	Deleted   FieldStatus
	Status    FieldStatus
	CreatedBy FieldStatus
}

//func (input AbstractDAO) SoftDelete(db *sql.Tx, tableName string, data repository.SoftDelete) (err errorModel.ErrorModel) {
//	funcName := "SoftDelete"
//
//	query := "UPDATE " + tableName + " " +
//		"SET" +
//		"	deleted = $1," +
//		"	updated_client = $2," +
//		"	updated_at = $3," +
//		"	updated_by = $4 " +
//		"WHERE" +
//		"	id = $5 AND " +
//		"	deleted = false"
//
//	stmt, dbError := db.Prepare(query)
//
//	if dbError != nil {
//		return errorModel.GenerateInternalDBServerError(input.FileName, funcName, dbError)
//	}
//
//	result, dbError := stmt.Exec(
//		data.Deleted.Bool,
//		data.UpdatedClient.String,
//		data.UpdatedAt.Time,
//		data.UpdatedBy.Int64,
//		data.Id.Int64)
//
//	if dbError != nil {
//		return errorModel.GenerateInternalDBServerError(input.FileName, funcName, dbError)
//	}
//
//	rowsAffected, rowsAffectedError := result.RowsAffected()
//	if rowsAffectedError != nil || rowsAffected < 1 {
//		return errorModel.GenerateInternalDBServerError(input.FileName, funcName, rowsAffectedError)
//	}
//
//	return errorModel.GenerateNonErrorModel()
//}

// updated at 08-18-2022
func (input AbstractDAO) IsDataExist(db *sql.DB, tableName string, id int64, colomIDName string) (result bool, output res.Payload) {
	funcName := "IsDataExist"
	query := "SELECT " +
		"CASE WHEN count(id) > 0 " +
		"	THEN TRUE " +
		"	ELSE FALSE " +
		"END " +
		"FROM " + tableName + " WHERE is NULL " +
		"AND " + colomIDName + " = $1"

	dbError := db.QueryRow(query, id).Scan(&result) // returnnya kolom case dengan value true/false
	if dbError != nil && dbError.Error() != "sql: no rows in result set" {
		output.Status.Code = constanta.CodeGetDataFailed
		output.Status.Message = dbError.Error()
		output.Status.Detail = funcName
		return
	}

	return
}
