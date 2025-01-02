package routes

import (
	"github.com/eduardonakaidev/todo-list-go/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/todos", controllers.AddTodo).Methods("POST")
	router.HandleFunc("/todos/{id}/complete", controllers.CompleteTodo).Methods("PUT")

	return router
}
