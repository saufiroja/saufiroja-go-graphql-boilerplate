package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/saufiroja/go-graphql-boilerplate/config"
)

func NewPostgres(conf *config.AppConfig) *sql.Tx {
	host := conf.Postgres.DBHost
	port := conf.Postgres.DBPort
	user := conf.Postgres.DBUser
	pass := conf.Postgres.DBPass
	dbname := conf.Postgres.DBName

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// connection pool
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	log.Println("Connected to Postgres")

	return tx
}
