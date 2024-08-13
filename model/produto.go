package model

import (
	"database/sql"
	"fmt"

	"github.com/roberthameyer/loja-digport-backend/db"
)

type Produto struct {
	Nome                string
	Descricao           string
	Categoria           string
	ID                  string
	Preco               float64
	Quantidade          int
	Imagem              string
	QuantidadeEmEstoque int
}

var id, nome string
var preco float64
var descricao, imagem string
var quantidade int

func BuscaTodosProdutos() []Produto {
	db := db.ConectaBancoDados()

	resultado, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
		// return nil, errors.New("erro ao buscar produtos")
	}

	p := Produto{}
	produtos := []Produto{}

	for resultado.Next() {

		err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.ID = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Imagem = imagem
		p.QuantidadeEmEstoque = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func BuscaProdutoPorNome(nomeProduto string) Produto {
	db := db.ConectaBancoDados()

	res := db.QueryRow("SELECT * FROM produtos where nome = $1", nomeProduto)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	if err == sql.ErrNoRows {
		fmt.Printf("Produto nao encontrado %s\n", nome)

	} else if err != nil {
		panic(err.Error())
	}

	var p Produto
	p.ID = id
	p.Nome = nomeProduto
	p.Descricao = descricao
	p.Preco = preco
	p.Imagem = imagem
	p.QuantidadeEmEstoque = quantidade

	defer db.Close() //sempre colocar para não ficar conexao aberta
	return p

}
