package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/juancassiano/gopportunities/docs"
	"github.com/juancassiano/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize Handler
	handler.InitializeHandler()
	BasePath := "/api/v1"
	docs.SwaggerInfo.BasePath = BasePath
	v1 := router.Group(BasePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings", handler.DeleteOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	//Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
