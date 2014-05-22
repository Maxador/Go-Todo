package server

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	// "strconv"

	"github.com/Maxador/Go-Todo/task"
	"github.com/gorilla/mux"
)


var tasks = task.NewTaskManager()

const PathPrefix = "/task/"

func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix, errorHandler(ListTasks)).Methods("GET")
}

func ListTasks(w http.ResponseWriter, r *http.Request) error {
	res := struct { Tasks []*task.Task }{tasks.All()}
	return json.NewEncoder(w).Encode(res)
}

func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		// case badRequest:
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// case notFound:
		// 	http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}