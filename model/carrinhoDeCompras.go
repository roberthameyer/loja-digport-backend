package model

type CarrinhoDeCompras struct {
	ID           string
	IdUsuario    string
	InfosProduto []InfosProduto
	ValorTotal   float64
}

type InfosProduto struct {
	IdProduto           string
	QuantidadeDeProduto int
}
