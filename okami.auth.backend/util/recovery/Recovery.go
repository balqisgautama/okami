package recovery

import (
	"okami.auth.backend/config"
	errorModel "okami.auth.backend/model/error"
	model "okami.auth.backend/model/logger"
	"okami.auth.backend/util"
)

//InputLog is user for usual logging that will save into main log file
//if there is an error it will save using logError method
func InputLog(err errorModel.ErrorModel, loggerModel model.LoggerModel) {
	if err.Error != nil {
		util.LogError(loggerModel.ToLoggerObject())
	} else {
		util.LogInfo(loggerModel.ToLoggerObject())
	}
}

func LogMessageInfo(msg string) {
	logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServerVersion())
	logModel.Status = 200
	logModel.Message = msg
	util.LogInfo(logModel.ToLoggerObject())
}

func LogMessageError(msg string) {
	logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServerVersion())
	logModel.Status = 200
	logModel.Message = msg
	util.LogError(logModel.ToLoggerObject())
}
