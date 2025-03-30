package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	var req Task
	json.NewDecoder(r.Body).Decode(&req)
	DB.Create(&req)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)

}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updates map[string]interface{}
	json.NewDecoder(r.Body).Decode(&updates)

	DB.Model(&Task{}).Where("id = ?", id).Updates(updates)

	var task Task
	DB.First(&task, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	DB.Delete(&Task{}, id)

	w.WriteHeader(http.StatusNoContent)
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
	router.HandleFunc("/api/task/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", TaskHandler).Methods("POST")
	router.HandleFunc("/api/task/delete/{id}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
