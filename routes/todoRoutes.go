package routes

import (
	"gin-quickstart/controllers/todo"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(r *gin.Engine) {
	r.POST("/addtodo", todo.AddTodo)
	r.PUT("/updatetodo", todo.UpdateTodo)
	r.DELETE("/deletetodo/:id", todo.DeleteTodo)
}
