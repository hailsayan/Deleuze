package server

import (
	"fmt"
	"net/http"
)

func (handler *server) HandlerFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working!")
}
