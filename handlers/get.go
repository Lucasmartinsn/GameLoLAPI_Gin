package handlers

import (
	"strconv"

	"github.com/Lucasmartinsn/new-api-gin/models"
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

	register, err := models.GetOne(int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "geristro nao existe",
		})
		return
	}
	c.JSON(200, register)
}
