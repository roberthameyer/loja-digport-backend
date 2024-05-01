package main

import (
	"errors"

	"github.com/roberthameyer/loja-digport-backend/model"
)

var produtos []model.Produto = []model.Produto{}

// função que cria o catalogo de produtos do site
func catalogoProdutos() {
	produtos = []model.Produto{
		{
			Nome:       "Pacote Espanha",
			Descricao:  "Pacote de viagem para Espanha",
			Categoria:  "Pacote",
			Id:         "1",
			Valor:      7000.00,
			Quantidade: 1,
			Imagem:     "Nome.png",
		},
		{
			Nome:       "Pacote Grécia",
			Descricao:  "Pacote de viagem para Grécia",
			Categoria:  "Pacote",
			Id:         "2",
			Valor:      8000.00,
			Quantidade: 1,
			Imagem:     "Nome1.png",
		},
		{
			Nome:       "Pacote Hungria",
			Descricao:  "Pacote de viagem para Hungria",
			Categoria:  "Pacote",
			Id:         "3",
			Valor:      9000.00,
			Quantidade: 1,
			Imagem:     "Nome2.png",
		},
	}

}

// função para buscar um produto por nome
func produtosPorNome(nome string) []model.Produto {
	resultado := []model.Produto{}

	for _, produtoBuscado := range produtos {
		if produtoBuscado.Nome == nome {
			resultado = append(resultado, produtoBuscado)
		}
	}

	return resultado
}

// função para criar um produto e retorna erro se o ID já for cadastrado
func criaProduto(produtoBuscado model.Produto) error {
	for _, idProduto := range produtos {
		if produtoBuscado.Id == idProduto.Id {
			return errors.New("pacote com ID já cadastrado")
		}
	}
	produtos = append(produtos, produtoBuscado)
	return nil
}
