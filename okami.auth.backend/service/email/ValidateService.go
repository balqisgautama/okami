package service

//
//import (
//	"net/http"
//	"okami.auth.backend/constanta"
//	"okami.auth.backend/dao"
//	res "okami.auth.backend/dto/out"
//	"okami.auth.backend/model"
//	"okami.auth.backend/util"
//	converter "okami.auth.backend/util/converter"
//	"strconv"
//)
//
//// created at 06-22-2022
//// updated at 06-23-2022
//func (input emailService) ValidateEmail(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
//	var templateData converter.DataHTMLFile
//
//	params := util.GenerateQueryParam(request)
//	code := params[constanta.CODE_key]
//	emailTo := params[constanta.EMAILTO_KEY]
//	username := params[constanta.USERNAME_KEY]
//
//	templateData.Title = constanta.ActivationFailed
//
//	checkUser, outputUser := dao.UserDAO.GetUserByEmailUsername(username, emailTo)
//	if outputUser.Status.Code != "" {
//		note := outputUser.Status.Code + "<br>" + outputUser.Status.Message + "<br>" + outputUser.Status.Detail
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.ButtonUrl = constanta.UserNotFound
//		templateData.AfterButton = note
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//	validateData, outputValidate := dao.ActivationAccountDAO.GetActivationByID(checkUser.ActivationID.Int64)
//	if outputValidate.Status.Code != "" {
//		note := outputValidate.Status.Code + "<br>" + outputValidate.Status.Message + "<br>" + outputValidate.Status.Detail + "</b>"
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.BeforeButton = constanta.UserNotFound
//		templateData.AfterButton = note
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//	codeData, verifyStatus := util.ValidateJWT(code, strconv.Itoa(int(validateData.Code.Int64)))
//	if verifyStatus != "" {
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.BeforeButton = constanta.InvalidToken
//		templateData.Button = constanta.ResendEmail
//		templateData.ButtonUrl = validateData.LinkResend.String
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
//	_, outputStatusActivation := dao.ActivationAccountDAO.UpdateStatusByID(checkUser.ActivationID.Int64)
//	if outputStatusActivation.Status.Code != "" {
//		templateData.Button = constanta.ResendEmail
//		note := outputStatusActivation.Status.Code + "<br>" + outputStatusActivation.Status.Message + "<br>" + outputStatusActivation.Status.Detail + "</b>"
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.AfterButton = note
//		templateData.ButtonUrl = validateData.LinkResend.String
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//	_, outputStatusUser := dao.UserDAO.UpdateStatusByUserID(checkUser.UserID.Int64)
//	if outputStatusUser.Status.Code != "" {
//		templateData.Button = constanta.ResendEmail
//		note := outputStatusUser.Status.Code + "<br>" + outputStatusUser.Status.Message + "<br>" + outputStatusUser.Status.Detail + "</b>"
//		output.Status.Detail = constanta.ContentTypeHTML
//		templateData.AfterButton = note
//		templateData.ButtonUrl = validateData.LinkResend.String
//		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//		return
//	}
//
//	output.Status.Detail = constanta.ContentTypeHTML
//	templateData.Title = constanta.ActivationSuccess
//	templateData.BeforeButton = constanta.MessageRegistrationSuccess
//	templateData.Button = constanta.BackToHome
//	templateData.ButtonUrl = "https://www.google.com/"
//	output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateData)
//	return
//}
