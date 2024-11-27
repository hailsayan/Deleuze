package main

import "github.com/hailsayan/miniProjects/webserver/server"

func main() {
	server := server.NewServer()
	server.Server(8080)
}
