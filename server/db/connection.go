package db

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/lib/pq"
)

var host string = "localhost"
var user string = "postgres"
var password string = "123"
var dbname string = "Banco_api2"

const port = 5432

func Conectar() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	err = conn.Ping()

	return conn, err
}
