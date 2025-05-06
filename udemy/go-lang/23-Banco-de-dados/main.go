package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	stringConection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", stringConection)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conection is opened")

}
