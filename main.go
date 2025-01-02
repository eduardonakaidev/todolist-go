package main

import (
	"log"
	"net/http"

	"github.com/eduardonakaidev/todo-list-go/database"
	"github.com/eduardonakaidev/todo-list-go/models"
	"github.com/eduardonakaidev/todo-list-go/routes"
)

func main() {
	// Inicializa o banco de dados
	database.InitDatabase()

	// Executa as migrações
	models.Migrate()   

	// Configura as rotas
	router := routes.SetupRoutes()

	// Inicia o servidor...
	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
