package handlers

import (
	"strconv"

	"github.com/Lucasmartinsn/new-api-gin/models"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao fazer parse do id " + err.Error(),
		})
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao remover register: " + err.Error(),
		})
		return
	}

	if rows > 1 {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro: foram removidos registros",
		})
	}

	resp := map[string]any{
		"Error":   false,
		"mensage": "registro removido com sucesso",
	}

	c.JSON(200, resp)
}
