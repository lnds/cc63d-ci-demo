# cc63d-ci-demo — un pipeline de CI, paso a paso

Ejemplo de la **Clase 5 (CI/CD)** del curso CC63D. Una mini aplicación en **Go**
con una pipeline de GitHub Actions que muestra las tres etapas de un CI:

```
compilar  →  pruebas unitarias  →  pruebas de integración
```

Lo usamos en clase para entender cómo se arma un workflow antes de hacerlo tú
en el laboratorio. (Go, además, adelanta la Clase 6: microservicios políglotas.)

## La aplicación

| Archivo | Qué es | Cómo se prueba |
|---------|--------|----------------|
| `total.go` | función pura `Total([]int) int` | **unitaria** (rápida, sin dependencias) |
| `server.go` | handler HTTP `GET /health` | **integración** (se le hace una petición real) |

## Las pruebas

```bash
go build ./...                    # 1. ¿compila?
go test ./...                     # 2. unitarias (total_test.go)
go test -tags=integration ./...   # 3. + integración (server_integration_test.go)
```

La prueba de integración lleva la etiqueta `//go:build integration`, así que
**solo** corre cuando pasas `-tags=integration`. Las unitarias corren siempre.

## La pipeline

[`.github/workflows/ci.yml`](.github/workflows/ci.yml) ejecuta esas tres etapas
en cada `push` y cada `pull_request`. Si una falla, el run queda **rojo**.


## test
test para demo
