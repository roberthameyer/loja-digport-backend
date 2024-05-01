package main

import (
	"encoding/json"
	"net/http"

	"github.com/roberthameyer/loja-digport-backend/model"
)

// função que inicia o server e cria o path de produtos
func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)
	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		buscaProduto(w, r)
	} else if r.Method == "POST" {
		addProduto(w, r)
	}
}

func addProduto(w http.ResponseWriter, r *http.Request) {
	var produtoBuscado model.Produto
	json.NewDecoder(r.Body).Decode(&produtoBuscado)
	criaProduto(produtoBuscado)

	err := criaProduto(produtoBuscado)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func buscaProduto(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome") //adiciona um query parameter nome no endpoint /produtos
	if queryNome != "" {
		produtosFiltradosPorNome := produtosPorNome(queryNome)
		json.NewEncoder(w).Encode(produtosFiltradosPorNome)
	} else {
		produtosBuscados := produtos
		json.NewEncoder(w).Encode(produtosBuscados)
	}
}
