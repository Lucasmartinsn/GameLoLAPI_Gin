package runas

import (
	"log"

	models "github.com/Lucasmartinsn/new-api-gin/models/runas"
	"github.com/gin-gonic/gin"
)

func GetrunaP(c *gin.Context) {
	resp, err := models.GetrRuna()
	if err != nil {
		log.Fatalln(err)
	} else {
		c.JSON(200, resp)
	}
}
