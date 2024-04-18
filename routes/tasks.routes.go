package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JoseDirazar/go-rest-api/db"
	"github.com/JoseDirazar/go-rest-api/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler( w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	db.DB.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Task not found."))
	return
	}

	json.NewEncoder(w).Encode(&task)
}

func CreateTasksHandler( w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)

	err := createdTask.Error
	if createdTask.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func DeleteTasksHandler( w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Task not found"))
		return
	}

	db.DB.Delete(&task)
	w.WriteHeader(http.StatusOK)
}