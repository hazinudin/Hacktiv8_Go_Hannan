package controllers

import (
	"go_final_project/helpers"
	"go_final_project/models"
	database "go_final_project/repository"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = UserID

	err := db.Where("user_id = ? and id = ?", UserID, Comment.PhotoID).Find(&models.Photo{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Photo Not Found",
		})
		return
	}

	err = db.Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	var Comments []models.Comment
	UserID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", UserID).Find(&Comments).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Comments,
	})
}

func GetByCommentID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Where("user_id = ? and id = ?", UserID, Comment.Id).First(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Comment Not Found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": Comment,
	})
}

func DeleteCommentById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Where("user_id = ? and id = ?", UserID, Comment.Id).First(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Comment Not Found",
		})
		return
	}

	db.Where("user_id = ? and id = ?", UserID, Comment.Id).Delete(&Comment)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Delete Data Success",
		"success": true,
	})
}

func UpdateCommentByID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Where("user_id = ? and id = ?", UserID, Comment.Id).First(&models.Comment{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": "Comment Not Found",
		})
		return
	}

	err = db.Where("user_id = ? and id = ?", UserID, Comment.Id).Updates(&Comment).Error

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Update Data Success",
		"success": true,
	})
}
