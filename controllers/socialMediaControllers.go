package controllers

import (
	"go_final_project/helpers"
	"go_final_project/models"
	database "go_final_project/repository"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateSocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Socmed := models.Socialmedia{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.UserID = UserID

	err := db.Create(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Socmed)
}

func GetAllSocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var Socmed []models.Socialmedia
	UserID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", UserID).Find(&Socmed).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Socmed,
	})
}

func GetBySocmedID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Socmed := models.Socialmedia{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	err := db.Where("user_id = ? and id = ?", UserID, Socmed.Id).First(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Social Media Not Found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Socmed,
	})
}

func DeleteSocmedById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Socmed := models.Comment{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	err := db.Where("user_id = ? and id = ?", UserID, Socmed.Id).First(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Social Media Not Found",
		})
		return
	}

	db.Where("user_id = ? and id = ?", UserID, Socmed.Id).Delete(&Socmed)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Delete Data Success",
		"success": true,
	})
}

func UpdateSocmedByID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Socmed := models.Socialmedia{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	err := db.Where("user_id = ? and id = ?", UserID, Socmed.Id).First(&models.Socialmedia{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Comment Not Found",
		})
		return
	}

	err = db.Where("user_id = ? and id = ?", UserID, Socmed.Id).Updates(&Socmed).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Update Data Success",
		"success": true,
	})
}
