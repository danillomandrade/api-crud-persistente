package data

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() {
	var err error

	//configure a string de conexão com postgres
	//postgres é o user
	//1234 é a senha
	//apicrud é o nome do banco
	connStr := "postgres://postgres:1234@localhost:5432/apicrud?sslmode=disable"

	DB, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Testa a conexão
	if err := DB.Ping(); err != nil {
		log.Fatalf("Não foi possível estabelecer conexão com o banco: %v", err)
	}

	log.Println("Conectado ao banco de dados com sucesso!")

}
