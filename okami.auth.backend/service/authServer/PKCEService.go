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
	jwt "okami.auth.backend/util/JWT"
	converter "okami.auth.backend/util/converter"
	"time"
)

type pkceService struct {
	service.AbstractService
}

var PKCEService = pkceService{service.AbstractService{
	FileName: "service/authServer/PKCEService.go",
}}

// PKCEStep1 ------------------------------------------------------------------------------------------
// created at 08-26-2022
// updated at 08-30-2022
func (input pkceService) PKCEStep1(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "PKCEStep1"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)
	tokenData, output := input.TokenChecker(tokenHeader, config.ApplicationConfiguration.GetClientCredentialsClientID(), funcName)
	if output.Status.Code != "" {
		return
	}

	result, output := input.readBodyAndValidatePKCEStep1(request, context, req.ValidateReqPKCEStep1)
	if output.Status.Code != "" {
		return
	}

	token, err := jwt.GenerateJWTPCKEStep1(result.CodeChallenger)
	if err != "" {
		output.Status.Code = constanta.CodeAuthorizationFailed
		output.Status.Message = constanta.GenerateFailed
		return
	}

	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestPOST, constanta.APIPKCEStep1, tokenData.ClientID))

	_, output = dao.LogPKCEDAO.InsertLogPKCE(result.CodeChallenger)
	if output.Status.Code != "" {
		return
	}

	MyMap := make(map[string]string)
	MyMap[constanta.TokenHeaderNameConstanta] = token
	header = MyMap

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.AuthSuccess
	return
}

// created at 08-26-2022
func (input pkceService) readBodyAndValidatePKCEStep1(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqPKCEStep1) (output res.Payload)) (inputStruct req.ReqPKCEStep1, output res.Payload) {
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

	return
}

// PKCEStep2 ------------------------------------------------------------------------------------------
// created at 08-30-2022
func (input pkceService) PKCEStep2(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	//funcName := "PKCEStep2"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)

	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestPOST, constanta.APIPKCEStep2, config.ApplicationConfiguration.GetClientCredentialsClientID()))

	result, output := input.readBodyAndValidatePKCEStep2(request, context, req.ValidateReqPKCEStep2)
	if output.Status.Code != "" {
		return
	}

	passwordHash, _ := util.HashPassword(result.Password)
	userFounded, output := dao.UserDAO.GetUserActiveByUsernamePassword(result.Username, passwordHash)
	if output.Status.Code != "" {
		return
	}

	//fmt.Println(tokenHeader)
	//fmt.Println(userFounded)
	tokenData, err := jwt.ValidateJWTPKCEStep1(tokenHeader)
	if err != "" {
		output.Status.Code = constanta.CodeRequestFailed
		output.Status.Message = constanta.InvalidToken
		return
	}

	dataFounded, output := dao.LogPKCEDAO.GetLogPKCEByCodeChallenger(tokenData.CodeChallenge)
	if output.Status.Code != "" {
		return
	}

	dataFounded.Step2.Time = time.Now()
	dataFounded.SecretCode.String = util.GetUUID()
	dataFounded.UserClientID.String = userFounded.ClientID.String
	_, output = dao.LogPKCEDAO.UpdateLogPKCEByCodeChallenger(dataFounded)

	MyMap := make(map[string]string)
	MyMap[constanta.HeaderKeySecretCode] = dataFounded.SecretCode.String
	header = MyMap

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.AuthSuccess
	return
}

// created at 08-30-2022
func (input pkceService) readBodyAndValidatePKCEStep2(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqPKCEStep2) (output res.Payload)) (inputStruct req.ReqPKCEStep2, output res.Payload) {
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

// PKCEStep3 ------------------------------------------------------------------------------------------
// created at 08-31-2022
func (input pkceService) PKCEStep3(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	//funcName := "PKCEStep3"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)

	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestPOST, constanta.APIPKCEStep3, config.ApplicationConfiguration.GetClientCredentialsClientID()))

	result, output := input.readBodyAndValidatePKCEStep3(request, context, req.ValidateReqPKCEStep3)
	if output.Status.Code != "" {
		return
	}

	dataFounded, output := dao.LogPKCEDAO.GetLogPKCEByCodeVerifier(tokenHeader)
	if output.Status.Code != "" {
		return
	}

	dataFounded.Step3.Time = time.Now()
	_, output = dao.LogPKCEDAO.UpdateLogPKCEByCodeChallenger(dataFounded)
	if output.Status.Code != "" {
		return
	}

	matched := util.SHA256(result.CodeVerifier)
	if matched != dataFounded.CodeChallenger.String {
		output.Status.Code = constanta.CodeRequestFailed
		output.Status.Message = constanta.DataDoesNotMatch
		return
	}

	userFounded, output := dao.UserDAO.GetUserByClientID(dataFounded.UserClientID.String)
	token, err := jwt.GenerateJWTUser(userFounded.ClientID.String, userFounded.UserID.Int64)
	if err != "" {
		output.Status.Code = constanta.CodeRequestFailed
		output.Status.Message = constanta.DataNotFound
		return
	}

	userFounded.UpdatedBy.Int64 = config.ApplicationConfiguration.GetClientCredentialsAuthUserID()
	userFounded.UpdatedClient.String = config.ApplicationConfiguration.GetClientCredentialsClientID()
	userFounded.LastToken.Time = time.Now()
	_, output = dao.UserDAO.UpdateUserByClientID(userFounded)

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.AuthSuccess
	output.Data = map[string]string{
		"user_token": token,
	}
	return
}

// created at 08-31-2022
func (input pkceService) readBodyAndValidatePKCEStep3(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqPKCEStep3) (output res.Payload)) (inputStruct req.ReqPKCEStep3, output res.Payload) {
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

	return
}
