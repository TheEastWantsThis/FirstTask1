package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type RequestBody struct {
	Task string `json:"task"`
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestBody
	json.NewDecoder(r.Body).Decode(&req)
	task = req.Task
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, ", task)
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")
	http.ListenAndServe(":8000", router)
}
