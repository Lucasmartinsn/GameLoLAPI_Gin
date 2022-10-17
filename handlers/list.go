package handlers

import (
	"log"

	"github.com/Lucasmartinsn/new-api-gin/models"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	resp, err := models.GetAll()

	if err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(200, resp)
	}

}
