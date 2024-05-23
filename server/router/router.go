package router

import (
	"github.com/Maysa87/Goland_to_do/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/DeleteTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/DeleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
