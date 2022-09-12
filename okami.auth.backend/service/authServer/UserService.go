package service

import (
	"encoding/json"
	"net/http"
	"okami.auth.backend/config"
	"okami.auth.backend/constanta"
	"okami.auth.backend/dao"
	req "okami.auth.backend/dto/in"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	"okami.auth.backend/service"
	"okami.auth.backend/util"
	converter "okami.auth.backend/util/converter"
	"time"
)

type userService struct {
	service.AbstractService
}

var UserService = userService{service.AbstractService{
	FileName: "service/authServer/UserService.go",
}}

// created at 08-19-2022
func (input userService) resourceIsExist(funcName string, result req.ReqUserGeneral) (output res.Payload) {
	resourceFounded, output := dao.ResourceDAO.GetResourceByClientID(result.ResourceClientID)
	if output.Status.Code != "" {
		return
	}

	if result.ResourceClientID == config.ApplicationConfiguration.GetClientCredentialsClientID() {
		output.Status.Code = constanta.CodeRequestFailed
		output.Status.Message = constanta.AccessForbidden
		output.Status.Detail = funcName
		return
	}

	// jika clientID tidak ada di database atau userID di request body tidak sama dengan data di database
	if resourceFounded.ResourceID.Int64 != result.ResourceID || resourceFounded.ResourceID.Int64 == 0 {
		output.Status.Code = constanta.CodeRequestFailed
		output.Status.Message = constanta.DataDoesNotMatch
		output.Status.Detail = funcName
		return
	}
	return
}

// CreateUser ------------------------------------------------------------------------------------------
// created at 08-19-2022
// updated at 08-27-2022
func (input userService) CreateUser(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "CreateUser"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)
	tokenData, output := input.TokenChecker(tokenHeader, config.ApplicationConfiguration.GetClientCredentialsClientID(), funcName)
	if output.Status.Code != "" {
		return
	}

	result, output := input.readBodyAndValidateGeneral(request, context, req.ValidateReqUserGeneral)
	if output.Status.Code != "" {
		return
	}

	output = input.resourceIsExist(funcName, result)
	if output.Status.Code != "" {
		return
	}

	converted := converter.ReqUserCreateToUserForInsertDB(result)
	userInserted, output := dao.UserDAO.InsertUser(converted)
	if output.Status.Code != "" {
		return
	}

	output = input.SendEmailActivationUser(userInserted, model.ActivationUser{})
	if output.Status.Code != "" {
		return
	}

	dataNew, _ := json.Marshal(converted)
	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestPOST, constanta.APIUser, tokenData.ClientID))
	input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionInsert, constanta.TableUsers,
		tokenData.ClientID, "", string(dataNew),
		constanta.ActionIDInsert))

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.InsertDataSuccess
	output.Status.Detail = funcName
	output.Data = converter.DBDataToUserGeneral(userInserted)
	return
}

// created at 08-19-2022
func (input userService) readBodyAndValidateGeneral(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqUserGeneral) (output res.Payload)) (inputStruct req.ReqUserGeneral, output res.Payload) {
	var stringBody string

	stringBody, err := input.ReadBody(request, contextModel)
	if err.Error != nil {
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Message = err.Error.Error()
		return
	}

	if stringBody != "" {
		errorS := json.Unmarshal([]byte(stringBody), &inputStruct)
		if errorS != nil {
			output.Status.Code = constanta.CodeValidationFailed
			output.Status.Message = errorS.Error()
			return
		}
	}
	output = validation(&inputStruct)
	if output.Status.Code != "" {
		return
	}

	if inputStruct.UserClientID != "" && inputStruct.Username == "" &&
		inputStruct.Email == "" && inputStruct.Password == "" {
		return
	}

	regexUsername, err1, err2 := util.IsUsernameStandardValid(inputStruct.Username)
	if !regexUsername {
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Message = constanta.RegexValidationFailed + " (" + err1 + " " + err2 + ")"
		output.Status.Detail = constanta.FormatUsernameIsWrong
		return
	}

	regexPassword, err1, err2 := util.IsPasswordStandardValid(inputStruct.Password)
	if !regexPassword {
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Message = constanta.RegexValidationFailed + " (" + err1 + " " + err2 + ")"
		output.Status.Detail = constanta.FormatPasswordIsWrong
		return
	}
	return
}

// ReadUser ------------------------------------------------------------------------------------------
// created at 08-19-2022
// updated at 08-27-2022
func (input userService) ReadUser(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "ReadUser"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)
	tokenData, output := input.TokenChecker(tokenHeader, config.ApplicationConfiguration.GetClientCredentialsClientID(), funcName)
	if output.Status.Code != "" {
		return
	}

	params := util.GenerateQueryParam(request)
	paramClientID := params[constanta.CLIENTID_KEY]

	// get user by clientID
	if paramClientID != "" {
		resource, output_ := dao.UserDAO.GetUserByClientID(paramClientID)
		if output_.Status.Code != "" {
			output.Status.Code = output_.Status.Code
			output.Status.Message = output_.Status.Message
			output.Status.Detail = output_.Status.Detail
			output.Status.AdditionalInfo = output_.Status.AdditionalInfo
			return
		}

		input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestGET, constanta.APIUser, tokenData.ClientID))

		output.Status.Code = constanta.PayloadStatusCode
		output.Status.Message = constanta.DataFounded
		output.Status.Detail = funcName
		output.Data = converter.DBDataToUserGeneral(resource)
		return
	}

	// get all user
	resources, output := dao.UserDAO.GetUsers()
	if output.Status.Code != "" {
		return
	}

	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestGET, constanta.APIUser, tokenData.ClientID))

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.DataFounded
	output.Status.Detail = funcName
	output.Data = converter.DBDataToUserGeneralArray(resources)
	return
}

// UpdateUser ------------------------------------------------------------------------------------------
// created at 08-19-2022
// updated at 08-27-2022
func (input userService) UpdateUser(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "UpdateUser"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)
	tokenData, output := input.TokenChecker(tokenHeader, config.ApplicationConfiguration.GetClientCredentialsClientID(), funcName)
	if output.Status.Code != "" {
		return
	}

	result, output := input.readBodyAndValidateGeneral(request, context, req.ValidateReqUserGeneral)
	if output.Status.Code != "" {
		return
	}

	output = input.resourceIsExist(funcName, result)
	if output.Status.Code != "" {
		return
	}

	if result.UserClientID == config.ApplicationConfiguration.GetClientCredentialsClientID() {
		output.Status.Code = constanta.CodeRequestFailed
		output.Status.Message = constanta.AccessForbidden
		output.Status.Detail = funcName
		return
	}

	userOld, output := dao.UserDAO.GetUserByClientID(result.UserClientID)
	if output.Status.Code != "" {
		return
	}

	password, err := util.HashPassword(result.Password)
	if err != nil {
		output.Status.Code = constanta.CodeRequestFailed
		output.Status.Message = constanta.UpdateDataFailed
		output.Status.Detail = funcName
		return
	}
	result.Password = password

	converted := converter.ReqUserGeneralToUserGeneralModel(result)
	converted.Status.Int64 = userOld.Status.Int64
	userUpdated, output := dao.UserDAO.UpdateUserByClientID(converted)
	if output.Status.Code != "" {
		return
	}

	dataOld, _ := json.Marshal(userOld)
	dataNew, _ := json.Marshal(userUpdated)
	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestPUT, constanta.APIUser, tokenData.ClientID))
	input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionUpdate, constanta.TableUsers,
		tokenData.ClientID, string(dataOld), string(dataNew),
		constanta.ActionIDUpdate))

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.UpdateDataSuccess
	output.Status.Detail = funcName
	output.Data = converter.DBDataToUserGeneral(userUpdated)
	return
}

// DeleteUser ------------------------------------------------------------------------------------------
// created at 08-19-2022
// updated at 08-27-2022
func (input userService) DeleteUser(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "DeleteUser"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)
	tokenData, output := input.TokenChecker(tokenHeader, config.ApplicationConfiguration.GetClientCredentialsClientID(), funcName)
	if output.Status.Code != "" {
		return
	}

	result, output := input.readBodyAndValidateGeneral(request, context, req.ValidateReqUserGeneral)
	if output.Status.Code != "" {
		return
	}

	output = input.resourceIsExist(funcName, result)
	if output.Status.Code != "" {
		return
	}

	if result.UserClientID == config.ApplicationConfiguration.GetClientCredentialsClientID() {
		output.Status.Code = constanta.CodeRequestFailed
		output.Status.Message = constanta.AccessForbidden
		output.Status.Detail = funcName
		return
	}

	userOld, output := dao.UserDAO.GetUserByClientID(result.UserClientID)
	if output.Status.Code != "" {
		return
	}

	user, output := dao.UserDAO.SoftDeleteUser(result)
	if output.Status.Code != "" {
		return
	}

	users, output := dao.UserDAO.GetUsers()

	dataOld, _ := json.Marshal(userOld)
	dataNew, _ := json.Marshal(user)
	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestDELETE, constanta.APIUser, tokenData.ClientID))
	input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionDelete, constanta.TableUsers,
		tokenData.ClientID, string(dataOld), string(dataNew),
		constanta.ActionIDDelete))

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.DeleteDataSuccess
	output.Status.Detail = funcName
	output.Data = map[string]interface{}{
		constanta.FieldDataDeleted: converter.DBDataToUserGeneral(user),
		constanta.FieldDataNewest:  converter.DBDataToUserGeneralArray(users),
	}
	return
}
