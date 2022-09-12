package errorModel

import (
	"strconv"

	"okami.auth.backend/constanta"
)

var DefaultError map[string]ErrorClass

type ErrorClass struct {
	ErrorCode    string
	ErrorMessage string
}

func GenerateUnsupportedResponseTypeError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-DTO-001", fileName, funcName)
}

func GenerateAlreadyUseError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(401, "E-1-AUT-DTO-011", fileName, funcName)
}

func GenerateInactiveAuditSystem(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-SRV-009", fileName, funcName)
}

func GenerateMissingResourceIDError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-DTO-005", fileName, funcName)
}
func GenerateUnauthorizedClientError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(401, "E-1-AUT-SRV-001", fileName, funcName)
}
func GenerateUnknownError(fileName string, funcName string, causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-5-AUT-SRV-001", fileName, funcName, causedBy)
}
func GenerateInternalDBServerError(fileName string, funcName string, causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-5-AUT-DBS-001", fileName, funcName, causedBy)
}
func GenerateInvalidRequestError(fileName string, funcName string, causedBy error) ErrorModel {
	return GenerateErrorModel(400, "E-4-AUT-DTO-003", fileName, funcName, causedBy)
}

func GenerateMissingCodeChallengeError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-DTO-004", fileName, funcName)
}
func GenerateInvalidJWTCodeError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(401, "E-4-AUT-SRV-001", fileName, funcName)
}
func GenerateRecoverError() ErrorModel {
	return GenerateSimpleErrorModel(500, "E-5-AUT-SRV-001")
}
func GenerateAPIError(code int, err string) ErrorModel {
	return GenerateSimpleErrorModel(code, err)
}
func GenerateFieldHaveMaxLimitError(fileName string, funcName string, fieldName string, digitMax int) ErrorModel {
	errorParam := make([]errorParameter, 2)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	errorParam[1].ErrorParameterKey = "DigitMax"
	errorParam[1].ErrorParameterValue = strconv.Itoa(digitMax)
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-009", fileName, funcName, errorParam)
}
func GenerateFieldHaveMinLimitError(fileName string, funcName string, fieldName string, digitMin int) ErrorModel {
	errorParam := make([]errorParameter, 2)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	errorParam[1].ErrorParameterKey = "DigitMax"
	errorParam[1].ErrorParameterValue = strconv.Itoa(digitMin)
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-010", fileName, funcName, errorParam)
}

// updated at 04-09-2022
func GenerateFormatFieldError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]errorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(200, constanta.CodeValidationFailed, fileName, funcName, errorParam)
}

// updated at 04-10-2022
func GenerateEmptyDatabaseFieldError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]errorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(200, constanta.CodeFieldIsEmpty, fileName, funcName, errorParam)
}

func GenerateFieldFormatWithRuleError(fileName string, funcName string, ruleName string, fieldName string, additionalInfo string) ErrorModel {
	errorParam := make([]errorParameter, 3)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	errorParam[1].ErrorParameterKey = "RuleName"
	errorParam[1].ErrorParameterValue = ruleName
	errorParam[2].ErrorParameterKey = "AdditionalInformation"
	errorParam[2].ErrorParameterValue = additionalInfo
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-008", fileName, funcName, errorParam)
}
func GenerateAlreadyExistDataError(fileName string, funcName string, primaryField string) ErrorModel {
	errorParam := make([]errorParameter, 1)
	errorParam[0].ErrorParameterKey = "PrimaryField"
	errorParam[0].ErrorParameterValue = primaryField
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-002", fileName, funcName, errorParam)
}
func GenerateForbiddenAccessClientError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(403, "E-3-AUT-SRV-001", fileName, funcName)
}
func GenerateUnsupportedRequestParam(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-DTO-001", fileName, funcName)
}
func GenerateUnknownDataError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]errorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-003", fileName, funcName, errorParam)
}

func GenerateDataUsedError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]errorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-010", fileName, funcName, errorParam)
}

func GenerateChangePasswordNotValidError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]errorParameter, 1)
	errorParam[0].ErrorParameterKey = "Password"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-006", fileName, funcName, errorParam)
}
func GenerateInvalidActivationCodeError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-SRV-004", fileName, funcName)
}
func GenerateActivationCodeExpiredError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-SRV-005", fileName, funcName)
}
func GenerateDataLockedError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]errorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-SRV-007", fileName, funcName, errorParam)
}

func GenerateDateIsLessThanNowError(fileName string, funcName string, fieldName string) ErrorModel {
	errorParam := make([]errorParameter, 1)
	errorParam[0].ErrorParameterKey = "FieldName"
	errorParam[0].ErrorParameterValue = fieldName

	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-011", fileName, funcName, errorParam)
}

func GenerateDataNotFoundError(fileName string, funcName string, causedBy error) ErrorModel {
	return GenerateErrorModel(404, "E-4-AUT-SRV-008", fileName, funcName, causedBy)
}

func GenerateAssignTicketError(fileName string, funcName string, ticketNumber string) ErrorModel {
	errorParam := make([]errorParameter, 2)
	errorParam[0].ErrorParameterKey = "TicketNumber"
	errorParam[0].ErrorParameterValue = ticketNumber
	return GenerateErrorModelWithErrorParam(400, "E-4-AUT-DTO-011", fileName, funcName, errorParam)
}

func GenerateCSIsLockedError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-SRV-011", fileName, funcName)
}

func GenerateTicketIsLockedError(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-SRV-012", fileName, funcName)
}
func GenerateFailedSavingDataError(fileName string, funcName string, errorMessage string) ErrorModel {
	errorParam := make([]errorParameter, 1)
	errorParam[0].ErrorParameterKey = "ErrorMessage"
	errorParam[0].ErrorParameterValue = errorMessage
	return GenerateErrorModelWithErrorParam(500, "E-4-AUT-DBS-002", fileName, funcName, errorParam)
}

func GenerateForbiddenDeleteTicket(fileName string, funcName string) ErrorModel {
	return GenerateErrorModelWithoutCaused(400, "E-4-AUT-SRV-013", fileName, funcName)
}
