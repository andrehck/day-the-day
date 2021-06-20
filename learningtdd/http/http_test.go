package main

import (
	"testing"
)

func Testhandler(t *testing.T) {
	resultado := handler("")
	esperado := "Hei, guys!"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
