package dao

import (
	"database/sql"
	serverconfig "okami.auth.backend/config/server"
	"okami.auth.backend/constanta"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	errorModel "okami.auth.backend/model/error"
)

type userActivationDao struct {
	AbstractDAO
}

var UserActivationDao = userActivationDao{}.New()

func (input userActivationDao) New() (output userActivationDao) {
	output.FileName = "ActivationAccountDAO.go"
	return
}

// InsertActivation ------------------------------------------------------------------------------------------
// created at 08-22-2022
func (input userActivationDao) InsertActivation(user model.ActivationUser) (result model.ActivationUser, output res.Payload) {
	funcName := "InsertActivation"
	result, err := input.insertActivation(serverconfig.ServerAttribute.DBConnection, user)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-22-2022
func (input userActivationDao) insertActivation(db *sql.DB, data model.ActivationUser) (result model.ActivationUser, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `INSERT INTO user_activation (code, expired_at, user_id, email_to, ` +
		`email_link_validate, email_link_resend) ` +
		`VALUES ($1, $2, $3, $4, $5, $6) ` +
		`RETURNING *`

	row := db.QueryRow(sqlStatement, data.Code.String, data.ExpiredAt.Time, data.UserID.Int64,
		data.EmailTo.String, data.EmailLinkValidate.String, data.EmailLinkResend.String)

	err.Error = row.Scan(&result.ActivationID, &result.Counter, &result.Code, &result.ExpiredAt,
		&result.Status, &result.UserID, &result.EmailTo, &result.EmailLinkValidate, &result.EmailLinkResend)

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

// GetActivationByUserID ------------------------------------------------------------------------------------------
// created at 08-23-2022
func (input userActivationDao) GetActivationByUserID(userID int64) (result model.ActivationUser, output res.Payload) {
	funcName := "GetActivationByUserID"
	result, err := input.getActivationByUserID(serverconfig.ServerAttribute.DBConnection, userID)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-23-2022
func (input userActivationDao) getActivationByUserID(db *sql.DB, userID int64) (result model.ActivationUser, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM user_activation WHERE user_id=$1 AND status=$2`
	row := db.QueryRow(sqlStatement, userID, constanta.UserPending)

	err.Error = row.Scan(&result.ActivationID, &result.Counter, &result.Code, &result.ExpiredAt,
		&result.Status, &result.UserID, &result.EmailTo, &result.EmailLinkValidate, &result.EmailLinkResend)

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

// UpdateActivationByActivationID ------------------------------------------------------------------------------------------
// created at 08-23-2022
func (input userActivationDao) UpdateActivationByActivationID(data model.ActivationUser) (result model.ActivationUser, output res.Payload) {
	funcName := "UpdateActivationByActivationID"
	result, output = input.updateActivationByActivationID(serverconfig.ServerAttribute.DBConnection, data)
	if output.Status.Code != "" {
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-23-2022
func (input userActivationDao) updateActivationByActivationID(db *sql.DB, data model.ActivationUser) (result model.ActivationUser, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `UPDATE user_activation SET counter_regenerate=$2, code=$3, expired_at=$4, ` +
		`status=$5, email_to=$6, email_link_validate=$7, email_link_resend=$8 ` +
		`WHERE activation_id=$1 RETURNING *`
	row := db.QueryRow(sqlStatement, data.ActivationID.Int64, data.Counter.Int64, data.Code.String,
		data.ExpiredAt.Time, data.Status.Int64, data.EmailTo.String, data.EmailLinkValidate.String,
		data.EmailLinkResend.String)

	err.Error = row.Scan(&result.ActivationID, &result.Counter, &result.Code, &result.ExpiredAt,
		&result.Status, &result.UserID, &result.EmailTo, &result.EmailLinkValidate, &result.EmailLinkResend)

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
