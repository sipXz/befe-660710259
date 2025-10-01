package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	_"github.com/lib/pq"
	"log"
	"os"
	"net/http"
	"time"
)
var db *sql.DB
type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	ISBN      string    `json:"isbn"`
	Year      int       `json:"year"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
func getEnv(key, defaultValue string) string{
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
func initDB(){

	host := getEnv("DB_HOST", "localhost")
	name := getEnv("DB_NAME", "bookstore")
	user := getEnv("DB_USER", "bookstore_user")
	password := getEnv("DB_PASSWORD", "your_strong_password")
	port := getEnv("DB_PORT", "5432")
	conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	// fmt.Println(conSt)
	var err error
	db,err = sql.Open("postgres", conSt)
	if err != nil {
		log.Fatal("Failed to open Database")
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}
	log.Println("successfully connected to database")
}
func getAllBooks(c *gin.Context){
	var rows *sql.Rows
	var err error
	rows, err = db.Query("SELECT id, title, author, isbn, year, price, created_at, update_at FROM books")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close() // ต้องปิด rows เสมอ เพื่อคืน Connection กลับ pool
    var books []Book
    for rows.Next() {
        var book Book
        err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Year, &book.Price, &book.CreatedAt, &book.UpdatedAt)
        if err != nil {
            // handle error
        }
        books = append(books, book)
    }
	if books == nil {
		books = []Book{}
	}

	c.JSON(http.StatusOK, books)
}
func main(){
	initDB()
	defer db.Close()
	r := gin.Default()
	r.GET("/health",func(c *gin.Context){
		err := db.Ping()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "unhealthy", "error": err})
			return
		}
		c.JSON(200, gin.H{"message": "healthy"})
	})
	api := r.Group("/api/v1")
	{
		api.GET("/books", getAllBooks)

	}

	r.Run(":8080")
}