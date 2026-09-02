package catalog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	repo RepositoryInterface
}

func NewHandler(db *pgxpool.Pool) Handler {
	return Handler{repo: NewRepository(db)}
}

func (h Handler) Facultades(c *gin.Context) {
	data, err := h.repo.ListFacultades(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) CreateFacultad(c *gin.Context) {
	var input struct {
		Nombre string `json:"nombre" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fac, err := h.repo.CreateFacultad(c.Request.Context(), input.Nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fac)
}

func (h Handler) UpdateFacultad(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Nombre string `json:"nombre" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fac, err := h.repo.UpdateFacultad(c.Request.Context(), id, input.Nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fac)
}

func (h Handler) DeleteFacultad(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeleteFacultad(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Facultad eliminada"})
}

func (h Handler) Departamentos(c *gin.Context) {
	data, err := h.repo.ListDepartamentos(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) CreateDepartamento(c *gin.Context) {
	var input struct {
		IDFacultad int    `json:"id_facultad" binding:"required"`
		Nombre     string `json:"nombre" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dept, err := h.repo.CreateDepartamento(c.Request.Context(), input.IDFacultad, input.Nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dept)
}

func (h Handler) UpdateDepartamento(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		IDFacultad int    `json:"id_facultad" binding:"required"`
		Nombre     string `json:"nombre" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dept, err := h.repo.UpdateDepartamento(c.Request.Context(), id, input.IDFacultad, input.Nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dept)
}

func (h Handler) DeleteDepartamento(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeleteDepartamento(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Departamento eliminado"})
}

func (h Handler) Escuelas(c *gin.Context) {
	data, err := h.repo.ListEscuelas(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) CreateEscuela(c *gin.Context) {
	var input CreateEscuelaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	esc, err := h.repo.CreateEscuela(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, esc)
}

func (h Handler) UpdateEscuela(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input UpdateEscuelaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	esc, err := h.repo.UpdateEscuela(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if esc == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Escuela no encontrada"})
		return
	}

	c.JSON(http.StatusOK, esc)
}

func (h Handler) DeleteEscuela(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeleteEscuela(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Escuela eliminada"})
}

func (h Handler) Aulas(c *gin.Context) {
	data, err := h.repo.ListAulas(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) AllAulas(c *gin.Context) {
	data, err := h.repo.ListAllAulas(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) SetAulaActivo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Activo bool `json:"activo"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.repo.SetAulaActivo(c.Request.Context(), id, input.Activo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aula actualizada correctamente"})
}

func (h Handler) CreateAula(c *gin.Context) {
	var input CreateAulaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aula, err := h.repo.CreateAula(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, aula)
}

func (h Handler) UpdateAula(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input UpdateAulaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aula, err := h.repo.UpdateAula(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if aula == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aula no encontrada"})
		return
	}

	c.JSON(http.StatusOK, aula)
}

func (h Handler) DeleteAula(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeleteAula(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aula eliminada correctamente"})
}

func (h Handler) Usuarios(c *gin.Context) {
	data, err := h.repo.ListUsuarios(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) PlanesEstudio(c *gin.Context) {
	data, err := h.repo.ListPlanesEstudio(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) CreatePlanEstudio(c *gin.Context) {
	var input CreatePlanEstudioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan, err := h.repo.CreatePlanEstudio(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, plan)
}

func (h Handler) UpdatePlanEstudio(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input UpdatePlanEstudioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plan, err := h.repo.UpdatePlanEstudio(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if plan == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan de estudio no encontrado"})
		return
	}

	c.JSON(http.StatusOK, plan)
}

func (h Handler) DeletePlanEstudio(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeletePlanEstudio(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plan de estudio eliminado"})
}

func (h Handler) Series(c *gin.Context) {
	escuela := c.Query("escuela")
	idEscuela := 0
	if escuela != "" {
		var err error
		idEscuela, err = strconv.Atoi(escuela)
		if err != nil {
			idEscuela = 0
		}
	}
	data, err := h.repo.ListSeries(c.Request.Context(), idEscuela)
	respond(c, data, err)
}

func (h Handler) Cursos(c *gin.Context) {
	escuela := c.Query("escuela")
	idEscuela := 0
	if escuela != "" {
		var err error
		idEscuela, err = strconv.Atoi(escuela)
		if err != nil {
			idEscuela = 0
		}
	}
	data, err := h.repo.ListCursos(c.Request.Context(), idEscuela)
	respond(c, data, err)
}

func (h Handler) CreateCurso(c *gin.Context) {
	var input CreateCursoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curso, err := h.repo.CreateCurso(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, curso)
}

func (h Handler) UpdateCurso(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input UpdateCursoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curso, err := h.repo.UpdateCurso(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if curso == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curso no encontrado"})
		return
	}

	c.JSON(http.StatusOK, curso)
}

func (h Handler) DeleteCurso(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeleteCurso(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Curso eliminado"})
}

func (h Handler) Docentes(c *gin.Context) {
	data, err := h.repo.ListDocentes(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) CreateDocente(c *gin.Context) {
	var input CreateDocenteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	docente, err := h.repo.CreateDocente(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, docente)
}

func (h Handler) UpdateDocente(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input UpdateDocenteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	docente, err := h.repo.UpdateDocente(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if docente == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Docente no encontrado"})
		return
	}

	c.JSON(http.StatusOK, docente)
}

func (h Handler) DeleteDocente(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeleteDocente(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Docente eliminado"})
}

func (h Handler) Periodos(c *gin.Context) {
	data, err := h.repo.ListPeriodos(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) CreatePeriodo(c *gin.Context) {
	var input CreatePeriodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	periodo, err := h.repo.CreatePeriodo(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, periodo)
}

func (h Handler) UpdatePeriodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input UpdatePeriodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	periodo, err := h.repo.UpdatePeriodo(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if periodo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Periodo no encontrado"})
		return
	}

	c.JSON(http.StatusOK, periodo)
}

func (h Handler) DeletePeriodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeletePeriodo(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Periodo eliminado"})
}

func (h Handler) SesionesDepartamento(c *gin.Context) {
	data, err := h.repo.ListSesionesDepartamento(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Locales(c *gin.Context) {
	data, err := h.repo.ListLocales(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Pabellones(c *gin.Context) {
	data, err := h.repo.ListPabellones(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) MatrizDistancias(c *gin.Context) {
	data, err := h.repo.ListDistancias(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) CargasAcademicas(c *gin.Context) {
	periodo := c.Query("periodo")
	escuela := c.Query("escuela")
	data, err := h.repo.ListCargasAcademicas(c.Request.Context(), periodo, escuela)
	respond(c, data, err)
}

func (h Handler) Grupos(c *gin.Context) {
	data, err := h.repo.ListGrupos(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Horarios(c *gin.Context) {
	data, err := h.repo.ListHorarios(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Bloques(c *gin.Context) {
	data, err := h.repo.ListBloquesHorario(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Bitacora(c *gin.Context) {
	data, err := h.repo.ListBitacoraAuditoria(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) CreateHorario(c *gin.Context) {
	var input CreateHorarioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existe, err := h.repo.ExistsHorario(c.Request.Context(), input.IDEscuela, input.IDPeriodo, input.IDSerie, input.Semestre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar horario existente"})
		return
	}
	if existe {
		c.JSON(http.StatusConflict, gin.H{"error": "Ya existe un horario para esta escuela, periodo, serie y semestre"})
		return
	}

	hdr, err := h.repo.CreateHorario(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, hdr)
}

func (h Handler) GenerateHorarios(c *gin.Context) {
	var input GenerateHorariosInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.repo.GenerateHorarios(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h Handler) DeleteHorario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeleteHorario(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Horario eliminado"})
}

func (h Handler) GetHorario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	data, err := h.repo.ListHorarios(c.Request.Context())
	if err != nil {
		respond(c, data, err)
		return
	}

	for _, h := range data {
		if h.ID == id {
			c.JSON(http.StatusOK, h)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Horario no encontrado"})
}

func (h Handler) GetBloquesByHorario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	data, err := h.repo.GetBloquesByHorario(c.Request.Context(), id)
	respond(c, data, err)
}

func (h Handler) GetGruposParaHorario(c *gin.Context) {
	escuelaStr := c.Query("escuela")
	periodoStr := c.Query("periodo")
	serieStr := c.Query("serie")
	semestre := c.Query("semestre")

	escuela, err := strconv.Atoi(escuelaStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de escuela inválido"})
		return
	}

	periodo, err := strconv.Atoi(periodoStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de periodo inválido"})
		return
	}

	var idSerie *int
	if serieStr != "" {
		val, err := strconv.Atoi(serieStr)
		if err == nil {
			idSerie = &val
		}
	}

	data, err := h.repo.GetGruposParaHorario(c.Request.Context(), escuela, periodo, idSerie, &semestre)
	respond(c, data, err)
}

func (h Handler) CreateBloque(c *gin.Context) {
	var input CreateBloqueInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.SlotFin <= input.SlotInicio {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slot_fin debe ser mayor que slot_inicio"})
		return
	}

	conflictos, err := h.repo.VerificarConflictoBloque(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(conflictos) > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error":      "Hay conflictos con el bloque propuesto",
			"conflictos": conflictos,
		})
		return
	}

	bloque, err := h.repo.CreateBloque(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, bloque)
}

func (h Handler) DeleteBloque(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.repo.DeleteBloque(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bloque eliminado"})
}

func (h Handler) VerificarConflictoBloque(c *gin.Context) {
	var input CreateBloqueInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.SlotFin <= input.SlotInicio {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slot_fin debe ser mayor que slot_inicio"})
		return
	}

	conflictos, err := h.repo.VerificarConflictoBloque(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tiene_conflicto": len(conflictos) > 0,
		"conflictos":       conflictos,
	})
}

func respond[T any](c *gin.Context, data T, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
