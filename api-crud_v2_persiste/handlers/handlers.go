package handlers

import (
	"api_cru_pestistencia/data"
	"api_cru_pestistencia/models"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Listar todos os produtos
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := data.DB.Query("SELECT id, nome, preco, quantidade FROM produtos")
	if err != nil {
		http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []models.Produto
	for rows.Next() {
		var p models.Produto
		if err := rows.Scan(&p.ID, &p.Nome, &p.Preco, &p.Quantidade); err != nil {
			http.Error(w, "Erro ao ler dados", http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, p)
	}

	if err := json.NewEncoder(w).Encode(produtos); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
	}
}

// Obter um produto por ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var p models.Produto
	err := data.DB.QueryRow("SELECT id, nome, preco, quantidade FROM produtos WHERE id = $1", id).
		Scan(&p.ID, &p.Nome, &p.Preco, &p.Quantidade)

	if err == sql.ErrNoRows {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Erro ao buscar produto", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
	}
}

// Criar um novo produto
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p models.Produto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err := data.DB.QueryRow(
		"INSERT INTO produtos (nome, preco, quantidade) VALUES ($1, $2, $3) RETURNING id",
		p.Nome, p.Preco, p.Quantidade,
	).Scan(&p.ID)

	if err != nil {
		http.Error(w, "Erro ao inserir produto", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
	}
}

// Atualizar um produto
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var p models.Produto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	_, err := data.DB.Exec(
		"UPDATE produtos SET nome = $1, preco = $2, quantidade = $3 WHERE id = $4",
		p.Nome, p.Preco, p.Quantidade, id,
	)

	if err != nil {
		http.Error(w, "Erro ao atualizar produto", http.StatusInternalServerError)
		return
	}

	p.ID = id
	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
	}
}

// Deletar um produto
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := data.DB.Exec("DELETE FROM produtos WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Erro ao deletar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Função auxiliar para obter o ID da requisição
func getIDFromRequest(r *http.Request) (int, error) {
	params := mux.Vars(r)
	return strconv.Atoi(params["id"])
}

//// Função auxiliar para buscar um produto por ID
//func findProductByID(id int) (models.Produto, bool) {
//	for _, produto := range data.Produtos {
//		if produto.ID == id {
//			return produto, true
//		}
//	}
//	return models.Produto{}, false
//}

//// Gerar o próximo ID (simples incremento)
//func getNextID() int {
//	if len(data.Produtos) == 0 {
//		return 1
//	}
//	return data.Produtos[len(data.Produtos)-1].ID + 1
//}
