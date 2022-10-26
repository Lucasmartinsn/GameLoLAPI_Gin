package router

import (
	info_game "github.com/Lucasmartinsn/new-api-gin/handles/info_game"
	info_user "github.com/Lucasmartinsn/new-api-gin/handles/records"
	info_runes "github.com/Lucasmartinsn/new-api-gin/handles/runas"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("user")
	{
		lol := main.Group("")
		{
			lol.GET("/", info_user.Get)
			lol.GET("/:id", info_user.GetOne)
			lol.POST("/", info_user.Create)
			lol.PUT("/:id", info_user.Update)
			lol.DELETE("/:id", info_user.Delete)
			lol.PUT("/:id/foto", info_user.Upfoto)
			lol.PUT("/:id/name", info_user.Upname)
			lol.GET("/info/:id", info_game.Getinfo)
			lol.PUT("/info/:id", info_game.Updateinfo)
			lol.GET("/runes/principais", info_runes.GetrunaP)
			lol.GET("/runes/secundarias", info_runes.Getrunes)
			lol.PUT("/runes/secundarias/:id", info_runes.Upimgrunas)
		}
	}

	return router
}
