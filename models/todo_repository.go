package models

import "github.com/eduardonakaidev/todo-list-go/database"

// Retorna todas as tarefas
func GetAllTodos() []Todo {
	var todos []Todo
	database.DB.Find(&todos)
	return todos
}

// Adiciona uma nova tarefa
func AddTodo(todo Todo) {
	database.DB.Create(&todo)
}

// Marca uma tarefa como concluída
func CompleteTodo(id uint) bool {
	var todo Todo
	result := database.DB.First(&todo, id)
	if result.Error == nil {
		todo.Completed = true
		database.DB.Save(&todo)
		return true
	}
	return false
}