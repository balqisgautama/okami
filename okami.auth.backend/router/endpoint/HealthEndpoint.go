package endpoint

import (
	"net/http"
	"okami.auth.backend/constanta"
	"okami.auth.backend/service"
)

type healthEndpoint struct {
	AbstractEndpoint
}

var HealthEndpoint healthEndpoint

func (input healthEndpoint) HealthStatus(responseWriter http.ResponseWriter, request *http.Request) {
	funcName := "healthEndpoint"
	input.FileName = "HealthEndpoint.go"

	switch request.Method {
	case constanta.RequestGET:
		input.ServeEndpoint(funcName, responseWriter, request, service.HealthService.GetHealthStatus)
		break
	}
}
