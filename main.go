package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Todos bool   `json:"todos"`
}

var todos = []todo{
	{ID: "1", Name: "wash some clothes", Todos: true},
	{ID: "2", Name: "Read BIO 101", Todos: false},
	{ID: "3", Name: "go shopping", Todos: true},
}

func getTodos(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, todos)

}

func addTodos(context *gin.Context) {
	var todo todo
	err := context.BindJSON(&todo)
	if err != nil {
		return
	}

	todos = append(todos, todo)
	context.IndentedJSON(http.StatusCreated, todos)
}

func getTodoId(id string) (*todo, error) {

	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoId(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	context.IndentedJSON(http.StatusOK, todo)

}

func toggleTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoId(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Todos = !todo.Todos

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todo", getTodos)
	router.POST("/todo", addTodos)
	router.GET("/todo/:id", getTodo)
	router.PATCH("/todo/:id", toggleTodo)
	router.Run("localhost:8080")
}
