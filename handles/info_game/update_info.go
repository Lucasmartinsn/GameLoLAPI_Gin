package infogame

import (
	"strconv"

	models "github.com/Lucasmartinsn/new-api-gin/models/info_game"
	"github.com/gin-gonic/gin"
)

func Updateinfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao fazer parse do id %v" + err.Error(),
		})
		return
	}

	var infogame models.Infogame

	err = c.ShouldBindJSON(&infogame)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao fazer decode do json: %v" + err.Error(),
		})
		return
	}

	rows, err := models.Updateinfo(int64(id), infogame)
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
