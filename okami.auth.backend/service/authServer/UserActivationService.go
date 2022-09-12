package service

import (
	"encoding/json"
	"net/http"
	"okami.auth.backend/config"
	"okami.auth.backend/constanta"
	"okami.auth.backend/dao"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	"okami.auth.backend/service"
	"okami.auth.backend/util"
	jwt "okami.auth.backend/util/JWT"
	converter "okami.auth.backend/util/converter"
	"time"
)

type userActivationService struct {
	service.AbstractService
}

var UserActivationService = userActivationService{service.AbstractService{
	FileName: "service/authServer/UserActivationService.go",
}}

// ValidateUserActivation ------------------------------------------------------------------------------------------
// created at 08-23-2022
func (input userActivationService) ValidateUserActivation(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	//funcName := "ValidateUserActivation"
	var templateHTML converter.DataHTMLFile
	var note string
	var output_ res.Payload

	params := util.GenerateQueryParam(request)
	codeJWT := params[constanta.CODE_key]
	username := params[constanta.USERNAME_KEY]
	emailTo := params[constanta.EMAILTO_KEY]

	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestGET, constanta.APIEmailValidate, config.ApplicationConfiguration.GetClientCredentialsClientID()))

	userFounded, output_ := dao.UserDAO.GetUserByEmailUsername(emailTo, username)
	if output_.Status.Code != "" {
		note = output_.Status.Code + "<br>" + output_.Status.Message + "<br>" + output_.Status.Detail
		templateHTML.Title = constanta.ActivationFailed
		templateHTML.BeforeButton = constanta.NoteActivationFailedUserNotFound
		templateHTML.Button = constanta.BackToHome
		templateHTML.ButtonUrl = constanta.LinkOkami
		templateHTML.AfterButton = note
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	userActivationFounded, output_ := dao.UserActivationDao.GetActivationByUserID(userFounded.UserID.Int64)
	if output.Status.Code != "" {
		note = output_.Status.Code + "<br>" + output_.Status.Message + "<br>" + output_.Status.Detail
		templateHTML.Title = constanta.ActivationFailed
		templateHTML.BeforeButton = constanta.NoteActivationFailedUserNotFound
		templateHTML.Button = constanta.BackToHome
		templateHTML.ButtonUrl = constanta.LinkOkami
		templateHTML.AfterButton = note
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	} else if userActivationFounded.ActivationID.Int64 == 0 {
		templateHTML.Title = constanta.ActivationSuccess
		templateHTML.BeforeButton = constanta.NoteActivationSuccessUserActivated
		templateHTML.Button = constanta.LoginHere
		templateHTML.ButtonUrl = constanta.LinkOkami
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	_, err := jwt.ValidateJWTActivation(codeJWT, userActivationFounded.Code.String)
	if err != "" {
		templateHTML.Title = constanta.ActivationFailed
		templateHTML.BeforeButton = constanta.NoteActivationFailedUserNotFound
		templateHTML.Button = constanta.ResendEmail
		templateHTML.ButtonUrl = userActivationFounded.EmailLinkResend.String
		templateHTML.AfterButton = err
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	userActivationFounded.Status.Int64 = int64(constanta.UserActive)
	userActivationUpdated, output := dao.UserActivationDao.UpdateActivationByActivationID(userActivationFounded)
	if output.Status.Code != "" {
		templateHTML.Title = constanta.ActivationFailed
		templateHTML.BeforeButton = constanta.NoteActivationFailedUserNotFound
		templateHTML.Button = constanta.ResendEmail
		templateHTML.ButtonUrl = userActivationFounded.EmailLinkResend.String
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	dataOld, _ := json.Marshal(userActivationFounded)
	dataNew, _ := json.Marshal(userActivationUpdated)
	input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionUpdate, constanta.TableUserActivation,
		config.ApplicationConfiguration.GetClientCredentialsClientID(), string(dataOld), string(dataNew),
		constanta.ActionIDUpdate))

	userFounded.UpdatedClient.String = config.ApplicationConfiguration.GetClientCredentialsClientID()
	userFounded.UpdatedBy.Int64 = config.ApplicationConfiguration.GetClientCredentialsAuthUserID()
	userFounded.Status.Int64 = constanta.UserActive
	userUpdated, output := dao.UserDAO.UpdateUserByClientID(userFounded)
	if output.Status.Code != "" {
		templateHTML.Title = constanta.ActivationFailed
		templateHTML.BeforeButton = constanta.NoteActivationFailedUserNotFound
		templateHTML.Button = constanta.ResendEmail
		templateHTML.ButtonUrl = userActivationFounded.EmailLinkResend.String
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	dataOld, _ = json.Marshal(userFounded)
	dataNew, _ = json.Marshal(userUpdated)
	input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionUpdate, constanta.TableUsers,
		config.ApplicationConfiguration.GetClientCredentialsClientID(), string(dataOld), string(dataNew),
		constanta.ActionIDUpdate))

	templateHTML.Title = constanta.ActivationSuccess
	templateHTML.BeforeButton = constanta.NoteActivationSuccessUserActivated
	templateHTML.Button = constanta.LoginHere
	templateHTML.ButtonUrl = constanta.LinkOkami
	output.Status.Detail = constanta.ContentTypeHTML
	output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
	return
}

// ResendUserActivation ------------------------------------------------------------------------------------------
// created at 08-23-2022
func (input userActivationService) ResendUserActivation(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	//funcName := "ResendUserActivation"
	var templateHTML converter.DataHTMLFile
	var note string
	var output_ res.Payload

	params := util.GenerateQueryParam(request)
	codeJWT := params[constanta.CODE_key]
	username := params[constanta.USERNAME_KEY]
	emailTo := params[constanta.EMAILTO_KEY]

	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestGET, constanta.APIEmailResend, config.ApplicationConfiguration.GetClientCredentialsClientID()))

	userFounded, output_ := dao.UserDAO.GetUserByEmailUsername(emailTo, username)
	if output_.Status.Code != "" {
		note = output_.Status.Code + "<br>" + output_.Status.Message + "<br>" + output_.Status.Detail
		templateHTML.Title = constanta.EmailResendFailed
		templateHTML.BeforeButton = constanta.NoteActivationFailedUserNotFound
		templateHTML.Button = constanta.BackToHome
		templateHTML.ButtonUrl = constanta.LinkOkami
		templateHTML.AfterButton = note
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	userActivationFounded, output_ := dao.UserActivationDao.GetActivationByUserID(userFounded.UserID.Int64)
	if output.Status.Code != "" {
		note = output_.Status.Code + "<br>" + output_.Status.Message + "<br>" + output_.Status.Detail
		templateHTML.Title = constanta.EmailResendFailed
		templateHTML.BeforeButton = constanta.NoteActivationFailedUserNotFound
		templateHTML.Button = constanta.BackToHome
		templateHTML.ButtonUrl = constanta.LinkOkami
		templateHTML.AfterButton = note
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	} else if userActivationFounded.ActivationID.Int64 == 0 {
		templateHTML.Title = constanta.ActivationSuccess
		templateHTML.BeforeButton = constanta.NoteActivationSuccessUserActivated
		templateHTML.Button = constanta.LoginHere
		templateHTML.ButtonUrl = constanta.LinkOkami
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	_, err := jwt.ValidateJWTActivation(codeJWT, userActivationFounded.Code.String)
	if err != "" {
		templateHTML.Title = constanta.EmailResendFailed
		templateHTML.BeforeButton = constanta.NoteActivationFailedUserNotFound
		templateHTML.Button = constanta.ResendEmail
		templateHTML.ButtonUrl = userActivationFounded.EmailLinkResend.String
		templateHTML.AfterButton = err
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	if userActivationFounded.Counter.Int64 >= 3 && userActivationFounded.ExpiredAt.Time.Day() == time.Now().Day() {
		templateHTML.Title = constanta.EmailResendFailed
		templateHTML.BeforeButton = constanta.NoteActivationFailedResendCounter
		templateHTML.Button = constanta.BackToHome
		templateHTML.ButtonUrl = constanta.LinkOkami
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	if userActivationFounded.Counter.Int64 >= 3 && userActivationFounded.ExpiredAt.Time.Day() != time.Now().Day() {
		userActivationFounded.Counter.Int64 = 0
		userActivationConverted := converter.ToUpdatedActivationUserData(userActivationFounded, userFounded)
		userActivationUpdated, _ := dao.UserActivationDao.UpdateActivationByActivationID(userActivationConverted)
		dataOld, _ := json.Marshal(userActivationFounded)
		dataNew, _ := json.Marshal(userActivationUpdated)
		input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionUpdate, constanta.TableUserActivation,
			config.ApplicationConfiguration.GetClientCredentialsClientID(), string(dataOld), string(dataNew),
			constanta.ActionIDUpdate))
	}

	output = input.SendEmailActivationUser(userFounded, userActivationFounded)
	if output.Status.Code != "" {
		templateHTML.Title = constanta.EmailResendFailed
		templateHTML.BeforeButton = constanta.PleaseCheckYourConnection
		templateHTML.Button = constanta.ResendEmail
		templateHTML.ButtonUrl = userActivationFounded.EmailLinkResend.String
		templateHTML.AfterButton = err
		output.Status.Detail = constanta.ContentTypeHTML
		output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
		return
	}

	templateHTML.Title = constanta.EmailResendSuccess
	templateHTML.BeforeButton = constanta.PleaseCheckYourEmail
	output.Status.Detail = constanta.ContentTypeHTML
	output.Status.Message = converter.ParseHTMLFileToString(constanta.ActivationTemplatePath, templateHTML)
	return
}
