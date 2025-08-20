package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
type Book struct {
    ID       string  `json:"id"`
    Book_Name     string  `json:"book"`
	Author   string  `json:"author"`
}
var books = []Book{
	{ID: "1", Book_Name: "Gachiakuta", Author: "Kei Urana"},
	{ID: "2", Book_Name: "Tokyo Revengers", Author: "Ken Wakui"},
	{ID: "3", Book_Name: "Chainsaw Man", Author: "Tatsuki Fujimoto"},
}
func getBooks(c *gin.Context) {
	BookQuery := c.Query("id")
	if BookQuery != "" {
		filteredBooks := []Book{}
		for _, book := range books {
			if fmt.Sprint(book.ID) == BookQuery {
				filteredBooks = append(filteredBooks, book)
			}
		}
		c.JSON(http.StatusOK, filteredBooks)
		return
	}
	c.JSON(http.StatusOK, books)
}
func main() {
    r := gin.Default()
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy"})
    })
	api := r.Group("/api/v1")
	{
		api.GET("/books", getBooks)
	}
	r.Run(":8080")
}