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

func respond[T any](w http.ResponseWriter, data T, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	_ = json.NewEncoder(w).Encode(data)
}
