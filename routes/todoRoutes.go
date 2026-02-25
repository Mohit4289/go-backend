package routes

import (
	"gin-quickstart/controllers/todo"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(r *gin.Engine) {
	r.POST("/addtodo", todo.AddTodo)
}
