package main

import (
	"testing"
)

type Retangulo struct {
	Largura float64
	Altura  float64
}
type Circulo struct {
	Raio float64
}

type Forma interface {
	Area() float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {
	testesArea := []struct {
		forma    Forma
		esperado float64
	}{
		{forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72.0},
		{forma: Circulo{Raio: 10}, esperado: 314.1592653589793},
		{forma: Triangulo{Base: 12, Altura: 6}, esperado: 36.0},
	}

	for _, tt := range testesArea {
		resultado := tt.forma.Area()
		if resultado != tt.esperado {
			t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.esperado)
		}
	}

}

//continuar na refatoração
