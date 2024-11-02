package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clear Room", Completed: true},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Interesting Video", Completed: true},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func helloWeb(context *gin.Context) {
	context.String(200, "Hello, World!")
}

func main() {
	router := gin.Default()
	router.GET("/", helloWeb)
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}
