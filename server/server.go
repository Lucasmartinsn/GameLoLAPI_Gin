package server

import (
	"log"

	"github.com/gin-gonic/gin"

	//"github.com/Lucasmartinsn/new-api-gin/server/db"
	"github.com/Lucasmartinsn/new-api-gin/server/routes"
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
	//err := db.Conectar()
	//if err != nil {
	//	panic(err)
	///}
	router := routes.ConfigRoutes(s.server)

	log.Print("server is running:", s.port)
	log.Fatal(router.Run(":" + s.port))
}
