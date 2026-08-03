package cargaacademica

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	repo *Repository
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{repo: NewRepository(db)}
}

func (h *Handler) ListCargas(c *gin.Context) {
	idPeriodoStr := c.Query("periodo")
	idEscuelaStr := c.Query("escuela")

	idPeriodo, err := strconv.Atoi(idPeriodoStr)
	if err != nil {
		idPeriodo = 0
	}

	idEscuela, err := strconv.Atoi(idEscuelaStr)
	if err != nil {
		idEscuela = 0
	}

	var cargas []CargaAcademica

	if idEscuela > 0 && idPeriodo > 0 {
		cargas, err = h.repo.ListCargasByEscuela(c.Request.Context(), idEscuela, idPeriodo)
	} else if idPeriodo > 0 {
		cargas, err = h.repo.ListCargasByPeriodo(c.Request.Context(), idPeriodo)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere parametro periodo"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i := range cargas {
		grupos, err := h.repo.GetGruposByCarga(c.Request.Context(), cargas[i].IDCarga)
		if err == nil {
			cargas[i].Grupos = grupos
		}
	}

	c.JSON(http.StatusOK, cargas)
}

func (h *Handler) GetCarga(c *gin.Context) {
	idCarga, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	carga, err := h.repo.GetCargaByID(c.Request.Context(), idCarga)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	grupos, err := h.repo.GetGruposByCarga(c.Request.Context(), idCarga)
	if err == nil {
		carga.Grupos = grupos
	}

	c.JSON(http.StatusOK, carga)
}

func (h *Handler) CreateCarga(c *gin.Context) {
	var input CreateCargaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carga, err := h.repo.CreateCarga(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, carga)
}

func (h *Handler) CreateGrupo(c *gin.Context) {
	idCarga, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var input CreateGrupoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.IDDocente != nil {
		periodoActivo, err := h.repo.GetPeriodoActivo(c.Request.Context())
		if err == nil {
			horasDocente, _ := h.repo.GetHorasDocente(c.Request.Context(), *input.IDDocente, *periodoActivo)
			cursoHoras := 0
			carga, _ := h.repo.GetCargaByID(c.Request.Context(), idCarga)
			if carga != nil {
				cursoHoras = carga.Curso.HorasTeoria + carga.Curso.HorasPractica
			}
			if horasDocente+cursoHoras > 16 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":            "Docente excederia las 16 horas lectivas",
					"horas_actuales":   horasDocente,
					"horas_nueva_curso": cursoHoras,
				})
				return
			}
		}
	}

	grupo, err := h.repo.CreateGrupo(c.Request.Context(), idCarga, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, grupo)
}

func (h *Handler) UpdateGrupo(c *gin.Context) {
	idGrupo, err := strconv.Atoi(c.Param("idGrupo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var input UpdateGrupoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.IDDocente != nil {
		periodoActivo, err := h.repo.GetPeriodoActivo(c.Request.Context())
		if err == nil {
			horasDocente, _ := h.repo.GetHorasDocente(c.Request.Context(), *input.IDDocente, *periodoActivo)
			if horasDocente > 16 {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":          "Docente excederia las 16 horas lectivas",
					"horas_actuales": horasDocente,
				})
				return
			}
		}
	}

	grupo, err := h.repo.UpdateGrupo(c.Request.Context(), idGrupo, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if grupo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grupo no encontrado"})
		return
	}

	c.JSON(http.StatusOK, grupo)
}

func (h *Handler) ApproveCarga(c *gin.Context) {
	idCarga, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	var input AprobarCargaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		input.IDUsuario = 0
		input.Justificacion = "Aprobacion automatica"
	}

	err = h.repo.ApproveCarga(c.Request.Context(), idCarga, input.IDUsuario, input.Justificacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Carga aprobada correctamente"})
}

func (h *Handler) GetResumenDocentes(c *gin.Context) {
	idEscuelaStr := c.Query("escuela")
	idPeriodoStr := c.Query("periodo")

	idEscuela, _ := strconv.Atoi(idEscuelaStr)
	idPeriodo, err := strconv.Atoi(idPeriodoStr)
	if idPeriodo == 0 {
		periodoActivo, err := h.repo.GetPeriodoActivo(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No hay periodo activo"})
			return
		}
		idPeriodo = *periodoActivo
	}

	var resumenes []ResumenDocente
	if idEscuela > 0 {
		resumenes, err = h.repo.GetDocentesResumen(c.Request.Context(), idEscuela, idPeriodo)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere parametro escuela"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resumenes)
}

func (h *Handler) GetHorasDocente(c *gin.Context) {
	idDocente, err := strconv.Atoi(c.Param("idDocente"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	idPeriodoStr := c.Query("periodo")
	idPeriodo, err := strconv.Atoi(idPeriodoStr)
	if idPeriodo == 0 {
		periodoActivo, err := h.repo.GetPeriodoActivo(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No hay periodo activo"})
			return
		}
		idPeriodo = *periodoActivo
	}

	horas, err := h.repo.GetHorasDocente(c.Request.Context(), idDocente, idPeriodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id_docente":      idDocente,
		"id_periodo":      idPeriodo,
		"horas_asignadas": horas,
		"horas_restantes": 16 - horas,
	})
}

func (h *Handler) GetDisponibilidadDocente(c *gin.Context) {
	idDocente, err := strconv.Atoi(c.Param("idDocente"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalido"})
		return
	}

	idPeriodoStr := c.Query("periodo")
	idPeriodo, err := strconv.Atoi(idPeriodoStr)
	if idPeriodo == 0 {
		periodoActivo, err := h.repo.GetPeriodoActivo(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No hay periodo activo"})
			return
		}
		idPeriodo = *periodoActivo
	}

	bloques, err := h.repo.GetBloquesDocente(c.Request.Context(), idDocente, idPeriodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"docente_id":        idDocente,
		"id_periodo":        idPeriodo,
		"bloques_asignados": bloques,
	})
}

func (h *Handler) VerificarDisponibilidad(c *gin.Context) {
	var input VerificarDisponibilidadInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idPeriodoStr := c.Query("periodo")
	idPeriodo, err := strconv.Atoi(idPeriodoStr)
	if idPeriodo == 0 {
		periodoActivo, err := h.repo.GetPeriodoActivo(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No hay periodo activo"})
			return
		}
		idPeriodo = *periodoActivo
	}

	conflicto, err := h.repo.VerificarConflictoDocente(c.Request.Context(), input.DocenteID, input.DiaSemana, input.SlotInicio, input.SlotFin, idPeriodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := VerificarDisponibilidadResponse{
		TieneConflicto: conflicto != nil,
	}

	if conflicto != nil {
		dias := []string{"", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}
		horas := map[int]string{
			1: "07:00", 2: "08:00", 3: "09:00", 4: "10:00", 5: "11:00", 6: "12:00",
			7: "13:00", 8: "14:00", 9: "15:00", 10: "16:00", 11: "17:00", 12: "18:00",
			13: "19:00", 14: "20:00",
		}
		slotSolicitado := horas[input.SlotInicio] + "-" + horas[input.SlotFin]
		slotOcupado := horas[conflicto.SlotInicio] + "-" + horas[conflicto.SlotFin]
		nombreDia := dias[conflicto.DiaSemana]
		if nombreDia == "" {
			nombreDia = "Día " + strconv.Itoa(conflicto.DiaSemana)
		}

		response.Conflictos = []ConflictoDetalle{
			{
				Tipo:           conflicto.Tipo,
				DiaSemana:      conflicto.DiaSemana,
				SlotSolicitado: slotSolicitado,
				SlotOcupado:    slotOcupado,
				Escuela:        conflicto.Escuela,
				CursoCodigo:    conflicto.CursoCodigo,
				CursoNombre:    conflicto.CursoNombre,
				GrupoCodigo:    conflicto.GrupoCodigo,
			},
		}
	}

	c.JSON(http.StatusOK, response)
}
