package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Maxador/Go-Todo/task"
	"github.com/gorilla/mux"
)


var tasks = task.NewTaskManager()

const PathPrefix = "/task/"

func RegisterHandlers() {
	r := mux.NewRouter()
	r.HandleFunc(PathPrefix, errorHandler(ListTasks)).Methods("GET")
	r.HandleFunc(PathPrefix, errorHandler(NewTask)).Methods("POST")
	r.HandleFunc(PathPrefix + "{id}", errorHandler(GetTask)).Methods("GET")
	r.HandleFunc(PathPrefix + "{id}", errorHandler(UpdateTask)).Methods("PUT")
	http.Handle(PathPrefix, r)
}

type badRequest struct{ error }

type notFound struct{ error }

func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

func parseID(r *http.Request) (int64, error) {
	str, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, fmt.Errorf("Task id not found")
	}
	return strconv.ParseInt(str, 10, 0)
}

func ListTasks(w http.ResponseWriter, r *http.Request) error {
	res := struct { Tasks []*task.Task }{tasks.All()}
	return json.NewEncoder(w).Encode(res)
}

func NewTask(w http.ResponseWriter, r *http.Request) error {
	req := struct{ Title string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return badRequest{err}
	}
	t, err := task.NewTask(req.Title)
	if err != nil {
		return badRequest{err}
	}
	return tasks.Save(t)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetTask(w http.ResponseWriter, r *http.Request) error {
	id, err:= parseID(r)
	if err != nil {
		return badRequest{err}
	}
	t, ok := tasks.Find(id)
	if !ok {
		return fmt.Errorf("Task not found")
	}
	return json.NewEncoder(w).Encode(t)
}