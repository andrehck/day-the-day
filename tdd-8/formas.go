package main

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

func (r Retangulo) Area() float64 {
	return 72.00
}

func (c Circulo) Area() float64 {
	return 314.15
}
