package main

import "testing"

// Teste de Soma
func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado %d. Esperado %d", total, 30)
	}
}
