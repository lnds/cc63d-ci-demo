package main

import "testing"

// Pruebas UNITARIAS: rápidas, sin red ni base de datos. Corren con `go test`.

func TestTotal(t *testing.T) {
	if got := Total([]int{10, 20, 5}); got != 35 {
		t.Fatalf("Total = %d; quiero 35", got)
	}
}

func TestTotalVacio(t *testing.T) {
	if got := Total(nil); got != 0 {
		t.Fatalf("Total = %d; quiero 0", got)
	}
}
