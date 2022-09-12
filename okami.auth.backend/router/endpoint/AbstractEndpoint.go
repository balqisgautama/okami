package endpoint

import (
	"context"
	"errors"
	"net/http"
	"runtime/debug"

	"okami.auth.backend/config"
	"okami.auth.backend/constanta"
	out "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	errorModel "okami.auth.backend/model/error"
	"okami.auth.backend/util"
)

type AbstractEndpoint struct {
	FileName string
}

func (input AbstractEndpoint) ServeEndpoint(funcName string, responseWriter http.ResponseWriter, request *http.Request, serveFunction func(*http.Request, *model.ContextModel) (out.Payload, map[string]string)) {
	serveEndpoint(input.FileName, funcName, responseWriter, request, serveFunction)
}

func serveEndpoint(fileName string, funcName string, responseWriter http.ResponseWriter, request *http.Request, serveFunction func(*http.Request, *model.ContextModel) (out.Payload, map[string]string)) {
	contextModel := request.Context().Value(constanta.ApplicationContextConstanta).(*model.ContextModel)

	ctx := context.WithValue(request.Context(), constanta.ApplicationContextConstanta, contextModel)
	request = request.WithContext(ctx)

	serve(fileName, funcName, responseWriter, request, serveFunction)
}

func serve(fileName string, funcName string, responseWriter http.ResponseWriter, request *http.Request, serve func(*http.Request, *model.ContextModel) (out.Payload, map[string]string)) {
	var err errorModel.ErrorModel
	var contextModel *model.ContextModel
	var output out.Payload
	var header map[string]string

	defer func() {
		if r := recover(); r != nil {
			err = errorModel.GenerateRecoverError()
			contextModel.LoggerModel.Message = string(debug.Stack())
		} else {
			if err.Error != nil {
				contextModel.LoggerModel.Class = "[" + err.FileName + "," + err.FuncName + "]"
				contextModel.LoggerModel.Code = err.Error.Error()
				if err.CausedBy != nil {
					contextModel.LoggerModel.Message = err.CausedBy.Error()
				} else {
					contextModel.LoggerModel.Message = err.Error.Error()
				}
			}
		}

		contextModel.LoggerModel.Status = err.Code
		ctx := context.WithValue(request.Context(), constanta.ApplicationContextConstanta, contextModel)
		request = request.WithContext(ctx)
		finish(responseWriter, err, contextModel, output)
	}()

	contextModel = request.Context().Value(constanta.ApplicationContextConstanta).(*model.ContextModel)
	contextModel.LoggerModel.Class = "[" + fileName + "," + funcName + "]"
	getDBSchema(contextModel)
	output, header = serve(request, contextModel)
	if err.Error != nil {
		return
	}

	setHeader(header, responseWriter)
}

func setHeader(header map[string]string, responseWriter http.ResponseWriter) {
	accessControlExpose := "Access-Control-Expose-Headers"
	exposeHeader := responseWriter.Header().Get(accessControlExpose)
	for key := range header {
		responseWriter.Header().Add(key, header[key])
		if exposeHeader == "" {
			exposeHeader = key
		} else {
			exposeHeader += ", " + key
		}
	}
	if exposeHeader != "" {
		responseWriter.Header().Set(accessControlExpose, exposeHeader)
	}
}

func finish(responseWriter http.ResponseWriter, err errorModel.ErrorModel, contextModel *model.ContextModel, output out.Payload) {
	if output.Status.Detail == constanta.ContentTypeHTML {
		responseWriter.Header().Set("Content-Type", constanta.ContentTypeHTML)
		responseWriter.Write([]byte(output.Status.Message))
	} else if err.Error != nil {
		writeErrorResponse(responseWriter, err, contextModel)
	} else {
		writeSuccessResponse(responseWriter, contextModel, output)
	}
}

func writeErrorResponse(responseWriter http.ResponseWriter, err errorModel.ErrorModel, contextModel *model.ContextModel) {
	if err.Code == 0 {
		responseWriter.WriteHeader(500)
		err.CausedBy = err.Error
		err.Error = errors.New("OGATE-SERVER-ERROR")
	} else {
		responseWriter.WriteHeader(err.Code)
	}

	errCode := err.Error.Error()

	errResponse := out.StatusResponse{
		Success: false,
		Code:    errCode,
	}

	responseMessage := out.APIResponse{
		Okami: out.OkamiMessage{
			Header: out.Header{
				RequestID: contextModel.LoggerModel.RequestID,
				Version:   config.ApplicationConfiguration.GetServerVersion(),
				Timestamp: util.GetTimeStamp(),
			},
			Payload: out.Payload{Status: errResponse}},
	}

	if err.AdditionalInformation != nil && len(err.AdditionalInformation) > 0 {
		responseMessage.Okami.Payload.Status.AdditionalInfo = err.AdditionalInformation
	}

	_, errorS := responseWriter.Write([]byte(responseMessage.String()))
	if errorS != nil {
		errModel := errorModel.GenerateUnknownError("AbstractEndpoint.go", "writeErrorResponse", errorS)
		contextModel.LoggerModel.Status = errModel.Code
		contextModel.LoggerModel.Code = errModel.Error.Error()
		contextModel.LoggerModel.Message = errorS.Error()
	}

	contextModel.LoggerModel.ByteOut = len([]byte(responseMessage.String()))
}

func writeSuccessResponse(responseWriter http.ResponseWriter, contextModel *model.ContextModel, output out.Payload) {
	success := true
	if output.Status.Code != constanta.PayloadStatusCode {
		success = false
	}
	output.Status.Success = success
	responseMessage := out.APIResponse{
		Okami: out.OkamiMessage{
			Header: out.Header{
				RequestID: contextModel.LoggerModel.RequestID,
				Version:   config.ApplicationConfiguration.GetServerVersion(),
				Timestamp: util.GetTimeStamp(),
			},
			Payload: output},
	}
	bodyMessage := responseMessage.String()

	responseWriter.WriteHeader(200)
	_, errorS := responseWriter.Write([]byte(bodyMessage))
	if errorS != nil {
		errModel := errorModel.GenerateUnknownError("AbstractEndpoint.go", "writeSuccessResponse", errorS)
		contextModel.LoggerModel.Status = errModel.Code
		contextModel.LoggerModel.Code = errModel.Error.Error()
		contextModel.LoggerModel.Message = errorS.Error()
	}
	contextModel.LoggerModel.ByteOut = len([]byte(responseMessage.String()))
}

func getDBSchema(model *model.ContextModel) {
	model.DBSchema = config.ApplicationConfiguration.GetPostgreSQLParam()
}
