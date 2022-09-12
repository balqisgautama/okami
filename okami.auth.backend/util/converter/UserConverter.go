package util

import (
	"okami.auth.backend/constanta"
	req "okami.auth.backend/dto/in"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	"okami.auth.backend/util"
	"time"
)

// created at 08-19-2022
func ReqUserCreateToUserForInsertDB(user req.ReqUserGeneral) (result model.UserGeneral) {
	result.Username.String = user.Username
	result.Email.String = user.Email
	password, _ := util.HashPassword(user.Password)
	result.Password.String = password
	result.ClientID.String = util.GetUUID()
	result.Status.Int64 = constanta.UserPending
	result.Locale.String = constanta.DEFAULT_LANGUAGE
	result.CreatedBy.Int64 = user.ResourceID
	result.CreatedClient.String = user.ResourceClientID
	result.CreatedAt.Time = time.Now()
	return
}

// created at 08-19-2022
func DBDataToUserGeneral(user model.UserGeneral) (result res.UserGeneral) {
	result.UserID = user.UserID.Int64
	result.Username = user.Username.String
	result.Email = user.Email.String
	//result.Password = user.Password.String
	result.ClientID = user.ClientID.String
	result.Status = user.Status.Int64
	result.Locale = user.Locale.String
	result.AdditionalInfo = user.AdditionalInfo.String
	result.LastToken = user.LastToken.Time
	result.CreatedBy = user.CreatedBy.Int64
	result.CreatedClient = user.CreatedClient.String
	result.CreatedAt = user.CreatedAt.Time
	result.UpdatedBy = user.UpdatedBy.Int64
	result.UpdatedClient = user.UpdatedClient.String
	result.UpdatedAt = user.UpdatedAt.Time
	result.DeletedBy = user.DeletedBy.Int64
	result.DeletedClient = user.DeletedClient.String
	result.DeletedAt = user.DeletedAt.Time
	return
}

// created at 08-19-2022
func DBDataToUserGeneralArray(dataDB []model.UserGeneral) (result []res.UserGeneral) {
	for _, value := range dataDB {
		temp := DBDataToUserGeneral(value)
		result = append(result, temp)
	}
	return
}

// created at 08-31-2022
func UserGeneralModelToReqUserGeneral(user model.UserGeneral) (result req.ReqUserGeneral) {
	result.UserClientID = user.ClientID.String
	result.Email = user.Email.String
	result.Password = user.Password.String
	result.Username = user.Username.String
	result.ResourceID = user.UpdatedBy.Int64
	result.ResourceClientID = user.UpdatedClient.String
	return
}

// created at 08-31-2022
func ReqUserGeneralToUserGeneralModel(reqUser req.ReqUserGeneral) (result model.UserGeneral) {
	result.Username.String = reqUser.Username
	result.Password.String = reqUser.Password
	result.Email.String = reqUser.Email
	result.UpdatedBy.Int64 = reqUser.ResourceID
	result.UpdatedClient.String = reqUser.ResourceClientID
	result.ClientID.String = reqUser.UserClientID
	return
}

//
//// created at 06-17-2022
//// updated at 06-24-2022
//func RegistToUserGeneral(registData req.ReqRegist, activationID int64) (user model.UserGeneral) {
//	user.Username.String = registData.Username
//	user.Email.String = registData.Email
//	password, _ := util.HashPassword(registData.Password)
//	user.Password.String = password
//	user.ClientID.String = util.GetUUID()
//	user.Status.Int64 = constanta.UserPending
//	//user.ActivationID.Int64 = activationID
//	user.Locale.String = constanta.DEFAULT_LANGUAGE
//	user.AdditionalInfo.String = constanta.UserHasNotActivated
//	user.LastToken.Int64 = time.Now().Unix()
//	user.CreatedBy.Int64 = config.ApplicationConfiguration.GetClientCredentialsAuthUserID()
//	user.CreatedClient.String = config.ApplicationConfiguration.GetClientCredentialsClientID()
//	user.CreatedAt.Int64 = time.Now().Unix()
//	user.UpdatedBy.Int64 = config.ApplicationConfiguration.GetClientCredentialsAuthUserID()
//	user.UpdatedClient.String = config.ApplicationConfiguration.GetClientCredentialsClientID()
//	user.UpdatedAt.Int64 = time.Now().Unix()
//	return
//}
//
//// created at 06-20-2022
//// updated at 06-24-2022
//func UserGeneralToUserSimple(user model.UserGeneral) (userSimple res.UserSimple) {
//	userSimple.UserID = user.UserID.Int64
//	userSimple.Username = user.Username.String
//	userSimple.Email = user.Email.String
//	userSimple.ClientID = user.ClientID.String
//	userSimple.Status = user.Status.Int64
//	userSimple.Locale = user.Locale.String
//	userSimple.AdditionalInfo = user.AdditionalInfo.String
//	return
//}
//
//// created at 07-05-2022
//func UserGeneralToUserSimpleWithToken(user model.UserGeneral, token string) (userSimple res.UserSimpleWithToken) {
//	userSimple.UserID = user.UserID.Int64
//	userSimple.Username = user.Username.String
//	userSimple.Email = user.Email.String
//	userSimple.ClientID = user.ClientID.String
//	userSimple.Status = user.Status.Int64
//	userSimple.Locale = user.Locale.String
//	userSimple.AdditionalInfo = user.AdditionalInfo.String
//	userSimple.UserToken = token
//	return
//}
