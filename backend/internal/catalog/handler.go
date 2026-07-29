package catalog

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	repo Repository
}

func NewHandler(db *pgxpool.Pool) Handler {
	return Handler{repo: NewRepository(db)}
}

func (h Handler) Facultades(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListFacultades(r.Context())
	respond(w, data, err)
}

func (h Handler) Departamentos(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListDepartamentos(r.Context())
	respond(w, data, err)
}

func (h Handler) Escuelas(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListEscuelas(r.Context())
	respond(w, data, err)
}

func (h Handler) Aulas(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListAulas(r.Context())
	respond(w, data, err)
}

func (h Handler) Usuarios(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListUsuarios(r.Context())
	respond(w, data, err)
}

func (h Handler) PlanesEstudio(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListPlanesEstudio(r.Context())
	respond(w, data, err)
}

func (h Handler) Series(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListSeries(r.Context())
	respond(w, data, err)
}

func (h Handler) Cursos(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListCursos(r.Context())
	respond(w, data, err)
}

func (h Handler) Docentes(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListDocentes(r.Context())
	respond(w, data, err)
}

func (h Handler) Periodos(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListPeriodos(r.Context())
	respond(w, data, err)
}

func (h Handler) SesionesDepartamento(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListSesionesDepartamento(r.Context())
	respond(w, data, err)
}

func (h Handler) Locales(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListLocales(r.Context())
	respond(w, data, err)
}

func (h Handler) Pabellones(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListPabellones(r.Context())
	respond(w, data, err)
}

func (h Handler) MatrizDistancias(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListDistancias(r.Context())
	respond(w, data, err)
}

func (h Handler) CargasAcademicas(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListCargasAcademicas(r.Context())
	respond(w, data, err)
}

func (h Handler) Grupos(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListGrupos(r.Context())
	respond(w, data, err)
}

func (h Handler) Horarios(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListHorarios(r.Context())
	respond(w, data, err)
}

func (h Handler) Bloques(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListBloquesHorario(r.Context())
	respond(w, data, err)
}

func (h Handler) Bitacora(w http.ResponseWriter, r *http.Request) {
	data, err := h.repo.ListBitacoraAuditoria(r.Context())
	respond(w, data, err)
}

func respond[T any](w http.ResponseWriter, data T, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	_ = json.NewEncoder(w).Encode(data)
}
