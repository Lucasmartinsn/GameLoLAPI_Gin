package router

import (
	handlers "github.com/Lucasmartinsn/new-api-gin/handles/records"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		lol := main.Group("lol")
		{
			lol.GET("/", handlers.Get)
			lol.GET("/:id", handlers.GetOne)
			lol.POST("/", handlers.Create)
			lol.PUT("/:id", handlers.Update)
			lol.DELETE("/:id", handlers.Delete)
		}
	}

	return router
}
