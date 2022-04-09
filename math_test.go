package main

import "testing"

func TestMDC(t *testing.T) {

	mdc := MDC(15, 25)
	var expected uint64 = 5
	if mdc != expected {
		t.Errorf("Resultado do MDC é inválido.\nResultado: %d\nEsperado: %d", mdc, expected)
	}

}
