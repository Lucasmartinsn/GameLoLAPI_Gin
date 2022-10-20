package router

import (
	handlers "github.com/Lucasmartinsn/new-api-gin/handles/records"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api")
	{
		lol := main.Group("")
		{
			lol.GET("/user", handlers.Get)
			lol.GET("/user/:id", handlers.GetOne)
			lol.POST("/user", handlers.Create)
			lol.PUT("/user/:id", handlers.Update)
			lol.DELETE("/user/:id", handlers.Delete)
			lol.PUT("/user/:id/foto", handlers.Upfoto)
		}
	}

	return router
}
