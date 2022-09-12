package service

import (
	"net/http"
	serverconfig "okami.auth.backend/config/server"
	"okami.auth.backend/constanta"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/model"
	"time"
)

type healthService struct {
	FileName string
	AbstractService
}

var HealthService = healthService{FileName: "HealthService.go"}

func (input healthService) GetHealthStatus(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
	// funcName := "GetHealthStatus"
	serverStatus := constanta.StringUp
	err := serverconfig.ServerAttribute.DBConnection.Ping()
	DBStatus := "UP"
	if err != nil {
		DBStatus = constanta.StringFailed
	}
	output.Status.Code = constanta.PayloadStatusCode
	timestamp := time.Now().Format(constanta.DefaultTimeFormat)
	output.Data = map[string]string{
		"timestamp": timestamp,
		"server":    serverStatus,
		"database":  DBStatus,
	}
	return
}
