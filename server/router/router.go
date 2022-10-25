package router

import (
	handlersa "github.com/Lucasmartinsn/new-api-gin/handles/info_game"
	handlers "github.com/Lucasmartinsn/new-api-gin/handles/records"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("user")
	{
		lol := main.Group("")
		{
			lol.GET("/", handlers.Get)
			lol.GET("/:id", handlers.GetOne)
			lol.POST("/", handlers.Create)
			lol.PUT("/:id", handlers.Update)
			lol.DELETE("/:id", handlers.Delete)
			lol.PUT("/:id/foto", handlers.Upfoto)
			lol.PUT("/:id/name", handlers.Upname)
			lol.GET("/info/:id", handlersa.Getinfo)
			lol.PUT("/info/:id", handlersa.Updateinfo)
		}
	}

	return router
}
