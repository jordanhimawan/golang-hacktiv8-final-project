package controllers

import (
	"fmt"
	"net/http"
	"sesi-final-project/database"
	"sesi-final-project/helpers"
	"sesi-final-project/models"

	"github.com/gin-gonic/gin"
)

var AppJson = "application/json"

func UserRegister(c *gin.Context) {
	db := database.GetDB()

	user := models.User{}
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error
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
		"payload": user,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	user := models.User{}
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	password := user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": err.Error(),
		})
		return
	}

	if err := helpers.ComparePass(user.Password, password); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "UNAUTHORIZED",
			"message": err.Error(),
		})
		return
	}

	token := helpers.GenerateToken(user.Id, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"status":       200,
		"access_token": token,
	})
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	userId := c.Param("userId")

	var user models.User
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Model(&user).Where("id = ?", userId).Updates(models.User{
		Email:    user.Email,
		Password: user.Password,
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
		"payload": user,
	})

}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	userId := c.Param("userId")

	var user models.User
	if AppJson == helpers.GetContentType(c) {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Model(&user).Where("id = ?", userId).Delete(models.User{
		Email:    user.Email,
		Password: user.Password,
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
		"message": "Your account has been succesfully deleted",
	})

}
