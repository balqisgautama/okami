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
//	"strconv"
//)
//
//// created at 06-24-2022
//// updated at 07-04-2022
//func (input emailService) ResendEmail(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
//	var templateData converter.DataHTMLFile
//
//	params := util.GenerateQueryParam(request)
//	code := params[constanta.CODE_key]
//	emailTo := params[constanta.EMAILTO_KEY]
//	username := params[constanta.USERNAME_KEY]
//
//	templateData.Title = constanta.ActivationFailed
//
//	checkUser, outputCheck := dao.UserDAO.GetUserByEmailUsername(username, emailTo)
//	if checkUser.UserID.Int64 == 0 && outputCheck.Status.Code != "" || checkUser.Status.Int64 == constanta.UserDeleted {
//		templateData.BeforeButton = constanta.UserNotFound
//		templateData.Button = constanta.ResendEmail
//		templateData.ButtonUrl = ""
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//	if checkUser.Status.Int64 == constanta.UserActive {
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.Title = constanta.UserHasAlreadyActivated
//		templateData.BeforeButton = constanta.PleaseCompleteYourProfile
//		templateData.Button = constanta.BackToHome
//		templateData.ButtonUrl = ""
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//
//	validateData, outputValidate := dao.ActivationAccountDAO.GetActivationByID(checkUser.ActivationID.Int64)
//	if outputValidate.Status.Code != "" {
//		note := outputValidate.Status.Code + "<br>" + outputValidate.Status.Message + "<br>" + outputValidate.Status.Detail + "</b>"
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.AfterButton = note
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//
//	codeData, verifyStatus := util.ValidateJWT(code, strconv.Itoa(int(validateData.Code.Int64)))
//	if verifyStatus != "" {
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.AfterButton = verifyStatus
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//	if checkUser.Email.String != codeData.Email || checkUser.Username.String != codeData.Username || verifyStatus != "" {
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.BeforeButton = constanta.DataDoesNotMatch
//		templateData.Button = constanta.ResendEmail
//		templateData.ButtonUrl = validateData.LinkResend.String
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//
//	outputActivation := input.SendActivationEmail(checkUser.ActivationID.Int64, username, emailTo)
//	if outputActivation.Status.Code != "" {
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.Title = constanta.EmailResendFailed
//		templateData.Button = constanta.ResendEmail
//		templateData.ButtonUrl = validateData.LinkResend.String
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//
//	output.Status.Detail = constanta.ContentTypeHTML
//	templateData.Title = constanta.EmailResendSuccess
//	templateData.BeforeButton = constanta.MessageResendSuccess
//	output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//	return
//}
//
//// created at 06-21-2022
//func (input emailService) readBodyAndValidateResend(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqEmail) (output res.Payload)) (inputStruct req.ReqEmail, output res.Payload) {
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
//	return
//}
