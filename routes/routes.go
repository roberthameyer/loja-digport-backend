package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roberthameyer/loja-digport-backend/controllers"
)

func HandleRequests() {
	route := mux.NewRouter()

	route.HandleFunc("/produtos", controllers.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produto", controllers.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produtos", controllers.CriaProdutosHandler).Methods("POST")

	http.ListenAndServe(":8080", route)
}
