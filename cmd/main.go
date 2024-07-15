package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/stxreocoma/todo-list/database"
	"github.com/stxreocoma/todo-list/handlers"
)

func main() {
	database.ConnectDB()

	r := chi.NewRouter()

	r.Get("/tasks", handlers.GetTasks)
	r.Post("/tasks", handlers.PostTask)
	r.Get("/tasks/{id}", handlers.GetTaskByID)
	r.Delete("/tasks/{id}", handlers.DeleteTask)

	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Printf("Start server error: %s", err.Error())
		return
	}
}
