package router

import (
	"go_final_project/controllers"
	"go_final_project/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/getall", controllers.GetAllPhoto)
		photoRouter.GET("/getone", controllers.GetByPhotoTitle)
		photoRouter.DELETE("/delete", controllers.DeletePhotoByTitle)
		photoRouter.PUT("/update", controllers.UpdatePhotoByTitle)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/getall", controllers.GetAllComment)
		commentRouter.GET("/getone", controllers.GetByCommentID)
		commentRouter.DELETE("/delete", controllers.DeleteCommentById)
		commentRouter.PUT("/update", controllers.UpdateCommentByID)
	}

	socmedRouter := router.Group("/socmed")
	{
		socmedRouter.Use(middlewares.Authentication())
		socmedRouter.POST("/", controllers.CreateSocmed)
		socmedRouter.GET("/getall", controllers.GetAllSocmed)
		socmedRouter.GET("/getone", controllers.GetBySocmedID)
		socmedRouter.DELETE("/delete", controllers.DeleteSocmedById)
		socmedRouter.PUT("/update", controllers.UpdateSocmedByID)
	}

	return router
}
