package request

import (
	"github.com/go-playground/validator"
	"okami.auth.backend/constanta"
	response "okami.auth.backend/dto/out"
)

// created at 07-05-2022
type ReqLoginGeneral struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// created at 07-05-2022
func ValidateReqLoginGeneral(inputStruct *ReqLoginGeneral) (output response.Payload) {
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output.Status.Message = err.Error()
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Detail = "LoginRequest.go"
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output.Status.Message = errV.Error()
			return
		}
		return
	}
	return
}
