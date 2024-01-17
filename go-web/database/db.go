package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func ConnectDatabase() *sql.DB {
	var envs map[string]string

	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatalf("Error loading .env")
	}

	ConnectionString := envs["DATABASE_URL"]

	fmt.Println("Connecting to database...")
	db, err := sql.Open("mysql", ConnectionString)

	if err != nil {
		fmt.Printf("Unable to connect to database %v\n", err)
		panic(err.Error())
	} else {
		fmt.Println("Connected")
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS go_web;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("use go_web;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY, 
		title VARCHAR(50) UNIQUE NOT NULL, 
		body TEXT
		);`)
	if err != nil {
		panic(err)
	}

	return db
}
