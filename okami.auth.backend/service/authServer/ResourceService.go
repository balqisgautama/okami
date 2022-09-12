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
	"strings"
	"time"
)

type resourceService struct {
	service.AbstractService
}

var ResourceService = resourceService{service.AbstractService{
	FileName: "service/authServer/ResourceService.go",
}}

// ---------------------------------- General --------------------------------------------------------
// created at 08-19-2022
func (input resourceService) resourceIsExist(accessTo string, funcName string) (output res.Payload) {
	accesses := strings.Split(accessTo, " ")
	resources, _ := dao.ResourceDAO.GetResources()
	countResources := len(resources)
	countNotFound := 0
	for _, access := range accesses {
		if access == constanta.ForbiddenResourceAuth {
			output.Status.Code = constanta.CodeRequestFailed
			output.Status.Message = constanta.AccessForbidden
			output.Status.Detail = funcName
			output.Status.AdditionalInfo = []string{access}
			return
		}

		for _, resource := range resources {
			if access != resource.Nickname.String {
				countNotFound++
			}
		}

		if countNotFound == countResources {
			output.Status.Code = constanta.CodeRequestFailed
			output.Status.Message = constanta.DataNotFound
			output.Status.Detail = funcName
			output.Status.AdditionalInfo = []string{access}
			return
		}
		countNotFound = 0
	}
	return
}

// CreateResource ------------------------------------------------------------------------------------------
// created at 08-16-2022
// updated at 08-27-2022
func (input resourceService) CreateResource(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "CreateResource"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)
	tokenData, output := input.TokenChecker(tokenHeader, config.ApplicationConfiguration.GetClientCredentialsClientID(), funcName)
	if output.Status.Code != "" {
		return
	}

	result, output := input.readBodyAndValidateCreate(request, context, req.ValidateReqResourceCreate)
	if output.Status.Code != "" {
		return
	}

	output = input.resourceIsExist(result.AccessTo, funcName)
	if output.Status.Code != "" {
		return
	}

	converted := converter.ReqResourceCreateToResourceForInsertDB(result)
	resourceInserted, output := dao.ResourceDAO.InsertResource(converted)
	if output.Status.Code != "" {
		return
	}

	dataNew, _ := json.Marshal(converted)
	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestPOST, constanta.APIResource, tokenData.ClientID))
	input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionInsert, constanta.TableResources,
		tokenData.ClientID, "", string(dataNew),
		constanta.ActionIDInsert))
	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.InsertDataSuccess
	output.Status.Detail = funcName
	output.Data = converter.DBDataToResouceGeneral(resourceInserted)
	return
}

// created at 08-16-2022
func (input resourceService) readBodyAndValidateCreate(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqResourceCreate) (output res.Payload)) (inputStruct req.ReqResourceCreate, output res.Payload) {
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

// ReadResource ------------------------------------------------------------------------------------------
// created at 08-18-2022
// updated at 08-27-2022
func (input resourceService) ReadResource(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "ReadResource"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)
	tokenData, output := input.TokenChecker(tokenHeader, config.ApplicationConfiguration.GetClientCredentialsClientID(), funcName)
	if output.Status.Code != "" {
		return
	}

	params := util.GenerateQueryParam(request)
	paramClientID := params[constanta.CLIENTID_KEY]

	// get resource by clientID
	if paramClientID != "" {
		resource, output_ := dao.ResourceDAO.GetResourceByClientID(paramClientID)
		if output_.Status.Code != "" {
			output.Status.Code = output_.Status.Code
			output.Status.Message = output_.Status.Message
			output.Status.Detail = output_.Status.Detail
			output.Status.AdditionalInfo = output_.Status.AdditionalInfo
			return
		}

		input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestGET, constanta.APIResource, tokenData.ClientID))

		output.Status.Code = constanta.PayloadStatusCode
		output.Status.Message = constanta.DataFounded
		output.Status.Detail = funcName
		output.Data = converter.DBDataToResouceGeneral(resource)
		return
	}

	// get all resource
	resources, output := dao.ResourceDAO.GetResources()
	if output.Status.Code != "" {
		return
	}

	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestGET, constanta.APIResource, tokenData.ClientID))

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.DataFounded
	output.Status.Detail = funcName
	output.Data = converter.DBDataToResouceGeneralArray(resources)
	return
}

// UpdateResource ------------------------------------------------------------------------------------------
// created at 08-18-2022
// updated at 08-27-2022
func (input resourceService) UpdateResource(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "UpdateResource"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)
	tokenData, output := input.TokenChecker(tokenHeader, config.ApplicationConfiguration.GetClientCredentialsClientID(), funcName)
	if output.Status.Code != "" {
		return
	}

	result, output := input.readBodyAndValidateUpdate(request, context, req.ValidateReqResourceUpdate)
	if output.Status.Code != "" {
		return
	}

	output = input.resourceIsExist(result.AccessTo, funcName)
	if output.Status.Code != "" {
		return
	}

	resourceOld, output := dao.ResourceDAO.GetResourceByClientID(result.ClientID)
	if output.Status.Code != "" {
		return
	}

	resourceUpdated, output := dao.ResourceDAO.UpdateResourceByClientID(result)
	if output.Status.Code != "" {
		return
	}

	dataOld, _ := json.Marshal(resourceOld)
	dataNew, _ := json.Marshal(resourceUpdated)
	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestPUT, constanta.APIResource, tokenData.ClientID))
	input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionUpdate, constanta.TableResources,
		tokenData.ClientID, string(dataOld), string(dataNew),
		constanta.ActionIDUpdate))

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.UpdateDataSuccess
	output.Status.Detail = funcName
	output.Data = converter.DBDataToResouceGeneral(resourceUpdated)
	return
}

// created at 08-18-2022
func (input resourceService) readBodyAndValidateUpdate(request *http.Request, contextModel *model.ContextModel, validation func(input *req.ReqResourceUpdate) (output res.Payload)) (inputStruct req.ReqResourceUpdate, output res.Payload) {
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

// DeleteResource ------------------------------------------------------------------------------------------
// created at 08-19-2022
// updated at 08-27-2022
func (input resourceService) DeleteResource(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	funcName := "DeleteResource"

	tokenHeader := util.ReadHeader(request, constanta.TokenHeaderNameConstanta)
	tokenData, output := input.TokenChecker(tokenHeader, config.ApplicationConfiguration.GetClientCredentialsClientID(), funcName)
	if output.Status.Code != "" {
		return
	}

	params := util.GenerateQueryParam(request)
	paramClientID := params[constanta.CLIENTID_KEY]

	// tidak dapat menghapus resource gate
	if paramClientID == config.ApplicationConfiguration.GetClientCredentialsClientID() {
		output.Status.Code = constanta.CodeRequestFailed
		output.Status.Message = constanta.AccessForbidden
		output.Status.Detail = funcName
		return
	}

	resourceOld, output := dao.ResourceDAO.GetResourceByClientID(paramClientID)
	if output.Status.Code != "" {
		return
	}
	resource, output := dao.ResourceDAO.SoftDeleteResource(paramClientID)
	if output.Status.Code != "" {
		return
	}

	resources, output := dao.ResourceDAO.GetResources()

	dataOld, _ := json.Marshal(resourceOld)
	dataNew, _ := json.Marshal(resource)
	input.LogActivity(converter.ToLogActivity(time.Now(), constanta.RequestDELETE, constanta.APIResource, tokenData.ClientID))
	input.LogAuditSystem(converter.ToLogAuditSystem(time.Now(), constanta.ActionDelete, constanta.TableResources,
		tokenData.ClientID, string(dataOld), string(dataNew),
		constanta.ActionIDDelete))

	output.Status.Code = constanta.PayloadStatusCode
	output.Status.Message = constanta.DeleteDataSuccess
	output.Status.Detail = funcName
	output.Data = map[string]interface{}{
		constanta.FieldDataDeleted: converter.DBDataToResouceGeneral(resource),
		constanta.FieldDataNewest:  converter.DBDataToResouceGeneralArray(resources),
	}
	return
}
