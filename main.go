package main

import (
	"fmt"
	"teste-automatizado/enderecos"
)

func main(){
	resultado := enderecos.TiposDeEnderecos("Avenida João Batista Loureiro")

	fmt.Println(resultado)
}