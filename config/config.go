package config

import "github.com/joho/godotenv"

type AppConfig struct {
	App struct {
		Env string
	}
	Postgres struct {
		DBName string
		DBUser string
		DBPass string
		DBHost string
		DBPort string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	_ = godotenv.Load()

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initPostgres(appConfig)
	}

	return appConfig
}
