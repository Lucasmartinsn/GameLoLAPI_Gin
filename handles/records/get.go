package handlers

import (
	"strconv"

	models "github.com/Lucasmartinsn/new-api-gin/models/records"
	"github.com/gin-gonic/gin"
)

func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "id tem de ser um int",
		})
		return
	}

	register, err := models.Get(int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "registro nao existe",
		})
		return
	}
	c.JSON(200, register)
}
