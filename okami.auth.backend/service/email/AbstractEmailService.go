package service

import (
	"okami.auth.backend/constanta"
	"okami.auth.backend/dao"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/service"
	"okami.auth.backend/util"
)

type emailService struct {
	FileName string
	service.AbstractService
}

var EmailService = emailService{FileName: "EmailService.go"}

// created at 06-21-2022
// updated at 06-23-2022
func (input emailService) SendActivationEmail(activationID int64, username string, emailTo string) (output res.Payload) {
	activationData, outputActivation := dao.ActivationAccountDAO.GetActivationByID(activationID)
	if activationData.ActivationID.Int64 == 0 && outputActivation.Status.Code != "" {
		output.Status.Code = constanta.CodeSendEmailFailed
		output.Status.Message = outputActivation.Status.Message
		return
	}

	send := util.SendEmailGeneral([]string{emailTo}, constanta.EmailVerificationSubject, activationData.LinkValidate.String, username)
	if !send {
		output.Status.Code = constanta.CodeSendEmailFailed
		output.Status.Message = constanta.EmailResendFailed
		return
	}
	return
}
