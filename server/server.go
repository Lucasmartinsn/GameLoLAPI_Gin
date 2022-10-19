package server

import (
	"log"

	"github.com/Lucasmartinsn/new-api-gin/server/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {

	router := router.ConfigRoutes(s.server)

	log.Print("server is running:", s.port)
	log.Fatal(router.Run(":" + s.port))
}
