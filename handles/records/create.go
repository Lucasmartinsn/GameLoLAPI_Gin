package handlers

import (
	"fmt"

	models "github.com/Lucasmartinsn/new-api-gin/models/records"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var records models.Record

	err := c.ShouldBindJSON(&records)
	if err != nil {
		c.JSON(400, gin.H{
			"Error":   true,
			"Mensage": "erro ao fazer decode do json %v" + err.Error(),
		})
		return
	}

	id, err := models.Insert(records)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":    true,
			"menssage": fmt.Sprintf("ocorreu um erro ao tenta inserir: %v ", err),
		}
	} else {
		resp = map[string]any{
			"Error":    false,
			"menssage": fmt.Sprintf("cadastrado com sucesso! Id : %v ", id),
		}
	}

	c.JSON(200, resp)
}
