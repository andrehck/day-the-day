package main

import "testing"

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("Diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"

		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("diz 'Olá, mundo' quando uma string", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificaMensagemCorreta(t, resultado, esperado)

	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "Espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em frances", func(t *testing.T) {
		resultado := Ola("monde", "Francês")
		esperado := "Bonjour, monde"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
