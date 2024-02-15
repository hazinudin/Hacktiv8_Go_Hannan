package controllers

import (
	"go_final_project/helpers"
	"go_final_project/models"
	database "go_final_project/repository"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = UserID

	err := db.Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var Photos []models.Photo
	UserID := uint(userData["id"].(float64))

	err := db.Preload("Comments").Where("user_id = ?", UserID).Find(&Photos).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Photos,
	})
}

func GetByPhotoTitle(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Preload("Comments").Where("user_id = ? and title = ?", UserID, Photo.Title).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Photo Not Found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Photo,
	})
}

func DeletePhotoByTitle(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Where("user_id = ? and title = ?", UserID, Photo.Title).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Photo Not Found",
		})
		return
	}

	db.Where("user_id = ? and photo_id = ?", UserID, Photo.Id).Delete(&models.Comment{})
	db.Where("user_id = ? and title = ?", UserID, Photo.Title).Delete(&Photo)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Delete Data Success",
		"success": true,
	})
}

func UpdatePhotoByTitle(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Where("user_id = ? and title = ?", UserID, Photo.Title).First(&models.Photo{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Photo Not Found",
		})
		return
	}

	err = db.Where("user_id = ? and title = ?", UserID, Photo.Title).Updates(&Photo).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Update Data Success",
		"success": true,
	})
}
