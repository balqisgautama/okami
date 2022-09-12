package endpoint

import (
	"net/http"
	"okami.auth.backend/constanta"
	service "okami.auth.backend/service/authServer"
)

type resourceEndpoint struct {
	AbstractEndpoint
}

var ResourceEndpoint = resourceEndpoint{AbstractEndpoint{
	FileName: "endpoint/authServer/ResourceEndpoint.go",
}}

// created at 08-16-2022
func (input resourceEndpoint) CRUDResource(responseWriter http.ResponseWriter, request *http.Request) {
	funcName := "CRUDResource"

	switch request.Method {
	case constanta.RequestPOST:
		input.ServeEndpoint(funcName, responseWriter, request, service.ResourceService.CreateResource)
		break
	case constanta.RequestGET:
		input.ServeEndpoint(funcName, responseWriter, request, service.ResourceService.ReadResource)
		break
	case constanta.RequestPUT:
		input.ServeEndpoint(funcName, responseWriter, request, service.ResourceService.UpdateResource)
		break
	case constanta.RequestDELETE:
		input.ServeEndpoint(funcName, responseWriter, request, service.ResourceService.DeleteResource)
		break
	}
}
