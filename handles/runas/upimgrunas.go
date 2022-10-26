package runas

import (
	"strconv"

	models "github.com/Lucasmartinsn/new-api-gin/models/runas"
	"github.com/gin-gonic/gin"
)

func Upimgrunas(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao fazer parse do id %v" + err.Error(),
		})
		return
	}

	var runas models.UpfotorunaS

	err = c.ShouldBindJSON(&runas)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao fazer decode do json: %v" + err.Error(),
		})
		return
	}

	rows, err := models.UpdateRuna(int64(id), runas)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao atualizar register: %v" + err.Error(),
		})
		return
	}

	if rows > 1 {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro: foram atualizador %d registros",
		})
	}

	resp := map[string]any{
		"Error":   false,
		"mensage": "dados atualizados com sucesso",
	}

	c.JSON(200, resp)

}
