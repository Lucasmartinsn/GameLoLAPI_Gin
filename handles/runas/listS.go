package runas

import (
	"log"

	models "github.com/Lucasmartinsn/new-api-gin/models/runas"
	"github.com/gin-gonic/gin"
)

func Getrunes(c *gin.Context) {
	resp, err := models.Getrunes()
	if err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(200, resp)
	}
}
