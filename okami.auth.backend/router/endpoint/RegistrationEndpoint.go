package endpoint

//
//import (
//	"net/http"
//	"okami.auth.backend/constanta"
//	"okami.auth.backend/service"
//)
//
//type registrationEndpoint struct {
//	AbstractEndpoint
//}
//
//var RegistrationEndpoint registrationEndpoint
//
//func (input registrationEndpoint) RegistrationEndpoint(responseWriter http.ResponseWriter, request *http.Request) {
//	funcName := "registrationEndpoint"
//	input.FileName = "RegistrationEndpoint.go"
//
//	switch request.Method {
//	case constanta.RequestPOST:
//		input.ServeEndpoint(funcName, responseWriter, request, service.RegistrationService.Registration)
//		break
//	}
//}
