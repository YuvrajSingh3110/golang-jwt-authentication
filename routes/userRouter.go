package routes

import (
	"github.com/YuvrajSingh3110/golang_jwt_authentication/controllers"
	"github.com/YuvrajSingh3110/golang_jwt_authentication/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authenticate())
	routes.GET("/users", controllers.GetUsers())
	routes.GET("/users/:user_id", controllers.GetUserById())
}