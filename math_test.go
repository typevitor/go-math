package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 15)
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resuldado: %d. Esperado era: %d", total, 30)
	}
}

func TestSub(t *testing.T) {
	total := Sub(15, 10)
	if total != 5 {
		t.Errorf("Resultado da subtração é inválido: Resuldado: %d. Esperado era: %d", total, 5)
	}
}

func TestTimes(t *testing.T) {
	total := Times(15, 2)
	if total != 30 {
		t.Errorf("Resultado da multiplicação é inválido: Resuldado: %d. Esperado era: %d", total, 30)
	}
}

func TestDiv(t *testing.T) {
	total,err := Div(15, 3)
	if total != 5 {
		t.Errorf("Resultado da divisão é inválido: Resuldado: %.2f. Esperado era: %.2f", total, 5.00)
	}
	if err != nil {
		t.Errorf("Erro não deveria ser retornado");
	}

	total,err = Div(10, 0)
	if err == nil {
		t.Errorf("Divisor não pode ser nulo.");
	}
}