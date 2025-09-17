package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	_"github.com/lib/pq"
	"log"
	"os"
	"net/http"
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
	r := gin.Default()
	r.GET("/health",func(c *gin.Context){
		err := db.Ping()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "unhealthy", "error": err})
			return
		}
		c.JSON(200, gin.H{"message": "healthy"})
	})

	r.Run(":8080")
}