package model

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
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

	produtos := []Produto{}

	for resultado.Next() {

		err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		var produto = populaProduto()

		produtos = append(produtos, produto)
	}
	defer db.Close()
	return produtos
}

func CriaProduto(prod Produto) error {
	//nome, descricao string, preco float64, image string, quantidade int

	if produtoCadastrado(prod.Nome) { // validação para conferir se o produto já existe
		fmt.Printf("Produto já cadastrado: %s\n", prod.Nome)
		return fmt.Errorf("Produto já cadastrado")
	}

	db := db.ConectaBancoDados()
	id := uuid.NewString() // evita que o número seja serial e ultrapasse a capacidade do banco
	nome := prod.Nome
	preco := prod.Preco
	descricao := prod.Descricao
	imagem := prod.Imagem
	quantidade := prod.QuantidadeEmEstoque

	strInsert := "INSERT INTO produtos VALUES($1, $2, $3, $4, $5, $6)" // envia o comando para o banco para inserir o produto

	result, err := db.Exec(strInsert, id, nome, strconv.FormatFloat(preco, 'f', 1, 64), descricao, imagem, strconv.Itoa(quantidade)) // passa os parametros do produto cadastrado para o banco
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Produto %s criado com sucesso (%d row affected)\n", id, rowsAffected)
	defer db.Close()
	return nil
}

func produtoCadastrado(nomeProduto string) bool {
	prod := BuscaProdutoPorNome(nomeProduto)
	return prod.Nome == nomeProduto
}

func RemoveProduto(id string) error {
	db := db.ConectaBancoDados()

	resultado, err := db.Exec("DELETE FROM PRODUTOS WHERE id = $1", id)
	if err != nil {
		fmt.Printf("Ocorreu um erro ao tentar excluir produto: %s", err.Error())
		return fmt.Errorf("Ocorreu um erro ao tentar excluir o produto: %w", err)
	}

	linesAffected, err := resultado.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("%d linhas afetadas\n", linesAffected)
	fmt.Println("Produto excluído")

	defer db.Close()
	return nil
}

func BuscaProdutoPorNome(nomeProduto string) Produto {
	db := db.ConectaBancoDados()
	res := db.QueryRow("SELECT * FROM produtos where nome = $1", nomeProduto)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	if err == sql.ErrNoRows {
		fmt.Printf("Produto nao encontrado %s\n", nomeProduto)
	} else if err != nil {
		panic(err.Error())
	}

	var produto1 = populaProduto()

	defer db.Close() //sempre colocar para não ficar conexao aberta
	return produto1
}

func populaProduto() Produto {
	var produto1 Produto
	produto1.ID = id
	produto1.Nome = nome
	produto1.Descricao = descricao
	produto1.Preco = preco
	produto1.Imagem = imagem
	produto1.QuantidadeEmEstoque = quantidade
	return produto1

}
