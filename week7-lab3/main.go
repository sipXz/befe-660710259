package main

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)
var db *sql.DB

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
func initDB(){
	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")
	conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	// fmt.Println(conSt)
	db,err := sql.Open("postgres", conSt)
	if err != nil {
		log.Fatal("Failed to open Database")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}
	log.Println("successfully connected to database")
}

func main(){
	initDB()
}