package request

import (
	"github.com/go-playground/validator"
	"okami.auth.backend/constanta"
	response "okami.auth.backend/dto/out"
)

// created at 06-21-2022
type ReqEmail struct {
	Username string `json:"username" validate:"required"`
	EmailTo  string `json:"email_to" validate:"required,email"`
}

// created at 06-21-2022
func ValidateReqEmail(inputStruct *ReqEmail) (output response.Payload) {
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output.Status.Message = err.Error()
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Detail = "EmailRequest.go"
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output.Status.Message = errV.Error()
			return
		}
	}
	return
}
