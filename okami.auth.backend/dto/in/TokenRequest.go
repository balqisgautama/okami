package request

import (
	"github.com/go-playground/validator"
	"okami.auth.backend/constanta"
	response "okami.auth.backend/dto/out"
)

// created at 08-18-2022
type ReqTokenResource struct {
	Token    string `json:"token" validate:"required,gte=20"`
	ClientID string `json:"client_id" validate:"required,gte=10"`
}

// created at 08-18-2022
func ValidateReqTokenResource(inputStruct *ReqTokenResource) (output response.Payload) {
	fileName = "TokenRequest.go"
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
