package router

import (
	"sesi-final-project/controllers"
	"sesi-final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", controllers.UserRegister)
		userRoute.POST("/login", controllers.UserLogin)
		userRoute.PUT("/:userId", controllers.UpdateUser)
		userRoute.DELETE("/:userId", controllers.DeleteUser)
	}

	photoRoute := r.Group("/photos")
	{
		photoRoute.Use(middlewares.Auth())
		photoRoute.POST("/", controllers.CreatePhoto)
		photoRoute.GET("/", controllers.GetPhotos)
		photoRoute.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		userRoute.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeleteUser)
	}

	commentRoute := r.Group("/comments")
	{
		commentRoute.Use(middlewares.Auth())
		commentRoute.POST("/", controllers.CreateComment)
		commentRoute.GET("/", controllers.GetComments)
		commentRoute.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRoute.DELETE("/:commentId", controllers.DeleteComment)
	}

	socialMediaRoute := r.Group("/socialmedias")
	{
		socialMediaRoute.Use(middlewares.Auth())
		socialMediaRoute.POST("/", controllers.CreatePhoto)
		socialMediaRoute.GET("/", controllers.GetSocialMedia)
		socialMediaRoute.PUT("/:socialMediaId", middlewares.CommentAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRoute.DELETE("/:socialMediaId", controllers.DeleteSocialMedia)
	}

	return r
}
