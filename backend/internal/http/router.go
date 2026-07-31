package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	// validation UI and APIs
	mux.HandleFunc("/validaciones", validationHandler.Page)
	mux.HandleFunc("/api/v1/validaciones/escenarios", validationHandler.Scenarios)
	mux.HandleFunc("/api/v1/validaciones/placement", scheduleHandler.ValidatePlacement)
	mux.HandleFunc("/api/v1/validaciones/audit", scheduleHandler.ValidateAuditChange)
	mux.HandleFunc("/api/v1/validaciones/carga", scheduleHandler.ValidateTeachingLoad)

	// serve the frontend static app under /app/
	// Priority: FRONTEND_DIR env var > dist folder > source folder
	frontendDir := os.Getenv("FRONTEND_DIR")
	if frontendDir == "" {
		candidates := []struct {
			path     string
			frontend string
		}{
			{"./frontend/dist/index.html", "./frontend/dist"},
			{"../frontend/dist/index.html", "../frontend/dist"},
			{"backend/frontend/dist/index.html", "backend/frontend/dist"},
			{"./frontend/index.html", "./frontend"},
			{"../frontend/index.html", "../frontend"},
		}
		for _, c := range candidates {
			if _, err := os.Stat(c.path); err == nil {
				frontendDir = c.frontend
				log.Printf("frontend: using %s (found %s)", frontendDir, c.path)
				break
			}
		}
		if frontendDir == "" {
			frontendDir = "./frontend"
			log.Printf("frontend: using fallback %s", frontendDir)
		}
	}
	fileServer := http.FileServer(http.Dir(frontendDir))
	mux.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		// if request targets the directory root, serve index.html with proper content-type
		if r.URL.Path == "/app/" || r.URL.Path == "/app" || r.URL.Path == "/app/index.html" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			http.ServeFile(w, r, filepath.Join(frontendDir, "index.html"))
			return
		}
		// otherwise serve static files from frontendDir
		http.StripPrefix("/app/", fileServer).ServeHTTP(w, r)
	})
	mux.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/app/", http.StatusFound) })

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
		// Only set application/json for API endpoints. Allow static files and UI to set their own content-type.
		if r.URL.Path == "/" || r.URL.Path == "/health" || r.URL.Path == "/ready" || strings.HasPrefix(r.URL.Path, "/api/") || strings.HasPrefix(r.URL.Path, "/validaciones") {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
