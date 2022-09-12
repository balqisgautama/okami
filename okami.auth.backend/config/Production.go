package config

type ProductionConfig struct {
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

func (input ProductionConfig) GetServerHost() string {
	return input.Server.Host
}
func (input ProductionConfig) GetServerPort() int {
	return convertStringParamToInt("Server Port", input.Server.Port)
}
func (input ProductionConfig) GetServerVersion() string {
	return input.Server.Version
}
func (input ProductionConfig) GetServerResourceID() string {
	return input.Server.ResourceID
}
func (input ProductionConfig) GetPostgreSQLAddress() string {
	return input.Postgresql.Address
}
func (input ProductionConfig) GetPostgreSQLParam() string {
	return input.Postgresql.Param
}
func (input ProductionConfig) GetPostgreSQLMaxOpenConnection() int {
	return input.Postgresql.MaxOpenConnection
}
func (input ProductionConfig) GetPostgreSQLMaxIdleConnection() int {
	return input.Postgresql.MaxIdleConnection
}
func (input ProductionConfig) GetPostgreSQLAddressView() string {
	return input.PostgresqlView.Address
}
func (input ProductionConfig) GetPostgreSQLParamView() string {
	return input.PostgresqlView.Param
}
func (input ProductionConfig) GetPostgreSQLMaxOpenConnectionView() int {
	return input.PostgresqlView.MaxOpenConnection
}
func (input ProductionConfig) GetPostgreSQLMaxIdleConnectionView() int {
	return input.PostgresqlView.MaxIdleConnection
}
func (input ProductionConfig) GetRedisHost() string {
	return input.Redis.Host
}
func (input ProductionConfig) GetRedisPort() int {
	return convertStringParamToInt("Redis Port", input.Redis.Port)
}
func (input ProductionConfig) GetRedisDB() int {
	return convertStringParamToInt("Redis DB", input.Redis.Db)
}
func (input ProductionConfig) GetRedisPassword() string {
	return input.Redis.Password
}
func (input ProductionConfig) GetRedisTimeout() int {
	return input.Redis.Timeout
}
func (input ProductionConfig) GetRedisRequestVolumeThreshold() int {
	return input.Redis.RequestVolumeThreshold
}
func (input ProductionConfig) GetRedisSleepWindow() int {
	return input.Redis.SleepWindow
}
func (input ProductionConfig) GetRedisErrorPercentThreshold() int {
	return input.Redis.ErrorPercentThreshold
}
func (input ProductionConfig) GetRedisMaxConcurrentRequests() int {
	return input.Redis.MaxConcurrentRequests
}
func (input ProductionConfig) GetClientCredentialsClientID() string {
	return input.ClientCredentials.ClientID
}
func (input ProductionConfig) GetClientCredentialsClientSecret() string {
	return input.ClientCredentials.ClientSecret
}
func (input ProductionConfig) GetClientCredentialsSecretKey() string {
	return input.ClientCredentials.SecretKey
}
func (input ProductionConfig) GetLogFile() []string {
	return input.LogFile
}
func (input ProductionConfig) GetJWTToken() JWTKey {
	return JWTKey{
		JWT:      input.JWTKey.JWT,
		Internal: input.JWTKey.Internal,
	}
}

func (input ProductionConfig) GetServerPrefixPath() string {
	return input.Server.PrefixPath
}

func (input ProductionConfig) GetClientCredentialsAuthUserID() int64 {
	return input.ClientCredentials.AuthUserID
}

func (input ProductionConfig) GetKafka() Kafka {
	return Kafka{
		URL:       input.Kafka.URL,
		GroupID:   input.Kafka.GroupID,
		Partition: input.Kafka.Partition,
	}
}
