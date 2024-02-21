package routes

import (
	"github.com/BottoCarlos/MusicApp/cmd/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *handlers.Handler) {
	discoGroup := router.Group("/disco")
	{
		discoGroup.GET("/", handler.GetAll())
		discoGroup.GET("/:id", handler.GetByID())
		discoGroup.POST("/", handler.Add())
		discoGroup.PUT("/:id", handler.Update())
		discoGroup.DELETE("/:id", handler.Delete())
	}
}
