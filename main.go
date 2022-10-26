package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        int    `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: 1, Item: "Pergi Ke kampus", Completed: false},
	{ID: 2, Item: "Pergi Ke kampus", Completed: false},
}

func getTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, todos)
}

func oneTodo(id int) (*todo, error) {
	//kenapa pakai *todo agar return nya bisa nil
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Print("Error get id param")
	}
	todo, err := oneTodo(id)

	if err != nil {
		fmt.Print("error get todo")
	}

	ctx.JSON(http.StatusOK, todo)
}

func addTodo(ctx *gin.Context) {
	var newTodo todo

	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
		fmt.Println(err)
		return
	}

	todos = append(todos, newTodo)
	ctx.IndentedJSON(http.StatusCreated, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)

	router.Run()
}
