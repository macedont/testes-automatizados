package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	Altura float64
	Base   float64
}

type Circulo struct {
	Raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func (r Retangulo) area() float64 {
	return r.Base * r.Altura
}

func SetArea(f Forma) {
	fmt.Printf("O resultado é %.2f", f.area())
}
