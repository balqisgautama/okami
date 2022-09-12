package serverconfig

import (
	"database/sql"
	"github.com/bukalapak/go-redis"
	"okami.auth.backend/config"
	dbconfig "okami.auth.backend/config/db"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
)

type serverAttribute struct {
	Version          string
	DBConnection     *sql.DB
	DBConnectionView *sql.DB
	RedisClient      *redis.Client
}

var ServerAttribute serverAttribute

func SetServerAttribute() {
	dbParam := config.ApplicationConfiguration.GetPostgreSQLParam()
	dbConnection := config.ApplicationConfiguration.GetPostgreSQLAddress()
	dbMaxOpenConnection := config.ApplicationConfiguration.GetPostgreSQLMaxOpenConnection()
	dbMaxIdleConnection := config.ApplicationConfiguration.GetPostgreSQLMaxIdleConnection()
	ServerAttribute.DBConnection = dbconfig.GetDbConnection(dbParam, dbConnection, dbMaxOpenConnection, dbMaxIdleConnection)

	dbParamView := config.ApplicationConfiguration.GetPostgreSQLParamView()
	dbConnectionView := config.ApplicationConfiguration.GetPostgreSQLAddressView()
	dbMaxOpenConnectionView := config.ApplicationConfiguration.GetPostgreSQLMaxOpenConnectionView()
	dbMaxIdleConnectionView := config.ApplicationConfiguration.GetPostgreSQLMaxIdleConnectionView()
	ServerAttribute.DBConnectionView = dbconfig.GetDbConnection(dbParamView, dbConnectionView, dbMaxOpenConnectionView, dbMaxIdleConnectionView)

	redisHost := config.ApplicationConfiguration.GetRedisHost()
	redisDB := config.ApplicationConfiguration.GetRedisDB()
	redisPassword := config.ApplicationConfiguration.GetRedisPassword()
	redisPort := config.ApplicationConfiguration.GetRedisPort()
	redisTimeout := config.ApplicationConfiguration.GetRedisTimeout()
	redisVolumeThreshold := config.ApplicationConfiguration.GetRedisRequestVolumeThreshold()
	redisSleepWindow := config.ApplicationConfiguration.GetRedisSleepWindow()
	redisErrorPercentThreshold := config.ApplicationConfiguration.GetRedisErrorPercentThreshold()
	redisMaxConcurrentRequest := config.ApplicationConfiguration.GetRedisMaxConcurrentRequests()

	optCB := &hystrix.CommandConfig{
		Timeout:                redisTimeout,
		RequestVolumeThreshold: redisVolumeThreshold,
		SleepWindow:            redisSleepWindow,
		ErrorPercentThreshold:  redisErrorPercentThreshold,
		MaxConcurrentRequests:  redisMaxConcurrentRequest,
	}

	ServerAttribute.RedisClient = getRedisClient(redisHost, redisPort, redisDB, redisPassword, optCB)
	ServerAttribute.Version = config.ApplicationConfiguration.GetServerVersion()
}

func getRedisClient(host string, port int, db int, password string, optCB *hystrix.CommandConfig) *redis.Client {
	redisAddress := host + ":" + strconv.Itoa(port)
	opts := &redis.Options{
		CircuitBreaker: optCB,
		Addr:           redisAddress,
		// Password:       password,
		DB: db,
	}
	return redis.NewClient(opts)
}
