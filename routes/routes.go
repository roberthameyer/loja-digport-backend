package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roberthameyer/loja-digport-backend/controllers"
	"github.com/rs/cors"
)

func HandleRequests() {
	route := mux.NewRouter()
	route.HandleFunc("/produtos", controllers.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produto", controllers.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produto", controllers.CriaProdutosHandler).Methods("POST")
	route.HandleFunc("/produto/{id}", controllers.RemoveProdutoHandler).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(route)
	http.ListenAndServe(":8080", handler)

}
