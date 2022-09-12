package config

import (
	"fmt"
	"okami.auth.backend/constanta"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/tkanos/gonfig"
)

var ApplicationConfiguration Configuration

type Configuration interface {
	GetServerHost() string
	GetServerPort() int
	GetServerVersion() string
	GetServerResourceID() string
	GetServerPrefixPath() string
	GetPostgreSQLAddress() string
	GetPostgreSQLParam() string
	GetPostgreSQLMaxOpenConnection() int
	GetPostgreSQLMaxIdleConnection() int
	GetPostgreSQLAddressView() string
	GetPostgreSQLParamView() string
	GetPostgreSQLMaxOpenConnectionView() int
	GetPostgreSQLMaxIdleConnectionView() int
	GetRedisHost() string
	GetRedisPort() int
	GetRedisDB() int
	GetRedisPassword() string
	GetRedisTimeout() int
	GetRedisRequestVolumeThreshold() int
	GetRedisSleepWindow() int
	GetRedisErrorPercentThreshold() int
	GetRedisMaxConcurrentRequests() int
	GetClientCredentialsClientID() string
	GetClientCredentialsClientSecret() string
	GetClientCredentialsSecretKey() string
	GetClientCredentialsAuthUserID() int64
	GetLogFile() []string
	GetJWTToken() JWTKey
	GetLanguageDirectoryPath() string
	GetNexCareFrontend() NexCareFrontend
	GetAuthenticationServer() AuthenticationServer
	GetCommonPath() CommonPath
	GetAudit() Audit
	GetAlertServer() AlertServer
	GetAzure() Azure
	GetCDN() CDN
	GetElasticSearchConnectionString() string
	GetMasterData() MasterData
	GetKafka() Kafka
}

type AuthenticationServer struct {
	Host                string                     `json:"host"`
	PathRedirect        AuthenticationPathRedirect `json:"path_redirect"`
	EmailLinkActivation string                     `json:"email_link_activation"`
}

type AuthenticationPathRedirect struct {
	CheckToken        string             `json:"check_token"`
	AddResourceClient string             `json:"add_resource_client"`
	Authorize         string             `json:"authorize"`
	Verify            string             `json:"verify"`
	VerifyRegister    VerifyRegisterType `json:"verify_register"`
	ResetPassword     string             `json:"reset_password"`
	EmailPassword     string             `json:"email_password"`
	Token             string             `json:"token"`
	UserDetail        string             `json:"user_detail"`
	Logout            string             `json:"logout"`
	RegisterUser      string             `json:"register_user"`
	Activation        string             `json:"activation"`
}

type VerifyRegisterType struct {
	VerifyEmail string `json:"verify_email"`
	VerifyPhone string `json:"verify_phone"`
}

type NexCareFrontend struct {
	Host         string              `json:"host"`
	PathRedirect NexCarePathRedirect `json:"path_redirect"`
}

type NexCarePathRedirect struct {
	ResetPasswordPath string `json:"reset_password_path"`
	VerifyUserPath    string `json:"verify_user_path"`
	ActivationPath    string `json:"activation_path"`
}

type CommonPath struct {
	ResourceClients string `json:"resource_clients"`
	ResourceToken   string `json:"resource_token"`
}

type Audit struct {
	IsActive bool `json:"is_active"`
}
type JWTKey struct {
	JWT      string `json:"jwt"`
	Internal string `json:"internal"`
}

type AlertServer struct {
	Host         string              `json:"host"`
	PathRedirect NexCarePathRedirect `json:"path_redirect"`
}

type AlertServerPathRedirect struct {
	Alert string `json:"alert"`
}

type Azure struct {
	AccountName string
	AccountKey  string
	Host        string
	Suffix      string
}

type CDN struct {
	RootPath string `json:"root_path"`
}

type MasterData struct {
	Host         string `json:"host"`
	PathRedirect struct {
		CompanyProfile string `json:"company_profile"`
	} `json:"path_redirect"`
}

type Kafka struct {
	URL       string `json:"url"`
	GroupID   string `json:"group_id"`
	Partition int    `json:"partition"`
}

func GenerateConfiguration(arguments string) {
	var err error
	enviName := "OkamiConfiguration"

	switch arguments {
	case "production":
		temp := ProductionConfig{}
		err = gonfig.GetConf(os.Getenv(enviName)+constanta.PathEnv+"config_production.json", &temp)
		if err != nil {
			fmt.Print(err)
			os.Exit(2)
		}
		err = envconfig.Process(os.Getenv(enviName)+constanta.PathEnv+"config_production.json", &temp)
		if err != nil {
			fmt.Print(err)
			os.Exit(2)
		}
		ApplicationConfiguration = &temp
	case "staging":
		temp := StagingConfig{}
		err = gonfig.GetConf(os.Getenv(enviName)+constanta.PathEnv+"config_staging.json", &temp)
		if err != nil {
			fmt.Print(err)
			os.Exit(2)
		}
		err = envconfig.Process(os.Getenv(enviName)+constanta.PathEnv+"config_sandbox.json", &temp)
		if err != nil {
			fmt.Print(err)
			os.Exit(2)
		}
		ApplicationConfiguration = &temp
	default:
		temp := DevelopmentConfig{}
		err = gonfig.GetConf(os.Getenv(enviName)+constanta.PathEnv+"config_development.json", &temp)
		if err != nil {
			fmt.Print(err)
			os.Exit(2)
		}
		err = envconfig.Process(os.Getenv(enviName)+constanta.PathEnv+"config_development.json", &temp)
		if err != nil {
			fmt.Print(err)
			os.Exit(2)
		}
		ApplicationConfiguration = &temp
	}

	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
}
