package endpoint

import (
	"net/http"
	"okami.auth.backend/constanta"
	serviceAuth "okami.auth.backend/service/authServer"
)

type tokenEndpoint struct {
	AbstractEndpoint
}

var TokenEndpoint tokenEndpoint

//func (input tokenEndpoint) TokenAuthFirebase(responseWriter http.ResponseWriter, request *http.Request) {
//	funcName := "TokenAuthFirebase"
//	input.FileName = "FirebaseEndpoint.go"
//
//	switch request.Method {
//	case constanta.RequestGET:
//		input.ServeEndpoint(funcName, responseWriter, request, service.FirebaseService.GenerateToken)
//		break
//	}
//}

// created at 08-18-2022
func (input tokenEndpoint) TokenResource(responseWriter http.ResponseWriter, request *http.Request) {
	funcName := "TokenResource"
	input.FileName = "TokenEndpoint.go"

	switch request.Method {
	case constanta.RequestGET:
		input.ServeEndpoint(funcName, responseWriter, request, serviceAuth.TokenService.GenerateTokenResource)
		break
	case constanta.RequestPOST:
		input.ServeEndpoint(funcName, responseWriter, request, serviceAuth.TokenService.VerifyTokenResource)
		break
	}
}
