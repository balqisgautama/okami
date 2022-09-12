package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"log"
	"okami.auth.backend/config"
	serverconfig "okami.auth.backend/config/server"
	model "okami.auth.backend/model/logger"
	"okami.auth.backend/router"
	"okami.auth.backend/util"
	"os"
	"strconv"

	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	var arguments = "development"
	args := os.Args
	if len(args) > 1 {
		arguments = args[1]
	}

	config.GenerateConfiguration(arguments)
	util.SetLogger(config.ApplicationConfiguration.GetLogFile())
	serverconfig.SetServerAttribute()
	dbMigrate()

	err := serverconfig.ServerAttribute.DBConnection.Ping()
	if err != nil {
		logModel1 := model.GenerateLogModel(config.ApplicationConfiguration.GetServerVersion())
		logModel1.Status = 200
		logModel1.Message = err.Error()
		util.LogInfo(logModel1.ToLoggerObject())
	}

	defer func() {
		err := serverconfig.ServerAttribute.DBConnection.Close()
		if err != nil {
			logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServerVersion())
			logModel.Status = 500
			logModel.Message = "Failed to close DB Connection " + err.Error()
			util.LogError(logModel.ToLoggerObject())
		}
	}()

	logModel := model.GenerateLogModel(config.ApplicationConfiguration.GetServerVersion())
	logModel.Status = 200
	logModel.Message = "Server started on port : " + strconv.Itoa(config.ApplicationConfiguration.GetServerPort())
	util.LogInfo(logModel.ToLoggerObject())

	router.ApiController(config.ApplicationConfiguration.GetServerPort())
}

func dbMigrate() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./config/db"),
	}
	n, err := migrate.Exec(serverconfig.ServerAttribute.DBConnection, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
		os.Exit(3)
	} else {
		fmt.Println("Applied " + strconv.Itoa(n) + " migrations!")
	}
}
