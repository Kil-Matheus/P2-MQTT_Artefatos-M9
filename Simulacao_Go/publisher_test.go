package main

import (
	"fmt"
	"testing"
)

func TestSimularLeituraSensor(t *testing.T) {
	zl1, zl2 := simularLeituraSensor()

	if zl1 < 0 || zl1 > 100 {
		t.Errorf("Valor de zl1 fora do intervalo esperado: %.2f", zl1)
	}

	if zl2 < 0 || zl2 > 100 {
		t.Errorf("Valor de zl2 fora do intervalo esperado: %.2f", zl2)
	}
}

func ExampleSimularLeituraSensor() {
	zl1, zl2 := simularLeituraSensor()
	fmt.Printf("Leitura do sensor: PM2.5 %.2f (µg/m³), PM10 %.2f (µg/m³), Penha - Zona Leste - SP\n", zl1, zl2)
}

func Testmain(){
	
}