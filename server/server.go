package server

import {
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Maxador/Go-Todo/task"
	"github.com/gorilla/mux"
}

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