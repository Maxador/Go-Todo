package main

import (
	"net/http"
	"github.com/Maxador/Go-Todo/server"
	"fmt"
)

func main() {
	server.RegisterHandlers()
	http.Handle("/", http.FileServer(http.Dir("static")))
	fmt.Printf("Listening on localhost:8080 \n")
	http.ListenAndServe(":8080", nil)
}