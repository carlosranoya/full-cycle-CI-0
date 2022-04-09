package main

import "testing"

func TestMDC(t *testing.T) {

	mdc := MDC(15, 24)
	var expected uint64 = 3
	if mdc != expected {
		t.Errorf("Resultado do MDC é inválido.\nResultado: %d\nEsperado: %d", mdc, expected)
	}

}
