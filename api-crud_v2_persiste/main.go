package main

import (
	"api_cru_pestistencia/data"
	"api_cru_pestistencia/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	data.Connect() // Estabelece conexão com o banco

	r := mux.NewRouter()
	r.HandleFunc("/produtos", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/produtos/{id}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/produtos", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/produtos/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/produtos/{id}", handlers.DeleteProduct).Methods("DELETE")

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
