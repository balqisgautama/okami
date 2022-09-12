package endpoint

import (
	"net/http"
	"okami.auth.backend/constanta"
	service "okami.auth.backend/service/authServer"
)

type userEndpoint struct {
	AbstractEndpoint
}

var UserEndpoint = userEndpoint{AbstractEndpoint{
	FileName: "endpoint/authServer/UserEndpoint.go",
}}

// created at 08-19-2022
func (input userEndpoint) CRUDUser(responseWriter http.ResponseWriter, request *http.Request) {
	funcName := "CRUDUser"

	switch request.Method {
	case constanta.RequestPOST:
		input.ServeEndpoint(funcName, responseWriter, request, service.UserService.CreateUser)
		break
	case constanta.RequestGET:
		input.ServeEndpoint(funcName, responseWriter, request, service.UserService.ReadUser)
		break
	case constanta.RequestPUT:
		input.ServeEndpoint(funcName, responseWriter, request, service.UserService.UpdateUser)
		break
	case constanta.RequestDELETE:
		input.ServeEndpoint(funcName, responseWriter, request, service.UserService.DeleteUser)
		break
	}
}
