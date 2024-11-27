package server

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) Server(port int) {
	http.HandleFunc("/h", s.HandlerFunction)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on port %d", port)

	log.Fatal(http.ListenAndServe(addr, nil))
}
