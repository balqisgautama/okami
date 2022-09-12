package dao

import (
	"database/sql"
	serverconfig "okami.auth.backend/config/server"
	"okami.auth.backend/constanta"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	errorModel "okami.auth.backend/model/error"
	"time"
)

type logPKCEDAO struct {
	AbstractDAO
}

var LogPKCEDAO = logPKCEDAO{}.New()

func (input logPKCEDAO) New() (output logPKCEDAO) {
	output.FileName = "LogPKCEDAO.go"
	return
}

// InsertLogPKCE ------------------------------------------------------------------------------------------
// created at 08-27-2022
func (input logPKCEDAO) InsertLogPKCE(codeChallenger string) (result model.LogPKCE, output res.Payload) {
	funcName := "insertLogPKCE"
	result, err := input.insertLogPKCE(serverconfig.ServerAttribute.DBConnection, codeChallenger)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-27-2022
func (input logPKCEDAO) insertLogPKCE(db *sql.DB, codeChallenger string) (result model.LogPKCE, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `INSERT INTO log_pkce (step1, code_challenger) ` +
		`VALUES ($1, $2) ` +
		`RETURNING *`

	row := db.QueryRow(sqlStatement, time.Now(), codeChallenger)

	err.Error = row.Scan(&result.LogPKCEID, &result.Step1, &result.Step2,
		&result.Step3, &result.UserClientID, &result.SecretCode, &result.CodeChallenger)

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

// GetUserByClientID ------------------------------------------------------------------------------------------
// created at 08-30-2022
func (input logPKCEDAO) GetLogPKCEByCodeChallenger(codeChallenger string) (result model.LogPKCE, output res.Payload) {
	funcName := "GetLogPKCEByCodeChallenger"
	result, err := input.getLogPKCEByCodeChallenger(serverconfig.ServerAttribute.DBConnection, codeChallenger)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-30-2022
func (input logPKCEDAO) getLogPKCEByCodeChallenger(db *sql.DB, codeChallenger string) (result model.LogPKCE, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM log_pkce WHERE code_challenger=$1`
	row := db.QueryRow(sqlStatement, codeChallenger)

	err.Error = row.Scan(&result.LogPKCEID, &result.Step1, &result.Step2, &result.Step3,
		&result.UserClientID, &result.SecretCode, &result.CodeChallenger)

	switch err.Error {
	case sql.ErrNoRows:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return
	case nil:
		return
	default:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return
	}
}

// UpdateUserByClientID ------------------------------------------------------------------------------------------
// created at 08-30-2022
func (input logPKCEDAO) UpdateLogPKCEByCodeChallenger(data model.LogPKCE) (result model.LogPKCE, output res.Payload) {
	funcName := "UpdateLogPKCEByCodeChallenger"
	result, output = input.updateLogPKCEByCodeChallenger(serverconfig.ServerAttribute.DBConnection, data)
	if output.Status.Code != "" {
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-30-2022
func (input logPKCEDAO) updateLogPKCEByCodeChallenger(db *sql.DB, data model.LogPKCE) (result model.LogPKCE, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `UPDATE log_pkce SET step1=$2, step2=$3, step3=$4, ` +
		`user_client_id=$5, secret_code=$6 ` +
		`WHERE code_challenger=$1 RETURNING *`
	row := db.QueryRow(sqlStatement, data.CodeChallenger.String, data.Step1.Time, data.Step2.Time,
		data.Step3.Time, data.UserClientID.String, data.SecretCode.String)

	err.Error = row.Scan(&result.LogPKCEID, &result.Step1, &result.Step2, &result.Step3,
		&result.UserClientID, &result.SecretCode, &result.CodeChallenger)

	switch err.Error {
	case sql.ErrNoRows:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return
	case nil:
		return
	default:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return
	}
}

// GetLogPKCEByCodeVerifier ------------------------------------------------------------------------------------------
// created at 08-31-2022
func (input logPKCEDAO) GetLogPKCEByCodeVerifier(codeVerifier string) (result model.LogPKCE, output res.Payload) {
	funcName := "GetLogPKCEByCodeVerifier"
	result, err := input.getLogPKCEByCodeVerifier(serverconfig.ServerAttribute.DBConnection, codeVerifier)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-31-2022
func (input logPKCEDAO) getLogPKCEByCodeVerifier(db *sql.DB, codeVerifier string) (result model.LogPKCE, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM log_pkce WHERE secret_code=$1`
	row := db.QueryRow(sqlStatement, codeVerifier)

	err.Error = row.Scan(&result.LogPKCEID, &result.Step1, &result.Step2, &result.Step3,
		&result.UserClientID, &result.SecretCode, &result.CodeChallenger)

	switch err.Error {
	case sql.ErrNoRows:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return
	case nil:
		return
	default:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return
	}
}
