package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockRepository struct {
	facultades      []Facultad
	departamentos   []Departamento
	escuelas        []Escuela
	aulas           []Aula
	usuarios        []Usuario
	planes          []PlanEstudio
	series          []Serie
	cursos          []Curso
	docentes        []Docente
	periodos        []PeriodoAcademico
	sesiones        []SesionDepartamento
	locales         []Local
	pabellones      []Pabellon
	distancias      []Distancia
	cargas          []CargaAcademica
	grupos          []Grupo
	horarios        []Horario
	bloques         []BloqueHorario
	bitacora        []BitacoraAuditoria
	err             error
}

func (m *mockRepository) ListFacultades(ctx context.Context) ([]Facultad, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.facultades, nil
}
func (m *mockRepository) ListDepartamentos(ctx context.Context) ([]Departamento, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.departamentos, nil
}
func (m *mockRepository) ListEscuelas(ctx context.Context) ([]Escuela, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.escuelas, nil
}
func (m *mockRepository) ListAulas(ctx context.Context) ([]Aula, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.aulas, nil
}
func (m *mockRepository) ListUsuarios(ctx context.Context) ([]Usuario, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.usuarios, nil
}
func (m *mockRepository) ListPlanesEstudio(ctx context.Context) ([]PlanEstudio, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.planes, nil
}
func (m *mockRepository) ListSeries(ctx context.Context) ([]Serie, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.series, nil
}
func (m *mockRepository) ListCursos(ctx context.Context) ([]Curso, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.cursos, nil
}
func (m *mockRepository) ListDocentes(ctx context.Context) ([]Docente, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.docentes, nil
}
func (m *mockRepository) ListPeriodos(ctx context.Context) ([]PeriodoAcademico, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.periodos, nil
}
func (m *mockRepository) ListSesionesDepartamento(ctx context.Context) ([]SesionDepartamento, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.sesiones, nil
}
func (m *mockRepository) ListLocales(ctx context.Context) ([]Local, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.locales, nil
}
func (m *mockRepository) ListPabellones(ctx context.Context) ([]Pabellon, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.pabellones, nil
}
func (m *mockRepository) ListDistancias(ctx context.Context) ([]Distancia, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.distancias, nil
}
func (m *mockRepository) ListCargasAcademicas(ctx context.Context) ([]CargaAcademica, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.cargas, nil
}
func (m *mockRepository) ListGrupos(ctx context.Context) ([]Grupo, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.grupos, nil
}
func (m *mockRepository) ListHorarios(ctx context.Context) ([]Horario, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.horarios, nil
}
func (m *mockRepository) ListBloquesHorario(ctx context.Context) ([]BloqueHorario, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.bloques, nil
}
func (m *mockRepository) ListBitacoraAuditoria(ctx context.Context) ([]BitacoraAuditoria, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.bitacora, nil
}

func testContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}

func TestFacultades_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		facultades: []Facultad{
			{ID: 1, Nombre: "Facultad de Ingeniería"},
			{ID: 2, Nombre: "Facultad de Medicina"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/facultades", nil)
	c, w := testContext()
	c.Request = req

	h.Facultades(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var result []Facultad
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 facultades, got %d", len(result))
	}
}

func TestFacultades_Returns500OnError(t *testing.T) {
	repo := &mockRepository{err: errors.New("database connection failed")}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/facultades", nil)
	c, w := testContext()
	c.Request = req

	h.Facultades(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", w.Code)
	}
}

func TestDepartamentos_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		departamentos: []Departamento{
			{ID: 1, IDFacultad: 1, Nombre: "Departamento de Computación"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/departamentos", nil)
	c, w := testContext()
	c.Request = req

	h.Departamentos(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var result []Departamento
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 departamento, got %d", len(result))
	}
}

func TestEscuelas_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		escuelas: []Escuela{
			{ID: 1, IDFacultad: 1, IDDepartamento: 1, Nombre: "Escuela de Ingeniería de Sistemas"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/escuelas", nil)
	c, w := testContext()
	c.Request = req

	h.Escuelas(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestAulas_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		aulas: []Aula{
			{ID: 1, IDPabellon: 1, Codigo: "A-101", Tipo: "TEORIA", Aforo: 40, EsCompartida: false},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/aulas", nil)
	c, w := testContext()
	c.Request = req

	h.Aulas(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestDocentes_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		docentes: []Docente{
			{ID: 1, IDDepartamento: 1, CodigoPlaza: "DOC001", Nombres: "Juan", Apellidos: "Pérez"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/docentes", nil)
	c, w := testContext()
	c.Request = req

	h.Docentes(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestPeriodos_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		periodos: []PeriodoAcademico{
			{ID: 1, Codigo: "2026-I", Activo: true},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/periodos", nil)
	c, w := testContext()
	c.Request = req

	h.Periodos(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestHorarios_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		horarios: []Horario{
			{ID: 1, IDEscuela: 1, IDPeriodo: 1, Estado: "BORRADOR", VersionReajuste: 0},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/horarios", nil)
	c, w := testContext()
	c.Request = req

	h.Horarios(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestBloques_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		bloques: []BloqueHorario{
			{ID: 1, IDHorario: 1, IDGrupo: 1, IDAula: 1, DiaSemana: 1, SlotInicio: 3, SlotFin: 5},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/bloques", nil)
	c, w := testContext()
	c.Request = req

	h.Bloques(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestBitacora_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		bitacora: []BitacoraAuditoria{
			{ID: 1, IDHorario: 1, IDUsuario: 1, Accion: "CREAR", MotivoJustificacion: "Horario inicial", VersionResultante: 0},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/bitacora", nil)
	c, w := testContext()
	c.Request = req

	h.Bitacora(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestCargasAcademicas_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		cargas: []CargaAcademica{
			{IDCarga: 1, IDCurso: 1, IDPeriodo: 1, IDEscuela: 1, Estado: "BORRADOR"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/cargas-academicas", nil)
	c, w := testContext()
	c.Request = req

	h.CargasAcademicas(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestGrupos_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		grupos: []Grupo{
			{ID: 1, IDCarga: 1, CodigoGrupo: "Grupo A", TipoComponente: "TEORIA", MatriculadosProyectados: 40},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/grupos", nil)
	c, w := testContext()
	c.Request = req

	h.Grupos(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestPlanesEstudio_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		planes: []PlanEstudio{
			{ID: 1, IDEscuela: 1, Codigo: "2024", Nombre: "Plan 2024"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/planes", nil)
	c, w := testContext()
	c.Request = req

	h.PlanesEstudio(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestSeries_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		series: []Serie{
			{ID: 1, IDPlan: 1, NumeroCiclo: 100},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/series", nil)
	c, w := testContext()
	c.Request = req

	h.Series(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestCursos_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		cursos: []Curso{
			{ID: 1, IDSerie: 1, Codigo: "CS101", Nombre: "Introducción a la Programación", Creditos: 4},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/cursos", nil)
	c, w := testContext()
	c.Request = req

	h.Cursos(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestSesionesDepartamento_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		sesiones: []SesionDepartamento{
			{ID: 1, IDDepartamento: 1, IDPeriodo: 1, DiaSemana: 1, HoraInicio: "08:00", HoraFin: "10:00"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sesiones-departamento", nil)
	c, w := testContext()
	c.Request = req

	h.SesionesDepartamento(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestLocales_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		locales: []Local{
			{ID: 1, Nombre: "Campus Central"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/locales", nil)
	c, w := testContext()
	c.Request = req

	h.Locales(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestPabellones_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		pabellones: []Pabellon{
			{ID: 1, IDLocal: 1, Codigo: "PAB-A", Nombre: "Pabellón A"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/pabellones", nil)
	c, w := testContext()
	c.Request = req

	h.Pabellones(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestMatrizDistancias_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		distancias: []Distancia{
			{DesdeID: 1, HastaID: 2, Minutos: 10},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/matriz-distancias", nil)
	c, w := testContext()
	c.Request = req

	h.MatrizDistancias(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestUsuarios_ReturnsJSON(t *testing.T) {
	repo := &mockRepository{
		usuarios: []Usuario{
			{ID: 1, Nombre: "Admin TI", Email: "admin@unsch.edu", Rol: "ADMIN_TI"},
		},
	}
	h := Handler{repo: repo}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/usuarios", nil)
	c, w := testContext()
	c.Request = req

	h.Usuarios(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}