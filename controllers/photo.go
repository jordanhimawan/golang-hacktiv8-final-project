package controllers

import (
	"fmt"
	"net/http"
	"sesi-final-project/database"
	"sesi-final-project/helpers"
	"sesi-final-project/models"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	id := c.MustGet("id")

	photo := models.Photo{}
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	photo.UserId = id.(uint)

	err := db.Debug().Create(&photo).Error
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
		"payload": photo,
	})
}

func GetPhotos(c *gin.Context) {
	db := database.GetDB()
	var (
		photos []models.Photo
		result gin.H
	)

	db.Find(&photos)

	if len(photos) == 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		var photo_slice []models.Photo

		for _, p := range photos {

			photos = append(photo_slice, p)
		}

		result = gin.H{
			"result": photo_slice,
			"count":  len(photo_slice),
		}
	}

	c.JSON(http.StatusOK, result)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	photoId := c.Param("photoId")

	var photo models.Photo
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&photo)
	} else {
		c.ShouldBind(&photo)
	}

	err := db.Debug().Model(&photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
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
		"payload": photo,
	})

}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	photoId := c.Param("photoId")

	var user models.User
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	var photo models.Photo
	err := db.Debug().Model(&photo).Where("id = ?", photoId).Delete(models.Photo{}).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been succesfully deleted",
	})

}
