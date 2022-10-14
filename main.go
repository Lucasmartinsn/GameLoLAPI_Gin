package main

import "github.com/Lucasmartinsn/new-api-gin/server"

func main() {
	server := server.NewServer()

	server.Run()
}
