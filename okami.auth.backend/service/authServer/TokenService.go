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

type tokenService struct {
	service.AbstractService
}

var TokenService = tokenService{service.AbstractService{
	FileName: "service/authServer/TokenService.go",
}}

// GenerateTokenResource ------------------------------------------------------------------------------------------
// created at 08-18-2022
// updated at 08-22-2022
func (input tokenService) GenerateTokenResource(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "GenerateTokenResource"
	params := util.GenerateQueryParam(request)
	paramClientID := params[constanta.CLIENTID_KEY]
	if paramClientID == "" {
		output.Status.Code = constanta.CodeGenerateFailed
		output.Status.Message = constanta.GenerateFailed
		return
	}

	resource, output := dao.ResourceDAO.GetResourceByClientID(paramClientID)
	if output.Status.Code != "" {
		return
	}
	token, err := jwt.GenerateJWTResource(
		config.ApplicationConfiguration.GetJWTToken().Internal,
		resource.Surname.String, resource.ClientID.String, resource.ResourceID.Int64)
	if err != "" {
		output.Status.Code = constanta.CodeGenerateFailed
		output.Status.Message = err
		output.Status.Detail = funcName
		return
	}

	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestGET, constanta.APITokenResource, resource.ClientID.String))

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.GenerateSuccess
	output.Status.Detail = funcName
	output.Data = map[string]string{
		"resource_token": token,
	}
	return
}

// VerifyTokenResource ------------------------------------------------------------------------------------------
// created at 08-18-2022
// updated at 08-22-2022
func (input tokenService) VerifyTokenResource(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "VerifyTokenResource"
	result, output := input.readBodyAndValidateVerifyTokenResource(request, context, req.ValidateReqTokenResource)
	if output.Status.Code != "" {
		return
	}

	tokenData, response := jwt.ValidateJWTResource(result.Token, config.ApplicationConfiguration.GetJWTToken().Internal)
	if response != "" {
		output.Status.Code = constanta.CodeVerifyFailed
		output.Status.Message = response
		output.Status.Detail = funcName
		return
	}
	if tokenData.ClientID != result.ClientID {
		output.Status.Code = constanta.CodeVerifyFailed
		output.Status.Message = constanta.InvalidToken
		output.Status.Detail = funcName
		return
	}

	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestPOST, constanta.APITokenResource, result.ClientID))

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.VerifySuccess
	return
}

// created at 08-18-2022
func (input tokenService) readBodyAndValidateVerifyTokenResource(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqTokenResource) (output res.Payload)) (inputStruct req.ReqTokenResource, output res.Payload) {
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
