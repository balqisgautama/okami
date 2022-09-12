package dao

import (
	"database/sql"
	serverconfig "okami.auth.backend/config/server"
	"okami.auth.backend/constanta"
	req "okami.auth.backend/dto/in"
	res "okami.auth.backend/dto/out"
	model "okami.auth.backend/model"
	errorModel "okami.auth.backend/model/error"
	"okami.auth.backend/util"
	"time"
)

type resourceDAO struct {
	AbstractDAO
}

var ResourceDAO = resourceDAO{}.New()

func (input resourceDAO) New() (output resourceDAO) {
	output.FileName = "UserDAO.go"
	return
}

// created at 08-16-2022
func (input resourceDAO) InsertResource(resource model.ResourceGeneral) (result model.ResourceGeneral, output res.Payload) {
	funcName := "InsertResource"
	result, err := input.insertResource(serverconfig.ServerAttribute.DBConnection, resource)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-16-2022
// updated at 08-18-2022
func (input resourceDAO) insertResource(db *sql.DB, resource model.ResourceGeneral) (result model.ResourceGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `INSERT INTO resources (client_id, surname, nickname, access_to, created_at) ` +
		`VALUES ($1, $2, $3, $4, $5) ` +
		`RETURNING *`

	row := db.QueryRow(sqlStatement, resource.ClientID.String, resource.Surname.String,
		resource.Nickname.String, resource.AccessTo.String, resource.CreatedAt.Time)

	err.Error = row.Scan(&result.ResourceID, &result.ClientID,
		&result.Surname, &result.Nickname, &result.AccessTo, &result.CreatedAt,
		&result.UpdatedAt, &result.DeletedAt)

	if err.Error != nil {
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return
	}
	return
}

// created at 08-16-2022
func (input resourceDAO) GetResources() (resources []model.ResourceGeneral, output res.Payload) {
	funcName := "GetResources"
	resources, err := input.getResources(serverconfig.ServerAttribute.DBConnection)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-16-2022
// updated at 08-19-2022
func (input resourceDAO) getResources(db *sql.DB) (resources []model.ResourceGeneral, output res.Payload) {
	var rows *sql.Rows
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM resources WHERE EXTRACT(EPOCH FROM deleted_at) is NULL`
	rows, err.Error = db.Query(sqlStatement)

	for rows.Next() {
		var resource model.ResourceGeneral
		err.Error = rows.Scan(&resource.ResourceID, &resource.ClientID,
			&resource.Surname, &resource.Nickname, &resource.AccessTo, &resource.CreatedAt,
			&resource.UpdatedAt, &resource.DeletedAt)
		resources = append(resources, resource)
	}

	switch err.Error {
	case sql.ErrNoRows:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.CausedBy.Error()
		return resources, output
	case nil:
		return resources, output
	default:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return resources, output
	}
}

// created at 08-18-2022
func (input resourceDAO) GetResourceByClientID(clientID string) (result model.ResourceGeneral, output res.Payload) {
	funcName := "GetResourceByClientID"
	result, err := input.getResourceByClientID(serverconfig.ServerAttribute.DBConnection, clientID)
	if err.Status.Code != "" {
		output.Status.Code = err.Status.Code
		output.Status.Message = err.Status.Message
		output.Status.Detail = funcName
		return
	}
	return
}

// created at 08-18-2022
func (input resourceDAO) getResourceByClientID(db *sql.DB, clientID string) (result model.ResourceGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `SELECT * FROM resources WHERE client_id=$1 AND 
                          EXTRACT(EPOCH FROM deleted_at) is NULL`
	row := db.QueryRow(sqlStatement, clientID)

	err.Error = row.Scan(&result.ResourceID, &result.ClientID,
		&result.Surname, &result.Nickname, &result.AccessTo, &result.CreatedAt,
		&result.UpdatedAt, &result.DeletedAt)

	switch err.Error {
	case sql.ErrNoRows:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return result, output
	case nil:
		return result, output
	default:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return result, output
	}
}

// created at 08-18-2022
func (input resourceDAO) UpdateResourceByClientID(data req.ReqResourceUpdate) (result model.ResourceGeneral, output res.Payload) {
	funcName := "UpdateResourceByClientID"
	result, output = input.updateResourceByClientID(serverconfig.ServerAttribute.DBConnection, data)
	if output.Status.Code != "" {
		output.Status.Detail = funcName
		return result, output
	}
	return result, output
}

// created at 08-18-2022
func (input resourceDAO) updateResourceByClientID(db *sql.DB, data req.ReqResourceUpdate) (result model.ResourceGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `UPDATE resources SET surname=$2, access_to=$3, updated_at=$4 WHERE client_id=$1 
					RETURNING *`
	row := db.QueryRow(sqlStatement, data.ClientID, data.Surname, data.AccessTo, time.Now())

	err.Error = row.Scan(&result.ResourceID, &result.ClientID,
		&result.Surname, &result.Nickname, &result.AccessTo, &result.CreatedAt,
		&result.UpdatedAt, &result.DeletedAt)

	switch err.Error {
	case sql.ErrNoRows:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return result, output
	case nil:
		return result, output
	default:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return result, output
	}
}

// created at 08-19-2022
func (input resourceDAO) SoftDeleteResource(clientID string) (result model.ResourceGeneral, output res.Payload) {
	funcName := "SoftDeleteResource"
	result, output = input.softDeleteResource(serverconfig.ServerAttribute.DBConnection, clientID)
	if output.Status.Code != "" {
		output.Status.Detail = funcName
		return result, output
	}
	return result, output
}

// created at 08-19-2022
func (input resourceDAO) softDeleteResource(db *sql.DB, clientID string) (result model.ResourceGeneral, output res.Payload) {
	var err errorModel.ErrorModel

	sqlStatement := `UPDATE resources SET deleted_at=$2, nickname=CONCAT(nickname, '` +
		constanta.PrefixDataDeleted + util.RandomString(7) +
		`') WHERE client_id=$1  RETURNING *`
	row := db.QueryRow(sqlStatement, clientID, time.Now())

	err.Error = row.Scan(&result.ResourceID, &result.ClientID,
		&result.Surname, &result.Nickname, &result.AccessTo, &result.CreatedAt,
		&result.UpdatedAt, &result.DeletedAt)

	switch err.Error {
	case sql.ErrNoRows:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return result, output
	case nil:
		return result, output
	default:
		output.Status.Code = constanta.CodeDBServerError
		output.Status.Message = err.Error.Error()
		return result, output
	}
}
