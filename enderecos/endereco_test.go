package enderecos

import "testing"

func TestTiposDeEnderecos(t *testing.T) {
	enderecoParaTest := "Estrada dos bandeirantes"
	tipoDeEnderecoEsperado := "Estrada"
	tipoDeEnderecoRecebido := TiposDeEnderecos(enderecoParaTest)
	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
			tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}
