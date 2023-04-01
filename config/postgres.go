package config

import (
	"os"
)

func initPostgres(conf *AppConfig) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	conf.Postgres.DBHost = host
	conf.Postgres.DBPort = port
	conf.Postgres.DBUser = user
	conf.Postgres.DBPass = pass
	conf.Postgres.DBName = dbname
}
