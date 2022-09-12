package util

import (
	req "okami.auth.backend/dto/in"
	response "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	uuid "okami.auth.backend/util"
	"strings"
	"time"
)

// created at 08-16-2022
// updated at 08-18-2022
func ReqResourceCreateToResourceForInsertDB(resource req.ReqResourceCreate) (result model.ResourceGeneral) {
	result.ClientID.String = uuid.GetUUID()
	result.Nickname.String = strings.ToLower(resource.Nickname)
	result.Surname.String = resource.Surname
	result.AccessTo.String = strings.ToLower(resource.AccessTo)
	result.CreatedAt.Time = time.Now()
	return
}

// created at 08-18-2022
func DBDataToResouceGeneral(dataDB model.ResourceGeneral) (result response.ResourceGeneral) {
	result.ResourceID = dataDB.ResourceID.Int64
	result.ClientID = dataDB.ClientID.String
	result.Surname = dataDB.Surname.String
	result.Nickname = dataDB.Nickname.String
	result.AccessTo = dataDB.AccessTo.String
	result.CreatedAt = dataDB.CreatedAt.Time
	result.UpdatedAt = dataDB.UpdatedAt.Time
	result.DeletedAt = dataDB.DeletedAt.Time
	return
}

// created at 08-18-2022
func DBDataToResouceGeneralArray(dataDB []model.ResourceGeneral) (result []response.ResourceGeneral) {
	for _, value := range dataDB {
		temp := DBDataToResouceGeneral(value)
		result = append(result, temp)
	}
	return
}

// created at 08-18-2022
func ReqResourceUpdateToResourceForInsertDB(resource req.ReqResourceUpdate) (result model.ResourceGeneral) {
	result.ClientID.String = resource.ClientID
	result.Surname.String = resource.Surname
	result.AccessTo.String = resource.AccessTo
	result.UpdatedAt.Time = time.Now()
	return
}
