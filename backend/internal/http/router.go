package http

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"unsch-horarios/backend/internal/auth"
	"unsch-horarios/backend/internal/cargaacademica"
	"unsch-horarios/backend/internal/catalog"
	"unsch-horarios/backend/internal/health"
	"unsch-horarios/backend/internal/schedule"
	"unsch-horarios/backend/internal/validationui"
)

func NewRouter(db *pgxpool.Pool) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/health", "/ready"},
	}))
	r.Use(gin.Recovery())
	r.Use(corsMiddleware())
	r.Use(rateLimitMiddleware(100, 60, 60*time.Second))

	r.GET("/health", healthHandler(db).Health)
	r.GET("/ready", healthHandler(db).Ready)
	r.GET("/", index)

	r.POST("/api/v1/login", auth.LoginHandler(db))

	cat := catalogHandler(db)
	sch := scheduleHandler()
	ca := cargaacademicaHandler(db)

	api := r.Group("/api/v1")
	{
		api.GET("/facultades", cat.Facultades)
		api.GET("/departamentos", cat.Departamentos)
		api.GET("/escuelas", cat.Escuelas)
		api.GET("/aulas", cat.Aulas)
		api.GET("/usuarios", cat.Usuarios)
		api.GET("/planes", cat.PlanesEstudio)
		api.GET("/series", cat.Series)
		api.GET("/cursos", cat.Cursos)
		api.GET("/docentes", cat.Docentes)
		api.GET("/periodos", cat.Periodos)
		api.GET("/sesiones-departamento", cat.SesionesDepartamento)
		api.GET("/locales", cat.Locales)
		api.GET("/pabellones", cat.Pabellones)
		api.GET("/matriz-distancias", cat.MatrizDistancias)
		api.GET("/cargas-academicas", cat.CargasAcademicas)
		api.GET("/grupos", cat.Grupos)
		api.GET("/horarios", cat.Horarios)
		api.POST("/horarios", cat.CreateHorario)
		api.GET("/horarios/:id", cat.GetHorario)
		api.GET("/horarios/:id/bloques", cat.GetBloquesByHorario)
		api.GET("/grupos-horario", cat.GetGruposParaHorario)
		api.GET("/bloques", cat.Bloques)
		api.POST("/bloques", cat.CreateBloque)
		api.POST("/bloques/verificar", cat.VerificarConflictoBloque)
		api.GET("/bitacora", cat.Bitacora)

		cargaRoutes := api.Group("/carga-academica")
		{
			cargaRoutes.GET("", ca.ListCargas)
			cargaRoutes.GET("/:id", ca.GetCarga)
			cargaRoutes.POST("", ca.CreateCarga)
			cargaRoutes.POST("/:id/grupos", ca.CreateGrupo)
			cargaRoutes.PUT("/grupos/:idGrupo", ca.UpdateGrupo)
			cargaRoutes.POST("/:id/aprobar", ca.ApproveCarga)
			cargaRoutes.GET("/resumen-docentes", ca.GetResumenDocentes)
			cargaRoutes.GET("/docente/:idDocente/horas", ca.GetHorasDocente)
		}

		disponibilidadRoutes := api.Group("/disponibilidad")
		{
			disponibilidadRoutes.GET("/docente/:idDocente", ca.GetDisponibilidadDocente)
			disponibilidadRoutes.POST("/verificar", ca.VerificarDisponibilidad)
		}

		validaciones := api.Group("/validaciones")
		{
			validaciones.GET("/escenarios", validationHandler().Scenarios)
			validaciones.POST("/placement", auth.AuthMiddleware(), auth.RequireRole("ADMIN_TI", "DIRECTOR_ESCUELA", "JEFE_DEPTO", "DGA"), sch.ValidatePlacement)
			validaciones.POST("/audit", auth.AuthMiddleware(), auth.RequireRole("ADMIN_TI", "DIRECTOR_ESCUELA", "JEFE_DEPTO", "DGA"), sch.ValidateAuditChange)
			validaciones.POST("/carga", auth.AuthMiddleware(), auth.RequireRole("ADMIN_TI", "DIRECTOR_ESCUELA", "JEFE_DEPTO", "DGA"), sch.ValidateTeachingLoad)
		}
	}

	r.GET("/validaciones", validationHandler().Page)

	frontend := setupFrontend()
	if frontend != "" {
		r.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path
			if path == "/app" || path == "/app/" || path == "/app/index.html" {
				c.File(filepath.Join(frontend, "index.html"))
				return
			}
			if strings.HasPrefix(path, "/app/") {
				filePath := filepath.Join(frontend, strings.TrimPrefix(path, "/app/"))
				if _, err := os.Stat(filePath); err == nil {
					c.File(filePath)
					return
				}
				indexPath := filepath.Join(frontend, "index.html")
				if _, err := os.Stat(indexPath); err == nil {
					c.File(indexPath)
					return
				}
			}
			c.JSON(404, gin.H{"error": "Not found"})
		})
	}

	return r
}

func healthHandler(db *pgxpool.Pool) health.Handler {
	return health.NewHandler(db)
}

func catalogHandler(db *pgxpool.Pool) catalog.Handler {
	return catalog.NewHandler(db)
}

func scheduleHandler() schedule.Handler {
	return schedule.NewHandler()
}

func validationHandler() validationui.Handler {
	return validationui.NewHandler()
}

func cargaacademicaHandler(db *pgxpool.Pool) *cargaacademica.Handler {
	return cargaacademica.NewHandler(db)
}

func setupFrontend() string {
	candidates := []struct {
		index  string
		folder string
	}{
		{"./frontend/dist/index.html", "./frontend/dist"},
		{"../frontend/dist/index.html", "../frontend/dist"},
		{"backend/frontend/dist/index.html", "backend/frontend/dist"},
		{"./frontend/index.html", "./frontend"},
		{"../frontend/index.html", "../frontend"},
	}

	for _, c := range candidates {
		if _, err := os.Stat(c.index); err == nil {
			log.Printf("frontend: using %s", c.folder)
			return c.folder
		}
	}

	log.Printf("frontend: no dist folder found")
	return ""
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func rateLimitMiddleware(requests int, burst int, window time.Duration) gin.HandlerFunc {
	type client struct {
		count    int
		lastSeen time.Time
	}

	clients := make(map[string]*client)

	go func() {
		for {
			for k, v := range clients {
				if time.Since(v.lastSeen) > window {
					delete(clients, k)
				}
			}
			time.Sleep(time.Second)
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()
		cl, exists := clients[ip]

		if !exists {
			clients[ip] = &client{count: 1, lastSeen: time.Now()}
			c.Next()
			return
		}

		if time.Since(cl.lastSeen) > window {
			cl.count = 1
			cl.lastSeen = time.Now()
			c.Next()
			return
		}

		if cl.count >= requests {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Demasiadas solicitudes"})
			c.Abort()
			return
		}

		cl.count++
		cl.lastSeen = time.Now()
		c.Next()
	}
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "Sistema de Gestion de Horarios UNSCH",
		"status": "running",
		"version": "1.0.0",
		"endpoints": []string{
			"POST /api/v1/login",
			"GET /health",
			"GET /api/v1/*catalogs",
		},
	})
}
