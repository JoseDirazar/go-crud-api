package main

import (
	"net/http"

	"github.com/JoseDirazar/go-rest-api/db"
	"github.com/JoseDirazar/go-rest-api/models"
	"github.com/JoseDirazar/go-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUsersHandler).Methods("POST") 
	router.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTasksHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}