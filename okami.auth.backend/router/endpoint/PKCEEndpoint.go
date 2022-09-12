package endpoint

import (
	"net/http"
	"okami.auth.backend/constanta"
	service "okami.auth.backend/service/authServer"
)

type pkceEndpoint struct {
	AbstractEndpoint
}

var PKCEEndpoint = pkceEndpoint{AbstractEndpoint{
	FileName: "endpoint/authServer/PKCEEndpoint.go",
}}

// PKCEStep1 ----------------------------------------------------------------------
// created at 08-26-2022
func (input pkceEndpoint) PKCEStep1(responseWriter http.ResponseWriter, request *http.Request) {
	funcName := "PKCEStep1"

	switch request.Method {
	case constanta.RequestPOST:
		input.ServeEndpoint(funcName, responseWriter, request, service.PKCEService.PKCEStep1)
		break
	}
}

// PKCEStep2 ----------------------------------------------------------------------
// created at 08-30-2022
func (input pkceEndpoint) PKCEStep2(responseWriter http.ResponseWriter, request *http.Request) {
	funcName := "PKCEStep2"

	switch request.Method {
	case constanta.RequestPOST:
		input.ServeEndpoint(funcName, responseWriter, request, service.PKCEService.PKCEStep2)
		break
	}
}

// PKCEStep3 ----------------------------------------------------------------------
// created at 08-31-2022
func (input pkceEndpoint) PKCEStep3(responseWriter http.ResponseWriter, request *http.Request) {
	funcName := "PKCEStep3"

	switch request.Method {
	case constanta.RequestPOST:
		input.ServeEndpoint(funcName, responseWriter, request, service.PKCEService.PKCEStep3)
		break
	}
}
