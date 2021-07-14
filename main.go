package main

import (
	"task-httpserver/pkg/server"
)

func main() {
	serv := server.NewServer()
	serv.Run()
}