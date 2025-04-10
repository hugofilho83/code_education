package main

import "testing"

// TestSoma testa a função Soma com diferentes cenários:
// - Soma de números positivos
// - Soma de números negativos
// - Soma de número positivo e negativo
// - Soma envolvendo o número zero
func TestSoma(t *testing.T) {
	t.Run("Soma de números positivos", func(t *testing.T) {
		result := Soma(1, 5)
		expected := 6

		if result != expected {
			t.Errorf("Resultado da soma é inválido: Resultado %d Esperado %d", result, expected)
		}
	})

	t.Run("Soma de números negativos", func(t *testing.T) {
		result := Soma(-3, -7)
		expected := -10

		if result != expected {
			t.Errorf("Resultado da soma é inválido: Resultado %d Esperado %d", result, expected)
		}
	})

	t.Run("Soma de número positivo e negativo", func(t *testing.T) {
		result := Soma(10, -4)
		expected := 6

		if result != expected {
			t.Errorf("Resultado da soma é inválido: Resultado %d Esperado %d", result, expected)
		}
	})

	t.Run("Soma com zero", func(t *testing.T) {
		result := Soma(0, 5)
		expected := 5

		if result != expected {
			t.Errorf("Resultado da soma é inválido: Resultado %d Esperado %d", result, expected)
		}
	})
}
