package dao

import (
	"database/sql"
	serverconfig "okami.auth.backend/config/server"
	"okami.auth.backend/constanta"
	res "okami.auth.backend/dto/out"
	errorModel "okami.auth.backend/model/error"
	model "okami.auth.backend/model/user"
)

type activationAccountDAO struct {
	AbstractDAO
}

var ActivationAccountDAO = activationAccountDAO{}.New()

func (input activationAccountDAO) New() (output activationAccountDAO) {
	output.FileName = "ActivationAccountDAO.go"
	return
}

// created at 06-20-2022
// updated at 06-21-2022
func (input activationAccountDAO) InsertData(data model.ActivationAccount) (id int64, output res.Payload) {
	funcName := "InsertData"
	id, err := input.insertData(serverconfig.ServerAttribute.DBConnection, data)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 06-20-2022
// updated at 06-21-2022
func (input activationAccountDAO) insertData(db *sql.DB, data model.ActivationAccount) (id int64, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `INSERT INTO activation_account (counter, code, email_to, expire, 
		status, link_validate, link_resend) ` +
		`VALUES ($1, $2, $3, $4, $5, $6, $7) ` +
		`RETURNING activation_id`

	err.Error = db.QueryRow(sqlStatement, data.Counter.Int64, data.Code.Int64, data.EmailTo.String,
		data.Expire.Int64, data.Status.Int64, data.LinkValidate.String, data.LinkResend.String).Scan(&id)

	if err.Error != nil {
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return
	}
	return
}

// created at 06-21-2022
func (input activationAccountDAO) GetActivationByID(id int64) (data model.ActivationAccount, output res.Payload) {
	funcName := "GetActivationByID"
	data, err := input.getActivationByID(serverconfig.ServerAttribute.DBConnection, id)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 06-21-2022
func (input activationAccountDAO) getActivationByID(db *sql.DB, id int64) (data model.ActivationAccount, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM activation_account WHERE activation_id=$1`
	row := db.QueryRow(sqlStatement, id)

	err.Error = row.Scan(&data.ActivationID, &data.Counter, &data.Code,
		&data.EmailTo, &data.Counter, &data.Expire, &data.LinkValidate, &data.LinkResend)

	switch err.Error {
	case sql.ErrNoRows:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return data, output
	case nil:
		return data, output
	default:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return data, output
	}
}

// created at 06-22-2022
func (input activationAccountDAO) UpdateStatusByID(id int64) (rowsAffected int64, output res.Payload) {
	funcName := "UpdateStatusByID"
	rowsAffected, output = input.updateStatusByID(serverconfig.ServerAttribute.DBConnection, id)
	if output.Status.Code != "" {
		output.Status.Detail = funcName
		return rowsAffected, output
	}
	return rowsAffected, output
}

// created at 06-22-2022
func (input activationAccountDAO) updateStatusByID(db *sql.DB, id int64) (rowsAffected int64, output res.Payload) {
	sqlStatement := `UPDATE activation_account 
				SET status=$2
				WHERE activation_id=$1`
	row, errs := db.Exec(sqlStatement, id, constanta.UserActive)
	if errs != nil {
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = errs.Error()
		return 0, output
	}
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error()
		return 0, output
	}
	return rowsAffected, output
}
