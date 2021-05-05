package main

import "math"

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
