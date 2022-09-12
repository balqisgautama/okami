package request

import (
	"github.com/go-playground/validator"
	"okami.auth.backend/constanta"
	response "okami.auth.backend/dto/out"
)

// created at 06-16-2022
type ReqRegist struct {
	Username       string `json:"username" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
	AdditionalInfo string `json:"additional_info" validate:"omitempty"`
	Locale         string `json:"locale" validate:"omitempty"`
}

// created at 06-16-2022
func ValidateReqRegist(inputStruct *ReqRegist) (output response.Payload) {
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output.Status.Message = err.Error()
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Detail = "RegistrationRequest.go"
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output.Status.Message = errV.Error()
			return
		}
		return
	}
	return
}
