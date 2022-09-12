package config

import (
	model "okami.auth.backend/model/logger"
	"okami.auth.backend/util"
	"os"
	"strconv"
)

type DevelopmentConfig struct {
	Configuration
	Server struct {
		Host       string `envconfig:"OAUTH_HOST"`
		Port       string `envconfig:"OAUTH_PORT"`
		Version    string `json:"version"`
		ResourceID string `envconfig:"OAUTH_RESOURCE_ID"`
		PrefixPath string `json:"prefix_path"`
	} `json:"server"`
	Postgresql struct {
		Address           string `envconfig:"OAUTH_DB_CONNECTION"`
		Param             string `envconfig:"OAUTH_DB_SCHEMA"`
		MaxOpenConnection int    `json:"max_open_connection"`
		MaxIdleConnection int    `json:"max_idle_connection"`
	} `json:"postgresql"`
	PostgresqlView struct {
		Address           string `envconfig:"OAUTH_DB_VIEW_CONNECTION"`
		Param             string `envconfig:"OAUTH_DB_VIEW_SCHEMA"`
		MaxOpenConnection int    `json:"max_open_connection"`
		MaxIdleConnection int    `json:"max_idle_connection"`
	} `json:"postgresql_view"`
	Redis struct {
		Host                   string `envconfig:"OAUTH_REDIS_HOST"`
		Port                   string `envconfig:"OAUTH_REDIS_PORT"`
		Db                     string `envconfig:"OAUTH_REDIS_DB"`
		Password               string `envconfig:"OAUTH_REDIS_PASSWORD"`
		Timeout                int    `json:"timeout"`
		RequestVolumeThreshold int    `json:"request_volume_threshold"`
		SleepWindow            int    `json:"sleep_window"`
		ErrorPercentThreshold  int    `json:"error_percent_threshold"`
		MaxConcurrentRequests  int    `json:"max_concurrent_requests"`
	} `json:"redis"`
	ClientCredentials struct {
		ClientID     string `envconfig:"OAUTH_CLIENTID"`
		ClientSecret string `envconfig:"OAUTH_CLIENT_SECRET"`
		SecretKey    string `envconfig:"OAUTH_SIGNATURE_KEY"`
		AuthUserID   int64  `envconfig:"OAUTH_AUTH_USERID"`
	} `json:"client_credentials"`
	LogFile []string `json:"log_file"`
	JWTKey  struct {
		JWT      string `envconfig:"OAUTH_JWT_KEY"`
		Internal string `envconfig:"OAUTH_INTERNAL_KEY"`
	} `json:"jwt_key"`
	Kafka struct {
		URL       string `envconfig:"OAUTH_KAFKA"`
		GroupID   string `json:"group_id"`
		Partition int    `json:"partition"`
	} `json:"kafka"`
}

func (input DevelopmentConfig) GetServerHost() string {
	return input.Server.Host
}
func (input DevelopmentConfig) GetServerPort() int {
	return convertStringParamToInt("Server Port", input.Server.Port)
}
func (input DevelopmentConfig) GetServerVersion() string {
	return input.Server.Version
}
func (input DevelopmentConfig) GetServerResourceID() string {
	return input.Server.ResourceID
}
func (input DevelopmentConfig) GetPostgreSQLAddress() string {
	return input.Postgresql.Address
}
func (input DevelopmentConfig) GetPostgreSQLParam() string {
	return input.Postgresql.Param
}
func (input DevelopmentConfig) GetPostgreSQLMaxOpenConnection() int {
	return input.Postgresql.MaxOpenConnection
}
func (input DevelopmentConfig) GetPostgreSQLMaxIdleConnection() int {
	return input.Postgresql.MaxIdleConnection
}
func (input DevelopmentConfig) GetPostgreSQLAddressView() string {
	return input.PostgresqlView.Address
}
func (input DevelopmentConfig) GetPostgreSQLParamView() string {
	return input.PostgresqlView.Param
}
func (input DevelopmentConfig) GetPostgreSQLMaxOpenConnectionView() int {
	return input.PostgresqlView.MaxOpenConnection
}
func (input DevelopmentConfig) GetPostgreSQLMaxIdleConnectionView() int {
	return input.PostgresqlView.MaxIdleConnection
}
func (input DevelopmentConfig) GetRedisHost() string {
	return input.Redis.Host
}
func (input DevelopmentConfig) GetRedisPort() int {
	return convertStringParamToInt("Redis Port", input.Redis.Port)
}
func (input DevelopmentConfig) GetRedisDB() int {
	return convertStringParamToInt("Redis DB", input.Redis.Db)
}
func (input DevelopmentConfig) GetRedisPassword() string {
	return input.Redis.Password
}
func (input DevelopmentConfig) GetRedisTimeout() int {
	return input.Redis.Timeout
}
func (input DevelopmentConfig) GetRedisRequestVolumeThreshold() int {
	return input.Redis.RequestVolumeThreshold
}
func (input DevelopmentConfig) GetRedisSleepWindow() int {
	return input.Redis.SleepWindow
}
func (input DevelopmentConfig) GetRedisErrorPercentThreshold() int {
	return input.Redis.ErrorPercentThreshold
}
func (input DevelopmentConfig) GetRedisMaxConcurrentRequests() int {
	return input.Redis.MaxConcurrentRequests
}
func (input DevelopmentConfig) GetClientCredentialsClientID() string {
	return input.ClientCredentials.ClientID
}
func (input DevelopmentConfig) GetClientCredentialsClientSecret() string {
	return input.ClientCredentials.ClientSecret
}
func (input DevelopmentConfig) GetClientCredentialsSecretKey() string {
	return input.ClientCredentials.SecretKey
}
func (input DevelopmentConfig) GetClientCredentialsAuthUserID() int64 {
	return input.ClientCredentials.AuthUserID
}
func (input DevelopmentConfig) GetLogFile() []string {
	return input.LogFile
}
func (input DevelopmentConfig) GetJWTToken() JWTKey {
	return JWTKey{
		JWT:      input.JWTKey.JWT,
		Internal: input.JWTKey.Internal,
	}
}

func (input DevelopmentConfig) GetServerPrefixPath() string {
	return input.Server.PrefixPath
}

func (input DevelopmentConfig) GetKafka() Kafka {
	return Kafka{
		URL:       input.Kafka.URL,
		GroupID:   input.Kafka.GroupID,
		Partition: input.Kafka.Partition,
	}
}

func convertStringParamToInt(key string, value string) int {
	intPort, err := strconv.Atoi(value)
	if err != nil {
		logModel := model.GenerateLogModel("-")
		logModel.Message = "Invalid " + key + " : " + err.Error()
		logModel.Status = 500
		util.LogError(logModel.ToLoggerObject())
		os.Exit(3)
	}
	return intPort
}
