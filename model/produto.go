package model

type Produto struct {
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Categoria  string  `json:"categoria"`
	Id         string  `json:"id"`
	Valor      float64 `json:"valor"`
	Quantidade int     `json:"quantidade"`
	Imagem     string  `json:"imagem"`
}
