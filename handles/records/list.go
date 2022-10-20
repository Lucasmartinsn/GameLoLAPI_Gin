package handlers

import (
	"log"

	models "github.com/Lucasmartinsn/new-api-gin/models/records"
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
