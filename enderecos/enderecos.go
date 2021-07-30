package enderecos

import (
	"strings"
)

//TiposDeEnderecos é uma função que verifica se o tipo do endereço (rua) é válido
func TiposDeEnderecos(endereco string) string{
	tiposValidos := []string{"rua", "avenida", "rodovia", "estrada"}
	enderecoToLowerCase := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoToLowerCase, " ")[0]
	tpEndereco := false

	for _, tipo := range tiposValidos{
		if tipo == primeiraPalavraDoEndereco {
			tpEndereco = true
		}
	}

	if tpEndereco{
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo de Endereço inválido!"
}