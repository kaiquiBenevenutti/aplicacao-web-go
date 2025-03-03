package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConectionDB() *sql.DB {
	connStr := "user=postgres dbname=web_aplicacao_go port=5432 sslmode=disable password=Kaiquirb1!"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
