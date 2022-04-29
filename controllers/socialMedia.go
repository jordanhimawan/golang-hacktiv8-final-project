package controllers

import (
	"fmt"
	"net/http"
	"sesi-final-project/database"
	"sesi-final-project/helpers"
	"sesi-final-project/models"

	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(c *gin.Context) {
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

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	var (
		socialMedias []models.SocialMedia
		result       gin.H
	)

	db.Find(&socialMedias)

	if len(socialMedias) == 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		var socialmedia_slice []models.SocialMedia

		for _, s := range socialMedias {
			socialMedias = append(socialmedia_slice, s)
		}

		result = gin.H{
			"result": socialmedia_slice,
			"count":  len(socialmedia_slice),
		}
	}

	c.JSON(http.StatusOK, result)
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMediaId := c.Param("socialMediaId")

	var socialMedia models.SocialMedia
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	err := db.Debug().Model(&socialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
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
		"payload": socialMedia,
	})

}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	socialMediaId := c.Param("socialMediaId")

	var user models.User
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	var socialMedia models.SocialMedia
	err := db.Debug().Model(&socialMedia).Where("id = ?", socialMediaId).Delete(models.SocialMedia{}).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been succesfully deleted",
	})

}
