package service

//
//import (
//	"encoding/json"
//	"net/http"
//	"okami.auth.backend/constanta"
//	"okami.auth.backend/dao"
//	req "okami.auth.backend/dto/in"
//	res "okami.auth.backend/dto/out"
//	"okami.auth.backend/model"
//	"okami.auth.backend/util"
//	converter "okami.auth.backend/util/converter"
//)
//
//type registrationService struct {
//	FileName string
//	AbstractService
//}
//
//var RegistrationService = registrationService{FileName: "RegistrationService.go"}
//
//// created at 06-16-2022
//// updated at 06-21-2022
//func (input registrationService) Registration(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
//	result, output := input.readBodyAndValidateRegist(request, context, req.ValidateReqRegist)
//	if output.Status.Code != "" {
//		return
//	}
//
//	checkDuplicate, duplicateOutput := dao.UserDAO.GetUsersByEmailUsername(result.Username, result.Email)
//	if duplicateOutput.Status.Code != "" {
//		output.Status.Code = duplicateOutput.Status.Code
//		output.Status.Message = duplicateOutput.Status.Message
//		output.Status.Detail = duplicateOutput.Status.Detail
//		return
//	}
//	if len(checkDuplicate) > 0 {
//		output.Status.Code = constanta.CodeRegistrationFailed
//		output.Status.Message = constanta.UsernameOrEmailAlreadyExist
//		return
//	}
//
//	activationData, linkOutput := converter.GenerateActivation(result.Email, result.Username)
//	if linkOutput.Status.Code != "" {
//		output.Status.Code = linkOutput.Status.Code
//		output.Status.Message = linkOutput.Status.Message
//		output.Status.Detail = linkOutput.Status.Detail
//		return
//	}
//
//	activationID, activationOutput := dao.ActivationAccountDAO.InsertData(activationData)
//	if activationID == 0 || activationOutput.Status.Code != "" {
//		output.Status.Code = activationOutput.Status.Code
//		output.Status.Message = activationOutput.Status.Message
//		output.Status.Detail = activationOutput.Status.Detail
//		return
//	}
//
//	output = SendActivationEmail(activationID, result.Username, result.Email)
//	if output.Status.Code != "" {
//		return
//	}
//
//	converted := converter.RegistToUserGeneral(result, activationID)
//	inserted, insertOutput := dao.UserDAO.InsertUser(converted)
//	if insertOutput.Status.Code != "" {
//		output.Status.Code = insertOutput.Status.Code
//		output.Status.Message = insertOutput.Status.Message
//		output.Status.Detail = insertOutput.Status.Detail
//		return
//	}
//
//	output.Status.Code = constanta.PayloadStatusCode
//	output.Status.Message = constanta.RegistrationSuccess + " " + constanta.PleaseCheckYourEmail
//	userData, _ := dao.UserDAO.GetUserByUserID(inserted)
//	output.Data.Content = converter.UserGeneralToUserSimple(userData)
//
//	return
//}
//
//// created at 06-16-2022
//// updated at 06-20-2022
//func (input registrationService) readBodyAndValidateRegist(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqRegist) (output res.Payload)) (inputStruct req.ReqRegist, output res.Payload) {
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
