package config

import (
	"github.com/Arijeet-webonise/go-react/pkg/logger"
	"github.com/spf13/viper"
)

//Config defines behaviour for constructing app configuration
type Config interface {
	ConstructAppConfig() *AppConfig
}

//AppConfig defines properties required by application
type AppConfig struct {
	DB             DbConfig
	Logger         logger.ILogger
	Port           string
	CSRFAuthkey    string
	SessionAuthkey string
}

// DbConfig wrapper for DB config
type DbConfig struct {
	DbUserName   string
	DbPassword   string
	DbHost       string
	DbName       string
	DbDriverName string
	DbDataSource string
}

//ConstructAppConfig prepares <AppConfig> from environment variables
func (appConfig *AppConfig) ConstructAppConfig() *AppConfig {
	viper.SetEnvPrefix("GR")
	viper.AutomaticEnv()

	appConfig.Port = appConfig.validateEnvVar("PORT")
	appConfig.CSRFAuthkey = appConfig.validateEnvVar("CSRF_AUTH_KEY")
	appConfig.SessionAuthkey = appConfig.validateEnvVar("SESSION_AUTH_KEY")
	return appConfig
}

//validateEnvVar fetches environment variable value for a given <key> if present
//else panics
func (appConfig *AppConfig) validateEnvVar(key string) string {
	value := viper.GetString(key)
	if value == "" {
		appConfig.Logger.Panicf("%s is not set", key)
	}
	return value
}