package errorModel

import "errors"

type ErrorModel struct {
	Code                  int
	Error                 error
	FileName              string
	FuncName              string
	CausedBy              error
	ErrorParameter        []errorParameter
	AdditionalInformation []string
}

type ErrorLogModel struct {
	Code         int
	ErrorMessage string
	FileName     string
	FuncName     string
}

type errorParameter struct {
	ErrorParameterKey   string
	ErrorParameterValue string
}

//type ErrorModel struct {
//	Code                  int
//	Error                 error
//	FileName              string
//	FuncName              string
//	CausedBy              error
//	ErrorParameter        []errorParameter
//	AdditionalInformation interface{}
//}
//
//

func GenerateErrorModel(code int, err string, fileName string, funcName string, causedBy error) ErrorModel {
	var errModel ErrorModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.FileName = fileName
	errModel.FuncName = funcName
	errModel.CausedBy = causedBy
	return errModel
}

func GenerateErrorModelWithoutCaused(code int, err string, fileName string, funcName string) ErrorModel {
	var errModel ErrorModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.FileName = fileName
	errModel.FuncName = funcName
	return errModel
}

func GenerateErrorModelWithErrorParam(code int, err string, fileName string, funcName string, errorParam []errorParameter) ErrorModel {
	var errModel ErrorModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.FileName = fileName
	errModel.FuncName = funcName
	errModel.ErrorParameter = errorParam
	return errModel
}

func GenerateSimpleErrorModel(code int, err string) ErrorModel {
	var errModel ErrorModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	return errModel
}

func GenerateNonErrorModel() ErrorModel {
	var errModel ErrorModel
	errModel.Code = 200
	errModel.Error = nil
	return errModel
}
