package infogame

import (
	"strconv"

	models "github.com/Lucasmartinsn/new-api-gin/models/info_game"
	"github.com/gin-gonic/gin"
)

func Getinfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "id tem de ser um int",
		})
		return
	}

	register, err := models.Getinfo(int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "registro nao existe",
		})
		return
	}
	c.JSON(200, register)
}
