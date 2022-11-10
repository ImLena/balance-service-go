package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	Conn *sql.DB
}

func Initialize(host, username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, 5432, username, password, database)
	var conn *sql.DB
	var err error

	conn, err = sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	log.Printf("%v\n", err)
	log.Printf("%v \n", dsn)
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}
