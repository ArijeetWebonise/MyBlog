package config

import (
	"github.com/spf13/viper"
)

//Config defines behaviour for constructing app configuration
type Config interface {
	ConstructAppConfig() *AppConfig
}

//AppConfig defines properties required by application
type AppConfig struct {
	DB             DbConfig
	Port           string
	CSRFAuthkey    string
	SessionAuthkey string
	GraphQL        GraphQLConfig
}

//GraphQLConfig config properties for GraphQl
type GraphQLConfig struct {
	Pretty     bool
	GraphiQL   bool
	Playground bool
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
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()
	appConfig.Port = viper.GetString("PORT")
	viper.SetEnvPrefix("GR")
	viper.AutomaticEnv()
	appConfig.CSRFAuthkey = viper.GetString("CSRF_AUTH_KEY")
	appConfig.SessionAuthkey = viper.GetString("SESSION_AUTH_KEY")

	appConfig.DB.DbUserName = viper.GetString("DB_USERNAME")
	appConfig.DB.DbPassword = viper.GetString("DB_PASSWORD")
	appConfig.DB.DbHost = viper.GetString("DB_HOST")
	appConfig.DB.DbName = viper.GetString("DB_NAME")
	appConfig.DB.DbDriverName = viper.GetString("DB_DRIVER_NAME")
	appConfig.DB.DbDataSource = viper.GetString("DB_DATA_SOURCE")
	appConfig.GraphQL.GraphiQL = viper.GetBool("GRAPHIQL")
	appConfig.GraphQL.Playground = viper.GetBool("GRAPHQL_PLAYGROUND")
	appConfig.GraphQL.Pretty = viper.GetBool("GRAPHQL_PRETTY")
	return appConfig
}
