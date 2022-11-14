package login

import (
	"strconv"

	models "github.com/Lucasmartinsn/new-api-gin/models/login"
	"github.com/gin-gonic/gin"
)

func Verificalogin(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "id tem de ser um int",
			"pq":    err.Error(),
		})
		return
	}

	register, err := models.Verificalogin(int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "geristro nao existe",
			"pq":    err.Error(),
		})
		return
	}
	c.JSON(200, register)
}
