package main

import (
	"github.com/roberthameyer/loja-digport-backend/model"
)

// função que cria o catalogo de produtos do site
func catalogoProdutos() []model.Produto {
	produtos := []model.Produto{
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
	return produtos
}

// função para buscar um produto por nome
func produtosPorNome(nome string) []model.Produto {
	produtos := catalogoProdutos()
	catalogo := []model.Produto{}

	for i := range produtos {
		produtoAtual := produtos[i]
		if produtoAtual.Nome == nome {
			catalogo = append(catalogo, produtoAtual)
		}
	}
	return catalogo
}
