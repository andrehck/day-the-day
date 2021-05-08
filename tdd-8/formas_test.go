package main

import (
	"math"
	"testing"
)

type Retangulo struct {
	Largura float64
	Altura  float64
}
type Circulo struct {
	Raio float64
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
	t.Run("retangulos", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		resultado := retangulo.Area()
		esperado := 72.0

		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}
	})

	t.Run("circulos", func(t *testing.T) {
		circulo := Circulo{10}
		resultado := circulo.Area()
		esperado := math.Pi * circulo.Raio * circulo.Raio

		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}

	})

}

//continuar na refatoração
