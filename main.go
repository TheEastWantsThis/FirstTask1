package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var req Task
	json.NewDecoder(r.Body).Decode(&req)
	DB.Create(&req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	DB.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	InitDB()
	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")
	http.ListenAndServe(":8000", router)
}
