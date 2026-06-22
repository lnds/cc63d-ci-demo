package main

// Total suma una lista de montos. Función pura: sin estado ni dependencias,
// perfecta para una prueba unitaria.
func Total(montos []int) int {
	suma := 0
	for _, m := range montos {
		suma += m
	}
	return suma
}
