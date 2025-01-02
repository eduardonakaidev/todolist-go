package models

import "github.com/eduardonakaidev/todo-list-go/database"

// Realiza as migrações
func Migrate() {
	database.DB.AutoMigrate(&Todo{})
}
