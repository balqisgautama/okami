package router

import (
	"context"
	"net/http"
	"strings"
	"time"

	"okami.auth.backend/config"
	"okami.auth.backend/constanta"
	"okami.auth.backend/model"
	errorModel "okami.auth.backend/model/error"
	logger "okami.auth.backend/model/logger"
	"okami.auth.backend/util"
	"okami.auth.backend/util/recovery"
)

func Middleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		util.CORSOriginHandler(&responseWriter)
		responseWriter.Header().Set("Content-Type", "application/json")
		if request.Method == "OPTIONS" {
			return
		} else {
			var contextModel *model.ContextModel
			defer func() {
				if r := recover(); r != nil {
					recovery.InputLog(errorModel.GenerateRecoverError(), contextModel.LoggerModel)
				}
			}()

			requestID := request.Header.Get(constanta.X_REQUEST_ID)
			if requestID == "" {
				requestID = util.GetUUID()
				request.Header.Set(constanta.X_REQUEST_ID, requestID)
			}

			var contextModels model.ContextModel

			contextModels.LoggerModel = logger.GenerateLogModel(config.ApplicationConfiguration.GetServerVersion() + " " + config.ApplicationConfiguration.GetServerResourceID()) //todo masukan ke config
			contextModels.LoggerModel.RequestID = requestID
			contextModels.LoggerModel.Class = "[Middleware.go,Middleware]"

			ctx := context.WithValue(request.Context(), constanta.ApplicationContextConstanta, &contextModels)
			request = request.WithContext(ctx)

			timestamp := time.Now()

			nextHandler.ServeHTTP(responseWriter, request)

			contextModel = request.Context().Value(constanta.ApplicationContextConstanta).(*model.ContextModel)
			contextModel.LoggerModel.Time = int64(time.Since(timestamp).Seconds())
			logMiddleware(contextModel.LoggerModel, request.RequestURI)
		}
	})
}

func logMiddleware(loggerModel logger.LoggerModel, requestURI string) {
	if !strings.Contains(requestURI, "health") && !strings.Contains(requestURI, "docs") && !strings.Contains(requestURI, "swagger") {
		recovery.InputLog(errorModel.GenerateNonErrorModel(), loggerModel)
	}
}
