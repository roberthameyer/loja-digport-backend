package main

import "fmt"

func main() {
	StartServer()
	catalogoFiltro := produtosPorNome("Pacote Hungria")
	fmt.Println(catalogoFiltro)
}
