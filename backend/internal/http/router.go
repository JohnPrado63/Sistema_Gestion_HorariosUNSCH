package http

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"unsch-horarios/backend/internal/catalog"
	"unsch-horarios/backend/internal/health"
	"unsch-horarios/backend/internal/schedule"
	"unsch-horarios/backend/internal/validationui"
)

func NewRouter(db *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()

	healthHandler := health.NewHandler(db)
	catalogHandler := catalog.NewHandler(db)
	validationHandler := validationui.NewHandler()
	scheduleHandler := schedule.NewHandler()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/health", healthHandler.Health)
	mux.HandleFunc("/ready", healthHandler.Ready)
	mux.HandleFunc("/api/v1/facultades", catalogHandler.Facultades)
	mux.HandleFunc("/api/v1/departamentos", catalogHandler.Departamentos)
	mux.HandleFunc("/api/v1/escuelas", catalogHandler.Escuelas)
	mux.HandleFunc("/api/v1/aulas", catalogHandler.Aulas)
	mux.HandleFunc("/api/v1/usuarios", catalogHandler.Usuarios)
	mux.HandleFunc("/api/v1/planes", catalogHandler.PlanesEstudio)
	mux.HandleFunc("/api/v1/series", catalogHandler.Series)
	mux.HandleFunc("/api/v1/cursos", catalogHandler.Cursos)
	mux.HandleFunc("/api/v1/docentes", catalogHandler.Docentes)
	mux.HandleFunc("/api/v1/periodos", catalogHandler.Periodos)
	mux.HandleFunc("/api/v1/sesiones-departamento", catalogHandler.SesionesDepartamento)
	mux.HandleFunc("/api/v1/locales", catalogHandler.Locales)
	mux.HandleFunc("/api/v1/pabellones", catalogHandler.Pabellones)
	mux.HandleFunc("/api/v1/matriz-distancias", catalogHandler.MatrizDistancias)
	mux.HandleFunc("/api/v1/cargas-academicas", catalogHandler.CargasAcademicas)
	mux.HandleFunc("/api/v1/grupos", catalogHandler.Grupos)
	mux.HandleFunc("/api/v1/horarios", catalogHandler.Horarios)
	mux.HandleFunc("/api/v1/bloques", catalogHandler.Bloques)
	mux.HandleFunc("/api/v1/bitacora", catalogHandler.Bitacora)
	mux.HandleFunc("/validaciones", validationHandler.Page)
	mux.HandleFunc("/api/v1/validaciones/escenarios", validationHandler.Scenarios)
	mux.HandleFunc("/api/v1/validaciones/placement", scheduleHandler.ValidatePlacement)
	mux.HandleFunc("/api/v1/validaciones/audit", scheduleHandler.ValidateAuditChange)
	mux.HandleFunc("/api/v1/validaciones/carga", scheduleHandler.ValidateTeachingLoad)

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
