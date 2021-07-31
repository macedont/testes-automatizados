package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		r := Retangulo{10, 10}
		areaEsperada := float64(100)
		areaRecebida := r.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada) //mata a aplicação aq diferentemente do Errorf
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		c := Circulo{10}
		areaEsperada := math.Pi * 100
		areaRecebida := c.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada) //mata a aplicação aq diferentemente do Errorf
		}

	})
}
