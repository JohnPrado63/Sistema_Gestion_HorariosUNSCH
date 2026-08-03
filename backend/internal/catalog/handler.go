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

func (h Handler) Departamentos(c *gin.Context) {
	data, err := h.repo.ListDepartamentos(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Escuelas(c *gin.Context) {
	data, err := h.repo.ListEscuelas(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Aulas(c *gin.Context) {
	data, err := h.repo.ListAulas(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Usuarios(c *gin.Context) {
	data, err := h.repo.ListUsuarios(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) PlanesEstudio(c *gin.Context) {
	data, err := h.repo.ListPlanesEstudio(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Series(c *gin.Context) {
	data, err := h.repo.ListSeries(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Cursos(c *gin.Context) {
	data, err := h.repo.ListCursos(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Docentes(c *gin.Context) {
	data, err := h.repo.ListDocentes(c.Request.Context())
	respond(c, data, err)
}

func (h Handler) Periodos(c *gin.Context) {
	data, err := h.repo.ListPeriodos(c.Request.Context())
	respond(c, data, err)
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
	data, err := h.repo.ListCargasAcademicas(c.Request.Context())
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

	hdr, err := h.repo.CreateHorario(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, hdr)
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

	data, err := h.repo.GetGruposParaHorario(c.Request.Context(), escuela, periodo)
	respond(c, data, err)
}

func (h Handler) CreateBloque(c *gin.Context) {
	var input CreateBloqueInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func (h Handler) VerificarConflictoBloque(c *gin.Context) {
	var input CreateBloqueInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
