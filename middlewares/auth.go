package middlewares

import (
	"fmt"
	"net/http"
	"sesi-final-project/database"
	"sesi-final-project/helpers"
	"sesi-final-project/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")
		bearer := strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "UNAUTHORIZED",
			})
			return
		}

		token := strings.Split(headerToken, " ")[1]

		id, email, err := helpers.ValidateToken(token)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "UNAUTHORIZED",
				"msg":    err.Error(),
			})
			return
		}

		c.Set("id", id)
		c.Set("email", email)
		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id := c.MustGet("id").(uint)
		productId := c.Param("productId")

		var photo models.Photo
		err := db.Debug().Where("id = ?", productId).Take(&photo).Error
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  "Photo not found!",
				"msg":    err.Error(),
			})
			return
		}

		if photo.UserId != id {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  "FORBIDDEN",
				"msg":    "You don't have right to access the resources",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id := c.MustGet("id").(uint)
		productId := c.Param("productId")

		var photo models.Photo
		err := db.Debug().Where("id = ?", productId).Take(&photo).Error
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  "Photo not found!",
				"msg":    err.Error(),
			})
			return
		}

		if photo.UserId != id {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  "FORBIDDEN",
				"msg":    "You don't have right to access the resources",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		id := c.MustGet("id").(uint)
		productId := c.Param("productId")

		var photo models.Photo
		err := db.Debug().Where("id = ?", productId).Take(&photo).Error
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"error":  "Photo not found!",
				"msg":    err.Error(),
			})
			return
		}

		if photo.UserId != id {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status": http.StatusForbidden,
				"error":  "FORBIDDEN",
				"msg":    "You don't have right to access the resources",
			})
			return
		}

		c.Next()
	}
}
