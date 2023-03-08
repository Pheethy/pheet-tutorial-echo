package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load(".env")/*Load Env*/
	if err != nil {
		log.Printf("please consider environment variable: %s", err)
	}

	psqlDB, err := sqlx.Open("postgres", os.Getenv("PSQL_DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	_ = psqlDB
}