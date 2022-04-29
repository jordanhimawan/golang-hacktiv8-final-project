package router

import (
	"sesi-final-project/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", controllers.UserRegister)
		userRoute.POST("/login", controllers.UserLogin)
	}

	// photoRoute := r.Group("/photo")
	// {
	// 	photoRoute.Use(middlewares.Auth())
	// 	photoRoute.POST("/", controllers.CreateProduct)
	// 	photoRoute.PUT("/:productId", middlewares.Authorization(), controllers.UpdateProduct)
	// }

	return r
}
