package request

import (
	"github.com/go-playground/validator"
	"okami.auth.backend/constanta"
	response "okami.auth.backend/dto/out"
)

// created at 08-16-2022
// updated at 08-18-2022
type ReqResourceCreate struct {
	Surname  string `json:"surname" validate:"required,gte=5"`
	Nickname string `json:"nickname" validate:"required,gte=4"`
	AccessTo string `json:"access_to" validate:"required,gte=4"`
}

// created at 08-16-2022
func ValidateReqResourceCreate(inputStruct *ReqResourceCreate) (output response.Payload) {
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output.Status.Message = err.Error()
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Detail = "ResourceRequest.go"
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output.Status.Message = errV.Error()
			return
		}
		return
	}
	return
}

// created at 08-18-2022
type ReqResourceUpdate struct {
	ClientID string `json:"client_id" validate:"required,gte=10"`
	Surname  string `json:"surname" validate:"required,gte=5"`
	AccessTo string `json:"access_to" validate:"required,gte=4"`
}

// created at 08-18-2022
func ValidateReqResourceUpdate(inputStruct *ReqResourceUpdate) (output response.Payload) {
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output.Status.Message = err.Error()
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Detail = "ResourceRequest.go"
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output.Status.Message = errV.Error()
			return
		}
		return
	}
	return
}
