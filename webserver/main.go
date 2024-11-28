package main

import "github.com/hailsayan/miniProjects/webserver/server"

func main() {
	server := server.New()
	server.Serve()
}
