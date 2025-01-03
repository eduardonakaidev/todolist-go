package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eduardonakaidev/todo-list-go/models"
	"github.com/eduardonakaidev/todo-list-go/views"
	"github.com/gorilla/mux"
)

// Handler para listar todas as tarefas
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := models.GetAllTodos()
	views.RenderJSON(w, todos, http.StatusOK) 
}

// Handler para adicionar uma nova tarefa
func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		views.RenderError(w, views.NewAPIError("Dados inválidos", http.StatusBadRequest))
		return
	}

	models.AddTodo(todo) 
	views.RenderJSON(w, todo, http.StatusCreated) 
}

// Handler para marcar uma tarefa como concluída
func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		views.RenderError(w, views.NewAPIError("ID inválido", http.StatusBadRequest))
		return
	}
	if models.CompleteTodo(uint(id)) { 
		views.RenderJSON(w, map[string]string{"message": "Tarefa concluída com sucesso"}, http.StatusOK) 
	} else {
		views.RenderError(w, views.NewAPIError("Tarefa não encontrada", http.StatusNotFound))
	}
}
