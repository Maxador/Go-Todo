package main

import (
	"net/http"
	"github.com/Maxador/Go-Todo/server"
)

func main() {
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("server/static")))
	http.ListenAndServe(":8080", nil)
}