package interacao

import (
	"fmt"
	"testing"
)

const qtdRpt = 5

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", qtdRpt)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", qtdRpt)
	}
}

func ExampleRepetir() {
	repeticoes := Repetir("a", qtdRpt)
	fmt.Println(repeticoes)
	// Output: aaaaa
}
