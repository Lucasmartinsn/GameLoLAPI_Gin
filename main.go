package main

import (
	"github.com/Lucasmartinsn/new-api-gin/configs"
	"github.com/Lucasmartinsn/new-api-gin/server"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	server := server.NewServer()

	server.Run()
}
