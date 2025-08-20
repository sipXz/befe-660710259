package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
func getUsers(c* gin.Context) {
	users := []User{{ID:"1",Name:"Sippakorn"}}
	c.JSON(200,users)
}

func main(){
	r := gin.Default()
	r.GET("/users", getUsers)
	r.Run(":8080")
}