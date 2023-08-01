package routes

import (
	"github.com/YuvrajSingh3110/golang_jwt_authentication/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("users/signup", controllers.Signup())
	routes.POST("users/login", controllers.Login())
}
