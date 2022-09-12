package service

//
//import (
//	"encoding/json"
//	"net/http"
//	"okami.auth.backend/config"
//	"okami.auth.backend/constanta"
//	"okami.auth.backend/dao"
//	req "okami.auth.backend/dto/in"
//	res "okami.auth.backend/dto/out"
//	"okami.auth.backend/model"
//	"okami.auth.backend/util"
//	converter "okami.auth.backend/util/converter"
//)
//
//// created at 07-05-2022
//func (input loginService) GeneralLogin(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
//	result, output := input.readBodyAndValidateLogin(request, context, req.ValidateReqLoginGeneral)
//	if output.Status.Code != "" {
//		return
//	}
//
//	userData, output := dao.UserDAO.GetUserByUsername(result.Username)
//	if output.Status.Code != "" {
//		output.Status.AdditionalInfo = []string{constanta.UserNotFound}
//		return
//	}
//
//	output = input.checkingPassword(result.Password, userData.Password.String)
//	if output.Status.Code != "" {
//		return
//	}
//
//	if userData.Status.Int64 != constanta.UserActive {
//		output.Status.Code = constanta.CodeLoginFailed
//		output.Status.Message = constanta.UserHasNotActivated
//		return
//	}
//
//	codeJWT, outputJWT := util.GenerateJWT(config.ApplicationConfiguration.GetJWTToken().JWT, "", config.ApplicationConfiguration.GetServerVersion(), userData.Username.String, userData.Email.String, userData.UserID.Int64)
//	if outputJWT != "" {
//		output.Status.Code = constanta.CodeAuthorizationFailed
//		output.Status.Message = outputJWT
//		return
//	}
//
//	output.Status.Code = constanta.PayloadStatusCode
//	output.Status.Message = constanta.LoginSuccess
//	output.Data.Content = converter.UserGeneralToUserSimpleWithToken(userData, codeJWT)
//	return
//}
//
//// created at 07-05-2022
//func (input loginService) readBodyAndValidateLogin(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqLoginGeneral) (output res.Payload)) (inputStruct req.ReqLoginGeneral, output res.Payload) {
//	var stringBody string
//
//	stringBody, err := input.ReadBody(request, contextModel)
//	if err.Error != nil {
//		output.Status.Code = constanta.CodeValidationFailed
//		output.Status.Message = err.Error.Error()
//		return
//	}
//
//	if stringBody != "" {
//		errorS := json.Unmarshal([]byte(stringBody), &inputStruct)
//		if errorS != nil {
//			output.Status.Code = constanta.CodeValidationFailed
//			output.Status.Message = errorS.Error()
//			return
//		}
//	}
//	output = validation(&inputStruct)
//	if output.Status.Code != "" {
//		return
//	}
//
//	regexUsername, err1, err2 := util.IsUsernameStandardValid(inputStruct.Username)
//	if !regexUsername {
//		output.Status.Code = constanta.CodeValidationFailed
//		output.Status.Message = constanta.RegexValidationFailed + " (" + err1 + " " + err2 + ")"
//		output.Status.Detail = constanta.FormatUsernameIsWrong
//		return
//	}
//
//	regexPassword, err1, err2 := util.IsPasswordStandardValid(inputStruct.Password)
//	if !regexPassword {
//		output.Status.Code = constanta.CodeValidationFailed
//		output.Status.Message = constanta.RegexValidationFailed + " (" + err1 + " " + err2 + ")"
//		output.Status.Detail = constanta.FormatPasswordIsWrong
//		return
//	}
//
//	return
//}
