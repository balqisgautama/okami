package endpoint

import (
	"net/http"
	"okami.auth.backend/constanta"
	service "okami.auth.backend/service/authServer"
)

type emailEndpoint struct {
	AbstractEndpoint
}

var EmailEndpoint emailEndpoint

// ValidateEmailEndpoint ------------------------------------------------------------------------------------------
// created at 08-23-2022
func (input emailEndpoint) ValidateEmailEndpoint(responseWriter http.ResponseWriter, request *http.Request) {
	funcName := "ValidateEmailEndpoint"
	input.FileName = "EmailEndpoint.go"

	switch request.Method {
	case constanta.RequestGET:
		input.ServeEndpoint(funcName, responseWriter, request, service.UserActivationService.ValidateUserActivation)
		break
	}
}

// ResendEmailEndpoint ------------------------------------------------------------------------------------------
// created at 08-23-2022
func (input emailEndpoint) ResendEmailEndpoint(responseWriter http.ResponseWriter, request *http.Request) {
	funcName := "ResendEmailEndpoint"
	input.FileName = "EmailEndpoint.go"

	switch request.Method {
	case constanta.RequestGET:
		input.ServeEndpoint(funcName, responseWriter, request, service.UserActivationService.ResendUserActivation)
		break
	}
}
