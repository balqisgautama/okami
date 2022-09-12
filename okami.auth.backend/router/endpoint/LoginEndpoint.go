package endpoint

//import (
//	"net/http"
//	"okami.auth.backend/constanta"
//	service "okami.auth.backend/service/login"
//)
//
//type loginEndpoint struct {
//	AbstractEndpoint
//}
//
//var LoginEndpoint loginEndpoint
//
//func (input loginEndpoint) GeneralLoginEndpoint(responseWriter http.ResponseWriter, request *http.Request) {
//	funcName := "GeneralLoginEndpoint"
//	input.FileName = "LoginEndpoint.go"
//
//	switch request.Method {
//	case constanta.RequestPOST:
//		input.ServeEndpoint(funcName, responseWriter, request, service.LoginService.GeneralLogin)
//		break
//	}
//}
