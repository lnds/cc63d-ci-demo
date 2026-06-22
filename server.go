package main

import (
	"fmt"
	"net/http"
)

// health responde 200 OK. Es un handler HTTP: para probarlo de verdad hay que
// hacerle una petición, así que su prueba es de INTEGRACIÓN.
func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}
