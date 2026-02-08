package routes

import (
	"gin-quickstart/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/createacc", controllers.CreateAcc)
	r.POST("/loginacc", controllers.LoginAcc)
}
