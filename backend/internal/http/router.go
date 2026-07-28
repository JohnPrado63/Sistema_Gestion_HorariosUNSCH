package http

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"unsch-horarios/backend/internal/catalog"
	"unsch-horarios/backend/internal/health"
)

func NewRouter(db *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()

	healthHandler := health.NewHandler(db)
	catalogHandler := catalog.NewHandler(db)

	mux.HandleFunc("GET /", index)
	mux.HandleFunc("GET /health", healthHandler.Health)
	mux.HandleFunc("GET /ready", healthHandler.Ready)
	mux.HandleFunc("GET /api/v1/facultades", catalogHandler.Facultades)
	mux.HandleFunc("GET /api/v1/departamentos", catalogHandler.Departamentos)
	mux.HandleFunc("GET /api/v1/escuelas", catalogHandler.Escuelas)
	mux.HandleFunc("GET /api/v1/aulas", catalogHandler.Aulas)
	mux.HandleFunc("GET /api/v1/usuarios", catalogHandler.Usuarios)

	return withJSONHeaders(mux)
}

func index(w http.ResponseWriter, _ *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]any{
		"name":   "Sistema de Gestion de Horarios UNSCH",
		"status": "running",
		"endpoints": []string{
			"GET /health",
			"GET /ready",
			"GET /api/v1/facultades",
			"GET /api/v1/departamentos",
			"GET /api/v1/escuelas",
			"GET /api/v1/aulas",
			"GET /api/v1/usuarios",
		},
	})
}

func withJSONHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
