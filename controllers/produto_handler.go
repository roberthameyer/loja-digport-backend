package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/roberthameyer/loja-digport-backend/model"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosProdutos()
	json.NewEncoder(w).Encode(produtos)
}

func BuscaProdutoPorNomeHandler(w http.ResponseWriter, r *http.Request) {
	//o parametro será processado como parte da url
	//http://localhost:8080/produto?nome=Revista moda
	nome := r.URL.Query().Get("nome")
	produto := model.BuscaProdutoPorNome(nome)
	json.NewEncoder(w).Encode(produto)
}

func CriaProdutosHandler(w http.ResponseWriter, r *http.Request) {

}
