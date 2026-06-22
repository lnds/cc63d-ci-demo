//go:build integration

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Prueba de INTEGRACIÓN: ejercita el handler HTTP de punta a punta.
// La etiqueta de arriba (//go:build integration) hace que SOLO corra con
// `go test -tags=integration`, no en la corrida unitaria normal.

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	health(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d; quiero 200", w.Code)
	}
	if w.Body.String() != "ok" {
		t.Fatalf("body = %q; quiero \"ok\"", w.Body.String())
	}
}
