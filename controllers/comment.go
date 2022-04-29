package controllers

import (
	"fmt"
	"net/http"
	"sesi-final-project/database"
	"sesi-final-project/helpers"
	"sesi-final-project/models"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	id := c.MustGet("id")

	comment := models.Comment{}
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	comment.UserId = id.(uint)

	err := db.Debug().Create(&comment).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  201,
		"payload": comment,
	})
}

func GetComments(c *gin.Context) {
	db := database.GetDB()
	var (
		comments []models.Comment
		result   gin.H
	)

	db.Find(&comments)

	if len(comments) == 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		var comment_slice []models.Comment

		for _, p := range comments {

			comments = append(comment_slice, p)
		}

		result = gin.H{
			"result": comment_slice,
			"count":  len(comment_slice),
		}
	}

	c.JSON(http.StatusOK, result)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	commentId := c.Param("commentId")

	var comment models.Comment
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&comment)
	} else {
		c.ShouldBind(&comment)
	}

	err := db.Debug().Model(&comment).Where("id = ?", commentId).Updates(models.Comment{
		Message: comment.Message,
	}).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"payload": comment,
	})

}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentId := c.Param("commentId")

	var user models.User
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	var comment models.Comment
	err := db.Debug().Model(&comment).Where("id = ?", commentId).Delete(models.Comment{}).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been succesfully deleted",
	})

}
