package enderecos_test

import (
	. "teste-automatizado/enderecos"
	"testing"
)

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTiposDeEnderecos(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioTeste{
		{"Rua Abc", "Rua"},
		{"Avenida das Hortas", "Avenida"},
		{"Rodovia dos Bandeirantes", "Rodovia"},
		{"Estrada do Matagal", "Estrada"},
		{"AVENIDA ESTREITA", "Avenida"},
		{"", "Tipo de Endereço inválido!"},
		{"Praça dos Bandeirantes", "Tipo de Endereço inválido!"},
		{"rua do jaumzinhu", "Rua"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TiposDeEnderecos(cenario.enderecoInserido)
		if cenario.retornoEsperado != retornoRecebido {
			t.Errorf("O tipo de Endereço retornado está inválido! Esperava receber %s recebeu %s",
				cenario.retornoEsperado, retornoRecebido)
		}
	}
}

func TestAleatorio(t *testing.T) {
	t.Parallel() //chama os testes em paralelo

	if 1 > 2 {
		t.Error("Quebrou o teste.")
	}
}
