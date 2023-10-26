package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	
	_ "github.com/lib/pq"
)

var db *sql.DB

type connection struct {
	Host string
	Port string
	User string
	Password string
	DBName string
}

func Init() {
	err := godotenv.Load("./.env")

	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err.Error())
		return
	}

	connInfo := connection {
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
		User: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),	
	}

	db, err = sql.Open("postgres", connToString(connInfo))
	if err != nil {
		fmt.Printf("Error connecting to the db: %s\n", err.Error())
	} else {
		fmt.Printf("DB is open\n")
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error could not ping database: %s\n", err.Error())
		return
	} else {
		fmt.Printf("DB pinged successfully\n")
	}
}

func connToString(info connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled", 
		info.Host, info.Port, info.User, info.Password, info.DBName)
}