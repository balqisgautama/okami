package service

import (
	"encoding/json"
	"net/http"
	"okami.auth.backend/config"
	"okami.auth.backend/constanta"
	"okami.auth.backend/dao"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	errorModel "okami.auth.backend/model/error"
	"okami.auth.backend/util"
	jwt "okami.auth.backend/util/JWT"
	converter "okami.auth.backend/util/converter"
	"time"
)

type AbstractService struct {
	FileName string
	Audit    bool
}

func (input AbstractService) ReadBody(request *http.Request, contextModel *model.ContextModel) (string, errorModel.ErrorModel) {
	var stringBody string
	var errorS error

	funcName := "ReadBody"

	if request.Method != "GET" {
		stringBody, contextModel.LoggerModel.ByteIn, errorS = util.ReadBody(request)
		if errorS != nil {
			return "", errorModel.GenerateInvalidRequestError(input.FileName, funcName, errorS)
		}
	}

	return stringBody, errorModel.GenerateNonErrorModel()
}

// TokenChecker ------------------------------------------------------------------------------------------
// created at 08-19-2022
// updateda at 08-27-2022
func (input AbstractService) TokenChecker(token, clientID, funcName string) (data *jwt.PayloadJWTResource, output res.Payload) {
	data, response := jwt.ValidateJWTResource(token, config.ApplicationConfiguration.GetJWTToken().Internal)
	if response != "" {
		output.Status.Code = constanta.CodeVerifyFailed
		output.Status.Message = response
		output.Status.Detail = funcName
		return
	}
	if data.ClientID != clientID {
		output.Status.Code = constanta.CodeAuthorizationFailed
		output.Status.Message = constanta.InvalidToken
		output.Status.Detail = funcName
		return
	}
	return
}

// LogActivity ------------------------------------------------------------------------------------------
// created at 08-22-2022
func (input AbstractService) LogActivity(data model.LogActivity) {
	dao.LogActivitiesDAO.InsertLogActivity(data)
}

// LogAuditSystem ------------------------------------------------------------------------------------------
// created at 08-22-2022
func (input AbstractService) LogAuditSystem(data model.LogAuditSystem) {
	dao.LogAuditSystemDAO.InsertLogAuditSystem(data)
}

// SendEmailActivationUser ------------------------------------------------------------------------------------------
// created at 08-22-2022
func (input AbstractService) SendEmailActivationUser(user model.UserGeneral, activationUpdated model.ActivationUser) (output res.Payload) {
	funcName := "SendEmailActivationUser"
	var activationData model.ActivationUser

	if activationUpdated.ActivationID.Int64 == 0 {
		data := converter.UserModelToActivationUserModel(user)
		activationData, output = dao.UserActivationDao.InsertActivation(data)
		if output.Status.Code != "" {
			return
		}
		input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestPOST, constanta.APIUser, config.ApplicationConfiguration.GetClientCredentialsClientID()))
	} else {
		userActivationConverted := converter.ToUpdatedActivationUserData(activationUpdated, user)
		activationData, output = dao.UserActivationDao.UpdateActivationByActivationID(userActivationConverted)
		if output.Status.Code != "" {
			return
		}
	}

	dataNew, _ := json.Marshal(activationData)
	input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionInsert, constanta.TableUserActivation,
		config.ApplicationConfiguration.GetClientCredentialsClientID(), "", string(dataNew),
		constanta.ActionIDInsert))

	send := util.SendEmailGeneral([]string{user.Email.String}, constanta.EmailVerificationSubject,
		activationData.EmailLinkValidate.String, user.Username.String)
	if !send {
		output.Status.Code = constanta.CodeSendEmailFailed
		output.Status.Message = constanta.EmailResendFailed
		output.Status.Detail = funcName
		return
	}
	return
}
