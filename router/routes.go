package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juancassiano/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine){
	v1 :=router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings", handler.DeleteOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}