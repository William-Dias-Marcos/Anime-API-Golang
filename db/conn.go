package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "dev"
	password = "12345678ab"
	dbname = "animes"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo :=  "host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable"
	psqlInfo =  fmt.Sprintf(psqlInfo, host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to " + dbname)
	return db, nil
}