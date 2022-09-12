package request

import (
	"github.com/go-playground/validator"
	"okami.auth.backend/constanta"
	response "okami.auth.backend/dto/out"
)

// created at 08-19-2022
type ReqUserGeneral struct {
	UserClientID     string `json:"user_client_id" validate:"omitempty,gte=10"`
	Username         string `json:"username" validate:"omitempty,gte=5"`
	Email            string `json:"email" validate:"omitempty,email"`
	Password         string `json:"password" validate:"omitempty,gte=7"`
	ResourceID       int64  `json:"resource_id" validate:"required,gt=0"`
	ResourceClientID string `json:"resource_client_id" validate:"required,gte=10"`
}

// created at 08-19-2022
func ValidateReqUserGeneral(inputStruct *ReqUserGeneral) (output response.Payload) {
	validate = validator.New()
	err := validate.Struct(inputStruct)
	if err != nil {
		output.Status.Message = err.Error()
		output.Status.Code = constanta.CodeValidationFailed
		output.Status.Detail = "UserRequest.go"
		if errV, ok := err.(*validator.InvalidValidationError); ok {
			output.Status.Message = errV.Error()
			return
		}
		return
	}
	return
}
