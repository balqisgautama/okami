package dao

import (
	"database/sql"
	serverconfig "okami.auth.backend/config/server"
	"okami.auth.backend/constanta"
	req "okami.auth.backend/dto/in"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	errorModel "okami.auth.backend/model/error"
	"okami.auth.backend/util"
	"strconv"
	"time"
)

type userDAO struct {
	AbstractDAO
}

var UserDAO = userDAO{}.New()

func (input userDAO) New() (output userDAO) {
	output.FileName = "UserDAO.go"
	return
}

// InsertUser ------------------------------------------------------------------------------------------
// created at 06-17-2022
// updated at 08-19-2022
func (input userDAO) InsertUser(user model.UserGeneral) (result model.UserGeneral, output res.Payload) {
	funcName := "InsertUser"
	result, err := input.insertUser(serverconfig.ServerAttribute.DBConnection, user)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 06-17-2022
// updated at 08-19-2022
func (input userDAO) insertUser(db *sql.DB, user model.UserGeneral) (result model.UserGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `INSERT INTO users (username, email, password, client_id, status, locale, 
		created_by, created_client, created_at) ` +
		`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ` +
		`RETURNING *`

	row := db.QueryRow(sqlStatement, user.Username.String, user.Email.String, user.Password.String,
		user.ClientID.String, user.Status.Int64, user.Locale.String,
		user.CreatedBy.Int64, user.CreatedClient.String, user.CreatedAt.Time)

	err.Error = row.Scan(&result.UserID, &result.Username, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.Locale, &result.AdditionalInfo, &result.LastToken, &result.CreatedBy,
		&result.CreatedClient, &result.CreatedAt, &result.UpdatedBy, &result.UpdatedClient, &result.UpdatedAt,
		&result.DeletedBy, &result.DeletedClient, &result.DeletedAt)

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

// GetUsers ------------------------------------------------------------------------------------------
// created at 08-19-2022
func (input userDAO) GetUsers() (result []model.UserGeneral, output res.Payload) {
	funcName := "GetUsers"
	result, err := input.getUsers(serverconfig.ServerAttribute.DBConnection)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-19-2022
func (input userDAO) getUsers(db *sql.DB) (users []model.UserGeneral, output res.Payload) {
	var rows *sql.Rows
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM users WHERE EXTRACT(EPOCH FROM deleted_at) is NULL`
	rows, err.Error = db.Query(sqlStatement)

	for rows.Next() {
		var user model.UserGeneral
		err.Error = rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.ClientID,
			&user.Status, &user.Locale, &user.AdditionalInfo, &user.LastToken, &user.CreatedBy,
			&user.CreatedClient, &user.CreatedAt, &user.UpdatedBy, &user.UpdatedClient, &user.UpdatedAt,
			&user.DeletedBy, &user.DeletedClient, &user.DeletedAt)
		users = append(users, user)
	}

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
}

// GetUserByClientID ------------------------------------------------------------------------------------------
// created at 08-19-2022
func (input userDAO) GetUserByClientID(clientID string) (result model.UserGeneral, output res.Payload) {
	funcName := "GetUserByClientID"
	result, err := input.getUserByClientID(serverconfig.ServerAttribute.DBConnection, clientID)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-19-2022
func (input userDAO) getUserByClientID(db *sql.DB, clientID string) (result model.UserGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM users WHERE client_id=$1 AND ` +
		`EXTRACT(EPOCH FROM deleted_at) is NULL`
	row := db.QueryRow(sqlStatement, clientID)

	err.Error = row.Scan(&result.UserID, &result.Username, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.Locale, &result.AdditionalInfo, &result.LastToken, &result.CreatedBy,
		&result.CreatedClient, &result.CreatedAt, &result.UpdatedBy, &result.UpdatedClient, &result.UpdatedAt,
		&result.DeletedBy, &result.DeletedClient, &result.DeletedAt)

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
// created at 08-19-2022
// updated at 08-31-2022
func (input userDAO) UpdateUserByClientID(data model.UserGeneral) (result model.UserGeneral, output res.Payload) {
	funcName := "UpdateUserByClientID"
	result, output = input.updateUserByClientID(serverconfig.ServerAttribute.DBConnection, data)
	if output.Status.Code != "" {
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-19-2022
// updated at 09-02-2022
func (input userDAO) updateUserByClientID(db *sql.DB, data model.UserGeneral) (result model.UserGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `UPDATE users SET username=$2, email=$3, password=$4, ` +
		`updated_by=$5, updated_client=$6, updated_at=$7, last_token=$8, status=$9 ` +
		`WHERE client_id=$1 RETURNING *`
	row := db.QueryRow(sqlStatement, data.ClientID.String, data.Username.String, data.Email.String,
		data.Password.String, data.UpdatedBy.Int64, data.UpdatedClient.String, time.Now(),
		data.LastToken.Time, data.Status.Int64)

	err.Error = row.Scan(&result.UserID, &result.Username, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.Locale, &result.AdditionalInfo, &result.LastToken, &result.CreatedBy,
		&result.CreatedClient, &result.CreatedAt, &result.UpdatedBy, &result.UpdatedClient, &result.UpdatedAt,
		&result.DeletedBy, &result.DeletedClient, &result.DeletedAt)

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

// SoftDeleteUser ------------------------------------------------------------------------------------------
// created at 08-19-2022
func (input userDAO) SoftDeleteUser(data req.ReqUserGeneral) (result model.UserGeneral, output res.Payload) {
	funcName := "SoftDeleteUser"
	result, output = input.softDeleteUser(serverconfig.ServerAttribute.DBConnection, data)
	if output.Status.Code != "" {
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-19-2022
func (input userDAO) softDeleteUser(db *sql.DB, data req.ReqUserGeneral) (result model.UserGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `UPDATE users SET deleted_at=$2, deleted_by=$3, deleted_client=$4, ` +
		`username=CONCAT(username, '` + constanta.PrefixDataDeleted + util.RandomString(7) + `'), ` +
		`email=CONCAT(email, '` + constanta.PrefixDataDeleted + util.RandomString(7) + `') ` +
		`WHERE client_id=$1  RETURNING *`
	row := db.QueryRow(sqlStatement, data.UserClientID, time.Now(), data.ResourceID, data.ResourceClientID)

	err.Error = row.Scan(&result.UserID, &result.Username, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.Locale, &result.AdditionalInfo, &result.LastToken, &result.CreatedBy,
		&result.CreatedClient, &result.CreatedAt, &result.UpdatedBy, &result.UpdatedClient, &result.UpdatedAt,
		&result.DeletedBy, &result.DeletedClient, &result.DeletedAt)

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

// GetUserByEmailUsername ------------------------------------------------------------------------------------------
// created at 08-23-2022
func (input userDAO) GetUserByEmailUsername(email, username string) (result model.UserGeneral, output res.Payload) {
	funcName := "GetUserByEmailUsername"
	result, err := input.getUserByEmailUsername(serverconfig.ServerAttribute.DBConnection, email, username)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-23-2022
func (input userDAO) getUserByEmailUsername(db *sql.DB, email, username string) (result model.UserGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM users WHERE email=$1 AND username=$2 AND ` +
		`EXTRACT(EPOCH FROM deleted_at) is NULL`
	row := db.QueryRow(sqlStatement, email, username)

	err.Error = row.Scan(&result.UserID, &result.Username, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.Locale, &result.AdditionalInfo, &result.LastToken, &result.CreatedBy,
		&result.CreatedClient, &result.CreatedAt, &result.UpdatedBy, &result.UpdatedClient, &result.UpdatedAt,
		&result.DeletedBy, &result.DeletedClient, &result.DeletedAt)

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

// GetUserActiveByUsernamePassword ------------------------------------------------------------------------------------------
// created at 08-27-2022
func (input userDAO) GetUserActiveByUsernamePassword(username string, password string) (result model.UserGeneral, output res.Payload) {
	funcName := "GetUserByUsername"
	result, err := input.getUserActiveByUsernamePassword(serverconfig.ServerAttribute.DBConnection, username, password)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-27-2022
func (input userDAO) getUserActiveByUsernamePassword(db *sql.DB, username string, password string) (result model.UserGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM users WHERE username=$1 AND ` +
		`EXTRACT(EPOCH FROM deleted_at) is NULL AND ` +
		`status=` + strconv.Itoa(constanta.UserActive)
	row := db.QueryRow(sqlStatement, username)

	err.Error = row.Scan(&result.UserID, &result.Username, &result.Email, &result.Password, &result.ClientID,
		&result.Status, &result.Locale, &result.AdditionalInfo, &result.LastToken, &result.CreatedBy,
		&result.CreatedClient, &result.CreatedAt, &result.UpdatedBy, &result.UpdatedClient, &result.UpdatedAt,
		&result.DeletedBy, &result.DeletedClient, &result.DeletedAt)

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

//// created at 06-20-2022
//func (input userDAO) GetUsersByEmailUsername(username string, email string) (users []model.UserGeneral, output res.Payload) {
//	funcName := "GetUserDataByEmailUsername"
//	users, err := input.getUsersByEmailUsername(serverconfig.ServerAttribute.DBConnection, username, email)
//	if err.Status.Code != "" {
//		output.Status.Code = err.Status.Code
//		output.Status.Message = err.Status.Message
//		output.Status.Detail = funcName
//		return
//	}
//	return
//}
//
//// created at 06-20-2022
//func (input userDAO) getUsersByEmailUsername(db *sql.DB, username string, email string) (users []model.UserGeneral, output res.Payload) {
//	var rows *sql.Rows
//	var err errorModel.ErrorModel
//
//	sqlStatement := `SELECT * FROM users WHERE username=$1 OR email=$2`
//	rows, err.Error = db.Query(sqlStatement, username, email)
//
//	for rows.Next() {
//		var user model.UserGeneral
//		err.Error = rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.ClientID,
//			&user.Status, &user.ActivationID, &user.Locale, &user.AdditionalInfo, &user.LastToken, &user.CreatedBy,
//			&user.CreatedClient, &user.CreatedAt, &user.UpdatedBy, &user.UpdatedClient, &user.UpdatedAt)
//		users = append(users, user)
//	}
//
//	switch err.Error {
//	case sql.ErrNoRows:
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = err.CausedBy.Error()
//		return users, output
//	case nil:
//		return users, output
//	default:
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = err.Error.Error()
//		return users, output
//	}
//}
//
//// created at 06-20-2022
//func (input userDAO) GetUserByUserID(id int64) (user model.UserGeneral, output res.Payload) {
//	funcName := "GetUserByUserID"
//	user, err := input.getUserByUserID(serverconfig.ServerAttribute.DBConnection, id)
//	if err.Status.Code != "" {
//		output.Status.Code = err.Status.Code
//		output.Status.Message = err.Status.Message
//		output.Status.Detail = funcName
//		return
//	}
//	return
//}
//
//// created at 06-20-2022
//func (input userDAO) getUserByUserID(db *sql.DB, id int64) (user model.UserGeneral, output res.Payload) {
//	var err errorModel.ErrorModel
//
//	sqlStatement := `SELECT * FROM users WHERE user_id=$1`
//	row := db.QueryRow(sqlStatement, id)
//
//	err.Error = row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.ClientID,
//		&user.Status, &user.ActivationID, &user.Locale, &user.AdditionalInfo, &user.LastToken, &user.CreatedBy,
//		&user.CreatedClient, &user.CreatedAt, &user.UpdatedBy, &user.UpdatedClient, &user.UpdatedAt)
//
//	switch err.Error {
//	case sql.ErrNoRows:
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = err.Error.Error()
//		return user, output
//	case nil:
//		return user, output
//	default:
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = err.Error.Error()
//		return user, output
//	}
//}
//
//// created at 06-20-2022
//func (input userDAO) GetUserByEmailUsername(username string, email string) (user model.UserGeneral, output res.Payload) {
//	funcName := "GetUserByEmailUsername"
//	user, err := input.getUserByEmailUsername(serverconfig.ServerAttribute.DBConnection, username, email)
//	if err.Status.Code != "" {
//		output.Status.Code = err.Status.Code
//		output.Status.Message = err.Status.Message
//		output.Status.Detail = funcName
//		return
//	}
//	return
//}
//
//// created at 06-20-2022
//func (input userDAO) getUserByEmailUsername(db *sql.DB, username string, email string) (user model.UserGeneral, output res.Payload) {
//	var err errorModel.ErrorModel
//
//	sqlStatement := `SELECT * FROM users WHERE username=$1 AND email=$2`
//	row := db.QueryRow(sqlStatement, username, email)
//
//	err.Error = row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.ClientID,
//		&user.Status, &user.ActivationID, &user.Locale, &user.AdditionalInfo, &user.LastToken, &user.CreatedBy,
//		&user.CreatedClient, &user.CreatedAt, &user.UpdatedBy, &user.UpdatedClient, &user.UpdatedAt)
//
//	switch err.Error {
//	case sql.ErrNoRows:
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = err.Error.Error()
//		return user, output
//	case nil:
//		return user, output
//	default:
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = err.Error.Error()
//		return user, output
//	}
//}
//
//// created at 06-22-2022
//func (input userDAO) UpdateStatusByUserID(id int64) (rowsAffected int64, output res.Payload) {
//	funcName := "UpdateStatusByUserID"
//	rowsAffected, output = input.updateStatusByUserID(serverconfig.ServerAttribute.DBConnection, id)
//	if output.Status.Code != "" {
//		output.Status.Detail = funcName
//		return rowsAffected, output
//	}
//	return rowsAffected, output
//}
//
//// created at 06-22-2022
//// updated at 07-05-2022
//func (input userDAO) updateStatusByUserID(db *sql.DB, id int64) (rowsAffected int64, output res.Payload) {
//	sqlStatement := `UPDATE users
//				SET status=$2, additional_info=$3
//				WHERE user_id=$1`
//	row, errs := db.Exec(sqlStatement, id, constanta.UserActive, constanta.UserHasAlreadyActivated)
//	if errs != nil {
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = errs.Error()
//		return 0, output
//	}
//	rowsAffected, err := row.RowsAffected()
//	if err != nil {
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = err.Error()
//		return 0, output
//	}
//	return rowsAffected, output
//}
//
//// created at 07-05-2022
//func (input userDAO) GetUserByUsername(username string) (user model.UserGeneral, output res.Payload) {
//	funcName := "GetUserByUsername"
//	user, err := input.getUserByUsername(serverconfig.ServerAttribute.DBConnection, username)
//	if err.Status.Code != "" {
//		output.Status.Code = err.Status.Code
//		output.Status.Message = err.Status.Message
//		output.Status.Detail = funcName
//		return
//	}
//	return
//}
//
//// created at 07-05-2022
//func (input userDAO) getUserByUsername(db *sql.DB, username string) (user model.UserGeneral, output res.Payload) {
//	var err errorModel.ErrorModel
//
//	sqlStatement := `SELECT * FROM users WHERE username=$1`
//	row := db.QueryRow(sqlStatement, username)
//
//	err.Error = row.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.ClientID,
//		&user.Status, &user.ActivationID, &user.Locale, &user.AdditionalInfo, &user.LastToken, &user.CreatedBy,
//		&user.CreatedClient, &user.CreatedAt, &user.UpdatedBy, &user.UpdatedClient, &user.UpdatedAt)
//
//	switch err.Error {
//	case sql.ErrNoRows:
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = err.Error.Error()
//		return user, output
//	case nil:
//		return user, output
//	default:
//		output.Status.Code = constanta.CodeDBServerError
//		output.Status.Message = err.Error.Error()
//		return user, output
//	}
//}
