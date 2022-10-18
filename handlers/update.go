package handlers

import (
	"strconv"

	"github.com/Lucasmartinsn/new-api-gin/models"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "id tem de ser um int",
		})
		return
	}

	var records models.Record

	rows, err := models.Update(int64(id), records)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Ao atualizar o registro",
		})

		return
	}
	if rows > 1 {
		c.JSON(400, gin.H{
			"Error": "Foi selecionado mais de um registro",
		})
		return
	}

	resp := map[string]any{
		"Error":   false,
		"mensage": "dados atualizados com sucesso",
	}

	c.JSON(200, resp)
}
