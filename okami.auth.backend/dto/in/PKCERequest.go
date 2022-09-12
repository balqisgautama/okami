package request

import (
	"github.com/go-playground/validator"
	"okami.auth.backend/constanta"
	response "okami.auth.backend/dto/out"
)

type ReqPKCEStep1 struct {
	CodeChallenger string `json:"code_challenger" validate:"required,gte=10"`
	Type           string `json:"type" validate:"required"`
}

type ReqPKCEStep2 struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=7"`
}

type ReqPKCEStep3 struct {
	CodeVerifier string `json:"code_verifier" validate:"required,gte=7"`
}

// ValidateReqPKCEStep1 ---------------------------------------------------------------
// created at 08-26-2022
func ValidateReqPKCEStep1(inputStruct *ReqPKCEStep1) (output response.Payload) {
	fileName = "PKCERequest.go"
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output.Status.Message = err.Error()
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Detail = fileName
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output.Status.Message = errV.Error()
			return
		}
		return
	}
	return
}

// ValidateReqPKCEStep2 ---------------------------------------------------------------
// created at 08-30-2022
func ValidateReqPKCEStep2(inputStruct *ReqPKCEStep2) (output response.Payload) {
	fileName = "PKCERequest.go"
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output.Status.Message = err.Error()
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Detail = fileName
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output.Status.Message = errV.Error()
			return
		}
		return
	}
	return
}

// ValidateReqPKCEStep3 ---------------------------------------------------------------
// created at 08-31-2022
func ValidateReqPKCEStep3(inputStruct *ReqPKCEStep3) (output response.Payload) {
	fileName = "PKCERequest.go"
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output.Status.Message = err.Error()
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Detail = fileName
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output.Status.Message = errV.Error()
			output.Status.Detail = fileName
			return
		}
		return
	}
	return
}
