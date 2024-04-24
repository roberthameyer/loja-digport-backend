package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)
	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {

	nomeQuery := r.URL.Query().Get("nome") // adiciona um query parameter nome no endpoint /produtos

	if nomeQuery != "" {
		produtosFiltrados := produtosPorNome(nomeQuery)
		json.NewEncoder(w).Encode(produtosFiltrados)
	} else {
		produtos := catalogoProdutos()
		json.NewEncoder(w).Encode(produtos)
	}
}
