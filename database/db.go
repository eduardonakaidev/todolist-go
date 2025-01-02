package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Inicializa a conexão com o banco de dados
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	log.Println("Banco de dados conectado com sucesso!")
}
