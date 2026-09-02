package catalog

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryInterface interface {
	ListFacultades(ctx context.Context) ([]Facultad, error)
	CreateFacultad(ctx context.Context, nombre string) (*Facultad, error)
	UpdateFacultad(ctx context.Context, id int, nombre string) (*Facultad, error)
	DeleteFacultad(ctx context.Context, id int) error
	ListDepartamentos(ctx context.Context) ([]Departamento, error)
	ListEscuelas(ctx context.Context) ([]Escuela, error)
	CreateEscuela(ctx context.Context, input CreateEscuelaInput) (*Escuela, error)
	UpdateEscuela(ctx context.Context, id int, input UpdateEscuelaInput) (*Escuela, error)
	DeleteEscuela(ctx context.Context, id int) error
	ListAulas(ctx context.Context) ([]Aula, error)
	ListAllAulas(ctx context.Context) ([]Aula, error)
	SetAulaActivo(ctx context.Context, idAula int, activo bool) error
	CreateAula(ctx context.Context, input CreateAulaInput) (*Aula, error)
	UpdateAula(ctx context.Context, id int, input UpdateAulaInput) (*Aula, error)
	DeleteAula(ctx context.Context, id int) error
	ListUsuarios(ctx context.Context) ([]Usuario, error)
	ListPlanesEstudio(ctx context.Context) ([]PlanEstudio, error)
	CreatePlanEstudio(ctx context.Context, input CreatePlanEstudioInput) (*PlanEstudio, error)
	UpdatePlanEstudio(ctx context.Context, id int, input UpdatePlanEstudioInput) (*PlanEstudio, error)
	DeletePlanEstudio(ctx context.Context, id int) error
	ListSeries(ctx context.Context, idEscuela int) ([]Serie, error)
	ListCursos(ctx context.Context, idEscuela int) ([]Curso, error)
	CreateCurso(ctx context.Context, input CreateCursoInput) (*Curso, error)
	UpdateCurso(ctx context.Context, id int, input UpdateCursoInput) (*Curso, error)
	DeleteCurso(ctx context.Context, id int) error
	ListDocentes(ctx context.Context) ([]Docente, error)
	CreateDocente(ctx context.Context, input CreateDocenteInput) (*Docente, error)
	UpdateDocente(ctx context.Context, id int, input UpdateDocenteInput) (*Docente, error)
	DeleteDocente(ctx context.Context, id int) error
	ListPeriodos(ctx context.Context) ([]PeriodoAcademico, error)
	CreatePeriodo(ctx context.Context, input CreatePeriodoInput) (*PeriodoAcademico, error)
	UpdatePeriodo(ctx context.Context, id int, input UpdatePeriodoInput) (*PeriodoAcademico, error)
	DeletePeriodo(ctx context.Context, id int) error
	ListSesionesDepartamento(ctx context.Context) ([]SesionDepartamento, error)
	ListLocales(ctx context.Context) ([]Local, error)
	ListPabellones(ctx context.Context) ([]Pabellon, error)
	ListDistancias(ctx context.Context) ([]Distancia, error)
	ListCargasAcademicas(ctx context.Context, periodo, escuela string) ([]CargaAcademica, error)
	ListGrupos(ctx context.Context) ([]Grupo, error)
	ListHorarios(ctx context.Context) ([]Horario, error)
	ListBloquesHorario(ctx context.Context) ([]BloqueHorario, error)
	ListBitacoraAuditoria(ctx context.Context) ([]BitacoraAuditoria, error)
	CreateHorario(ctx context.Context, input CreateHorarioInput) (*Horario, error)
	ExistsHorario(ctx context.Context, idEscuela, idPeriodo int, idSerie *int, semestre *string) (bool, error)
	GenerateHorarios(ctx context.Context, input GenerateHorariosInput) (*GenerateHorariosResult, error)
	DeleteHorario(ctx context.Context, id int) error
	VerificarConflictoBloque(ctx context.Context, input CreateBloqueInput) ([]ConflictoBloque, error)
	CreateBloque(ctx context.Context, input CreateBloqueInput) (*BloqueHorario, error)
	DeleteBloque(ctx context.Context, id int) error
	GetBloquesByHorario(ctx context.Context, idHorario int) ([]BloqueContexto, error)
	GetGruposParaHorario(ctx context.Context, idEscuela int, idPeriodo int, idSerie *int, semestre *string) ([]GrupoInfo, error)
	CreateDepartamento(ctx context.Context, idFacultad int, nombre string) (*Departamento, error)
	UpdateDepartamento(ctx context.Context, id int, idFacultad int, nombre string) (*Departamento, error)
	DeleteDepartamento(ctx context.Context, id int) error
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return Repository{db: db}
}

func (r Repository) ListFacultades(ctx context.Context) ([]Facultad, error) {
	rows, err := r.db.Query(ctx, `SELECT id_facultad, nombre FROM facultad ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Facultad, 0)
	for rows.Next() {
		var item Facultad
		if err := rows.Scan(&item.ID, &item.Nombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) CreateFacultad(ctx context.Context, nombre string) (*Facultad, error) {
	var id int
	err := r.db.QueryRow(ctx, `INSERT INTO facultad (nombre) VALUES ($1) RETURNING id_facultad`, nombre).Scan(&id)
	if err != nil {
		return nil, err
	}

	var f Facultad
	err = r.db.QueryRow(ctx, `SELECT id_facultad, nombre FROM facultad WHERE id_facultad = $1`, id).Scan(&f.ID, &f.Nombre)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (r Repository) UpdateFacultad(ctx context.Context, id int, nombre string) (*Facultad, error) {
	_, err := r.db.Exec(ctx, `UPDATE facultad SET nombre = $1 WHERE id_facultad = $2`, nombre, id)
	if err != nil {
		return nil, err
	}

	var f Facultad
	err = r.db.QueryRow(ctx, `SELECT id_facultad, nombre FROM facultad WHERE id_facultad = $1`, id).Scan(&f.ID, &f.Nombre)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func (r Repository) DeleteFacultad(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM facultad WHERE id_facultad = $1`, id)
	return err
}

func (r Repository) ListDepartamentos(ctx context.Context) ([]Departamento, error) {
	rows, err := r.db.Query(ctx, `SELECT id_departamento, id_facultad, nombre FROM departamento_academico ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Departamento, 0)
	for rows.Next() {
		var item Departamento
		if err := rows.Scan(&item.ID, &item.IDFacultad, &item.Nombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListEscuelas(ctx context.Context) ([]Escuela, error) {
	rows, err := r.db.Query(ctx, `SELECT id_escuela, id_facultad, id_departamento, nombre FROM escuela_profesional ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Escuela, 0)
	for rows.Next() {
		var item Escuela
		if err := rows.Scan(&item.ID, &item.IDFacultad, &item.IDDepartamento, &item.Nombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) CreateEscuela(ctx context.Context, input CreateEscuelaInput) (*Escuela, error) {
	var id int
	err := r.db.QueryRow(ctx, `
		INSERT INTO escuela_profesional (id_facultad, id_departamento, nombre)
		VALUES ($1, $2, $3)
		RETURNING id_escuela
	`, input.IDFacultad, input.IDDepartamento, input.Nombre).Scan(&id)
	if err != nil {
		return nil, err
	}

	var e Escuela
	err = r.db.QueryRow(ctx, `
		SELECT id_escuela, id_facultad, id_departamento, nombre FROM escuela_profesional WHERE id_escuela = $1
	`, id).Scan(&e.ID, &e.IDFacultad, &e.IDDepartamento, &e.Nombre)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r Repository) UpdateEscuela(ctx context.Context, id int, input UpdateEscuelaInput) (*Escuela, error) {
	query := `UPDATE escuela_profesional SET `
	args := []interface{}{}
	argIdx := 1
	setClauses := []string{}

	if input.IDFacultad != nil {
		setClauses = append(setClauses, fmt.Sprintf("id_facultad = $%d", argIdx))
		args = append(args, *input.IDFacultad)
		argIdx++
	}
	if input.IDDepartamento != nil {
		setClauses = append(setClauses, fmt.Sprintf("id_departamento = $%d", argIdx))
		args = append(args, *input.IDDepartamento)
		argIdx++
	}
	if input.Nombre != "" {
		setClauses = append(setClauses, fmt.Sprintf("nombre = $%d", argIdx))
		args = append(args, input.Nombre)
		argIdx++
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no hay campos para actualizar")
	}

	query += strings.Join(setClauses, ", ") + fmt.Sprintf(" WHERE id_escuela = $%d", argIdx)
	args = append(args, id)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, nil
	}

	var e Escuela
	err = r.db.QueryRow(ctx, `
		SELECT id_escuela, id_facultad, id_departamento, nombre FROM escuela_profesional WHERE id_escuela = $1
	`, id).Scan(&e.ID, &e.IDFacultad, &e.IDDepartamento, &e.Nombre)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r Repository) DeleteEscuela(ctx context.Context, id int) error {
	result, err := r.db.Exec(ctx, `DELETE FROM escuela_profesional WHERE id_escuela = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("escuela no encontrada")
	}
	return nil
}

func (r Repository) ListAulas(ctx context.Context) ([]Aula, error) {
	rows, err := r.db.Query(ctx, `SELECT id_aula, id_pabellon, id_escuela, codigo, tipo::text, aula.aforo, aula.es_compartida, aula.activo FROM aula WHERE aula.activo = true ORDER BY codigo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Aula, 0)
	for rows.Next() {
		var item Aula
		if err := rows.Scan(&item.ID, &item.IDPabellon, &item.IDEscuela, &item.Codigo, &item.Tipo, &item.Aforo, &item.EsCompartida, &item.Activo); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListAllAulas(ctx context.Context) ([]Aula, error) {
	rows, err := r.db.Query(ctx, `SELECT id_aula, id_pabellon, id_escuela, codigo, tipo::text, aula.aforo, aula.es_compartida, aula.activo FROM aula ORDER BY codigo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Aula, 0)
	for rows.Next() {
		var item Aula
		if err := rows.Scan(&item.ID, &item.IDPabellon, &item.IDEscuela, &item.Codigo, &item.Tipo, &item.Aforo, &item.EsCompartida, &item.Activo); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) SetAulaActivo(ctx context.Context, idAula int, activo bool) error {
	_, err := r.db.Exec(ctx, `UPDATE aula SET activo = $1 WHERE id_aula = $2`, activo, idAula)
	return err
}

func (r Repository) CreateAula(ctx context.Context, input CreateAulaInput) (*Aula, error) {
	var id int
	err := r.db.QueryRow(ctx, `
		INSERT INTO aula (id_pabellon, codigo, tipo, aforo, es_compartida, activo)
		VALUES ($1, $2, $3::tipo_aula_enum, $4, $5, true)
		RETURNING id_aula
	`, input.IDPabellon, input.Codigo, input.Tipo, input.Aforo, input.EsCompartida).Scan(&id)
	if err != nil {
		return nil, err
	}

	var aula Aula
	err = r.db.QueryRow(ctx, `
		SELECT id_aula, id_pabellon, id_escuela, codigo, tipo::text, aula.aforo, aula.es_compartida, aula.activo
		FROM aula WHERE id_aula = $1
	`, id).Scan(&aula.ID, &aula.IDPabellon, &aula.IDEscuela, &aula.Codigo, &aula.Tipo, &aula.Aforo, &aula.EsCompartida, &aula.Activo)
	if err != nil {
		return nil, err
	}
	return &aula, nil
}

func (r Repository) UpdateAula(ctx context.Context, id int, input UpdateAulaInput) (*Aula, error) {
	query := `UPDATE aula SET `
	args := []interface{}{}
	argIdx := 1
	setClauses := []string{}

	if input.Codigo != "" {
		setClauses = append(setClauses, fmt.Sprintf("codigo = $%d", argIdx))
		args = append(args, input.Codigo)
		argIdx++
	}
	if input.Tipo != "" {
		setClauses = append(setClauses, fmt.Sprintf("tipo = $%d::tipo_aula_enum", argIdx))
		args = append(args, input.Tipo)
		argIdx++
	}
	if input.Aforo > 0 {
		setClauses = append(setClauses, fmt.Sprintf("aforo = $%d", argIdx))
		args = append(args, input.Aforo)
		argIdx++
	}
	if input.EsCompartida != nil {
		setClauses = append(setClauses, fmt.Sprintf("es_compartida = $%d", argIdx))
		args = append(args, *input.EsCompartida)
		argIdx++
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no hay campos para actualizar")
	}

	query += strings.Join(setClauses, ", ") + fmt.Sprintf(" WHERE id_aula = $%d", argIdx)
	args = append(args, id)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, nil
	}

	var aula Aula
	err = r.db.QueryRow(ctx, `
		SELECT id_aula, id_pabellon, id_escuela, codigo, tipo::text, aula.aforo, aula.es_compartida, aula.activo
		FROM aula WHERE id_aula = $1
	`, id).Scan(&aula.ID, &aula.IDPabellon, &aula.IDEscuela, &aula.Codigo, &aula.Tipo, &aula.Aforo, &aula.EsCompartida, &aula.Activo)
	if err != nil {
		return nil, err
	}
	return &aula, nil
}

func (r Repository) DeleteAula(ctx context.Context, id int) error {
	result, err := r.db.Exec(ctx, `DELETE FROM aula WHERE id_aula = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("aula no encontrada")
	}
	return nil
}

func (r Repository) ListUsuarios(ctx context.Context) ([]Usuario, error) {
	rows, err := r.db.Query(ctx, `SELECT id_usuario, nombre, email, rol::text FROM usuario ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Usuario, 0)
	for rows.Next() {
		var item Usuario
		if err := rows.Scan(&item.ID, &item.Nombre, &item.Email, &item.Rol); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r Repository) ListPlanesEstudio(ctx context.Context) ([]PlanEstudio, error) {
	rows, err := r.db.Query(ctx, `SELECT id_plan, id_escuela, codigo_plan, nombre FROM plan_estudio ORDER BY codigo_plan`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]PlanEstudio, 0)
	for rows.Next() {
		var item PlanEstudio
		if err := rows.Scan(&item.ID, &item.IDEscuela, &item.Codigo, &item.Nombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) CreatePlanEstudio(ctx context.Context, input CreatePlanEstudioInput) (*PlanEstudio, error) {
	var id int
	err := r.db.QueryRow(ctx, `
		INSERT INTO plan_estudio (id_escuela, codigo_plan, nombre)
		VALUES ($1, $2, $3)
		RETURNING id_plan
	`, input.IDEscuela, input.Codigo, input.Nombre).Scan(&id)
	if err != nil {
		return nil, err
	}

	var p PlanEstudio
	err = r.db.QueryRow(ctx, `
		SELECT id_plan, id_escuela, codigo_plan, nombre FROM plan_estudio WHERE id_plan = $1
	`, id).Scan(&p.ID, &p.IDEscuela, &p.Codigo, &p.Nombre)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r Repository) UpdatePlanEstudio(ctx context.Context, id int, input UpdatePlanEstudioInput) (*PlanEstudio, error) {
	query := `UPDATE plan_estudio SET `
	args := []interface{}{}
	argIdx := 1
	setClauses := []string{}

	if input.IDEscuela != nil {
		setClauses = append(setClauses, fmt.Sprintf("id_escuela = $%d", argIdx))
		args = append(args, *input.IDEscuela)
		argIdx++
	}
	if input.Codigo != "" {
		setClauses = append(setClauses, fmt.Sprintf("codigo_plan = $%d", argIdx))
		args = append(args, input.Codigo)
		argIdx++
	}
	if input.Nombre != "" {
		setClauses = append(setClauses, fmt.Sprintf("nombre = $%d", argIdx))
		args = append(args, input.Nombre)
		argIdx++
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no hay campos para actualizar")
	}

	query += strings.Join(setClauses, ", ") + fmt.Sprintf(" WHERE id_plan = $%d", argIdx)
	args = append(args, id)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, nil
	}

	var p PlanEstudio
	err = r.db.QueryRow(ctx, `
		SELECT id_plan, id_escuela, codigo_plan, nombre FROM plan_estudio WHERE id_plan = $1
	`, id).Scan(&p.ID, &p.IDEscuela, &p.Codigo, &p.Nombre)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r Repository) DeletePlanEstudio(ctx context.Context, id int) error {
	result, err := r.db.Exec(ctx, `DELETE FROM plan_estudio WHERE id_plan = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("plan de estudio no encontrado")
	}
	return nil
}

func (r Repository) ListSeries(ctx context.Context, idEscuela int) ([]Serie, error) {
	query := `SELECT s.id_serie, s.id_plan, s.numero_ciclo, s.subciclo, p.codigo_plan
		FROM serie s
		JOIN plan_estudio p ON p.id_plan = s.id_plan`
	var args []interface{}
	if idEscuela > 0 {
		query += ` WHERE p.id_escuela = $1`
		args = append(args, idEscuela)
	}
	query += ` ORDER BY p.id_escuela, s.numero_ciclo, s.subciclo`

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Serie, 0)
	for rows.Next() {
		var item Serie
		if err := rows.Scan(&item.ID, &item.IDPlan, &item.NumeroCiclo, &item.Subciclo, &item.CodigoPlan); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListCursos(ctx context.Context, idEscuela int) ([]Curso, error) {
	query := `
		SELECT c.id_curso, c.id_serie, c.codigo, c.nombre, c.creditos, c.horas_teoria, c.horas_practica,
		       e.id_escuela, e.nombre as escuela_nombre
		FROM curso c
		JOIN serie s ON s.id_serie = c.id_serie
		JOIN plan_estudio p ON p.id_plan = s.id_plan
		JOIN escuela_profesional e ON e.id_escuela = p.id_escuela
	`
	var args []interface{}
	if idEscuela > 0 {
		query += ` WHERE e.id_escuela = $1`
		args = append(args, idEscuela)
	}
	query += ` ORDER BY e.nombre, c.codigo`

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Curso, 0)
	for rows.Next() {
		var item Curso
		if err := rows.Scan(&item.ID, &item.IDSerie, &item.Codigo, &item.Nombre, &item.Creditos, &item.HorasTeoria, &item.HorasPractica, &item.IDEscuela, &item.EscuelaNombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) CreateCurso(ctx context.Context, input CreateCursoInput) (*Curso, error) {
	var id int
	err := r.db.QueryRow(ctx, `
		INSERT INTO curso (id_serie, codigo, nombre, creditos, horas_teoria, horas_practica)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_curso
	`, input.IDSerie, input.Codigo, input.Nombre, input.Creditos, input.HorasTeoria, input.HorasPractica).Scan(&id)
	if err != nil {
		return nil, err
	}

	var c Curso
	err = r.db.QueryRow(ctx, `
		SELECT c.id_curso, c.id_serie, c.codigo, c.nombre, c.creditos, c.horas_teoria, c.horas_practica,
		       e.id_escuela, e.nombre as escuela_nombre
		FROM curso c
		JOIN serie s ON s.id_serie = c.id_serie
		JOIN plan_estudio p ON p.id_plan = s.id_plan
		JOIN escuela_profesional e ON e.id_escuela = p.id_escuela
		WHERE c.id_curso = $1
	`, id).Scan(&c.ID, &c.IDSerie, &c.Codigo, &c.Nombre, &c.Creditos, &c.HorasTeoria, &c.HorasPractica, &c.IDEscuela, &c.EscuelaNombre)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r Repository) UpdateCurso(ctx context.Context, id int, input UpdateCursoInput) (*Curso, error) {
	query := `UPDATE curso SET `
	args := []interface{}{}
	argIdx := 1
	setClauses := []string{}

	if input.IDSerie != nil {
		setClauses = append(setClauses, fmt.Sprintf("id_serie = $%d", argIdx))
		args = append(args, *input.IDSerie)
		argIdx++
	}
	if input.Codigo != "" {
		setClauses = append(setClauses, fmt.Sprintf("codigo = $%d", argIdx))
		args = append(args, input.Codigo)
		argIdx++
	}
	if input.Nombre != "" {
		setClauses = append(setClauses, fmt.Sprintf("nombre = $%d", argIdx))
		args = append(args, input.Nombre)
		argIdx++
	}
	if input.Creditos != nil {
		setClauses = append(setClauses, fmt.Sprintf("creditos = $%d", argIdx))
		args = append(args, *input.Creditos)
		argIdx++
	}
	if input.HorasTeoria != nil {
		setClauses = append(setClauses, fmt.Sprintf("horas_teoria = $%d", argIdx))
		args = append(args, *input.HorasTeoria)
		argIdx++
	}
	if input.HorasPractica != nil {
		setClauses = append(setClauses, fmt.Sprintf("horas_practica = $%d", argIdx))
		args = append(args, *input.HorasPractica)
		argIdx++
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no hay campos para actualizar")
	}

	query += strings.Join(setClauses, ", ") + fmt.Sprintf(" WHERE id_curso = $%d", argIdx)
	args = append(args, id)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, nil
	}

	var c Curso
	err = r.db.QueryRow(ctx, `
		SELECT c.id_curso, c.id_serie, c.codigo, c.nombre, c.creditos, c.horas_teoria, c.horas_practica,
		       e.id_escuela, e.nombre as escuela_nombre
		FROM curso c
		JOIN serie s ON s.id_serie = c.id_serie
		JOIN plan_estudio p ON p.id_plan = s.id_plan
		JOIN escuela_profesional e ON e.id_escuela = p.id_escuela
		WHERE c.id_curso = $1
	`, id).Scan(&c.ID, &c.IDSerie, &c.Codigo, &c.Nombre, &c.Creditos, &c.HorasTeoria, &c.HorasPractica, &c.IDEscuela, &c.EscuelaNombre)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r Repository) DeleteCurso(ctx context.Context, id int) error {
	result, err := r.db.Exec(ctx, `DELETE FROM curso WHERE id_curso = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("curso no encontrado")
	}
	return nil
}

func (r Repository) ListDocentes(ctx context.Context) ([]Docente, error) {
	rows, err := r.db.Query(ctx, `SELECT id_docente, id_departamento, codigo_plaza, nombres, apellidos, email FROM docente ORDER BY apellidos, nombres`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Docente, 0)
	for rows.Next() {
		var item Docente
		if err := rows.Scan(&item.ID, &item.IDDepartamento, &item.CodigoPlaza, &item.Nombres, &item.Apellidos, &item.Email); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) CreateDocente(ctx context.Context, input CreateDocenteInput) (*Docente, error) {
	var id int
	err := r.db.QueryRow(ctx, `
		INSERT INTO docente (id_departamento, codigo_plaza, nombres, apellidos, email)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id_docente
	`, input.IDDepartamento, input.CodigoPlaza, input.Nombres, input.Apellidos, input.Email).Scan(&id)
	if err != nil {
		return nil, err
	}

	var d Docente
	err = r.db.QueryRow(ctx, `
		SELECT id_docente, id_departamento, codigo_plaza, nombres, apellidos, email FROM docente WHERE id_docente = $1
	`, id).Scan(&d.ID, &d.IDDepartamento, &d.CodigoPlaza, &d.Nombres, &d.Apellidos, &d.Email)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r Repository) UpdateDocente(ctx context.Context, id int, input UpdateDocenteInput) (*Docente, error) {
	query := `UPDATE docente SET `
	args := []interface{}{}
	argIdx := 1
	setClauses := []string{}

	if input.IDDepartamento != nil {
		setClauses = append(setClauses, fmt.Sprintf("id_departamento = $%d", argIdx))
		args = append(args, *input.IDDepartamento)
		argIdx++
	}
	if input.CodigoPlaza != "" {
		setClauses = append(setClauses, fmt.Sprintf("codigo_plaza = $%d", argIdx))
		args = append(args, input.CodigoPlaza)
		argIdx++
	}
	if input.Nombres != "" {
		setClauses = append(setClauses, fmt.Sprintf("nombres = $%d", argIdx))
		args = append(args, input.Nombres)
		argIdx++
	}
	if input.Apellidos != "" {
		setClauses = append(setClauses, fmt.Sprintf("apellidos = $%d", argIdx))
		args = append(args, input.Apellidos)
		argIdx++
	}
	if input.Email != "" {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argIdx))
		args = append(args, input.Email)
		argIdx++
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no hay campos para actualizar")
	}

	query += strings.Join(setClauses, ", ") + fmt.Sprintf(" WHERE id_docente = $%d", argIdx)
	args = append(args, id)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, nil
	}

	var d Docente
	err = r.db.QueryRow(ctx, `
		SELECT id_docente, id_departamento, codigo_plaza, nombres, apellidos, email FROM docente WHERE id_docente = $1
	`, id).Scan(&d.ID, &d.IDDepartamento, &d.CodigoPlaza, &d.Nombres, &d.Apellidos, &d.Email)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r Repository) DeleteDocente(ctx context.Context, id int) error {
	result, err := r.db.Exec(ctx, `DELETE FROM docente WHERE id_docente = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("docente no encontrado")
	}
	return nil
}

func (r Repository) ListPeriodos(ctx context.Context) ([]PeriodoAcademico, error) {
	rows, err := r.db.Query(ctx, `SELECT id_periodo, codigo, activo FROM periodo_academico ORDER BY id_periodo DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]PeriodoAcademico, 0)
	for rows.Next() {
		var item PeriodoAcademico
		if err := rows.Scan(&item.ID, &item.Codigo, &item.Activo); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) CreatePeriodo(ctx context.Context, input CreatePeriodoInput) (*PeriodoAcademico, error) {
	var id int
	err := r.db.QueryRow(ctx, `
		INSERT INTO periodo_academico (codigo, activo)
		VALUES ($1, $2)
		RETURNING id_periodo
	`, input.Codigo, input.Activo).Scan(&id)
	if err != nil {
		return nil, err
	}

	var p PeriodoAcademico
	err = r.db.QueryRow(ctx, `
		SELECT id_periodo, codigo, activo FROM periodo_academico WHERE id_periodo = $1
	`, id).Scan(&p.ID, &p.Codigo, &p.Activo)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r Repository) UpdatePeriodo(ctx context.Context, id int, input UpdatePeriodoInput) (*PeriodoAcademico, error) {
	query := `UPDATE periodo_academico SET `
	args := []interface{}{}
	argIdx := 1
	setClauses := []string{}

	if input.Codigo != nil {
		setClauses = append(setClauses, fmt.Sprintf("codigo = $%d", argIdx))
		args = append(args, *input.Codigo)
		argIdx++
	}
	if input.Activo != nil {
		setClauses = append(setClauses, fmt.Sprintf("activo = $%d", argIdx))
		args = append(args, *input.Activo)
		argIdx++
	}

	if len(setClauses) == 0 {
		return nil, fmt.Errorf("no hay campos para actualizar")
	}

	query += strings.Join(setClauses, ", ") + fmt.Sprintf(" WHERE id_periodo = $%d", argIdx)
	args = append(args, id)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected() == 0 {
		return nil, nil
	}

	var p PeriodoAcademico
	err = r.db.QueryRow(ctx, `
		SELECT id_periodo, codigo, activo FROM periodo_academico WHERE id_periodo = $1
	`, id).Scan(&p.ID, &p.Codigo, &p.Activo)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r Repository) DeletePeriodo(ctx context.Context, id int) error {
	result, err := r.db.Exec(ctx, `DELETE FROM periodo_academico WHERE id_periodo = $1`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("periodo no encontrado")
	}
	return nil
}

func (r Repository) ListSesionesDepartamento(ctx context.Context) ([]SesionDepartamento, error) {
	rows, err := r.db.Query(ctx, `SELECT id_sesion, id_departamento, id_periodo, dia_semana, hora_inicio, hora_fin FROM sesion_departamento ORDER BY id_periodo, id_departamento`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]SesionDepartamento, 0)
	for rows.Next() {
		var item SesionDepartamento
		if err := rows.Scan(&item.ID, &item.IDDepartamento, &item.IDPeriodo, &item.DiaSemana, &item.HoraInicio, &item.HoraFin); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListLocales(ctx context.Context) ([]Local, error) {
	rows, err := r.db.Query(ctx, `SELECT id_local, nombre FROM local ORDER BY nombre`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Local, 0)
	for rows.Next() {
		var item Local
		if err := rows.Scan(&item.ID, &item.Nombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListPabellones(ctx context.Context) ([]Pabellon, error) {
	rows, err := r.db.Query(ctx, `SELECT id_pabellon, id_local, codigo, nombre FROM pabellon ORDER BY codigo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Pabellon, 0)
	for rows.Next() {
		var item Pabellon
		if err := rows.Scan(&item.ID, &item.IDLocal, &item.Codigo, &item.Nombre); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListDistancias(ctx context.Context) ([]Distancia, error) {
	rows, err := r.db.Query(ctx, `SELECT id_pabellon_origen, id_pabellon_destino, tiempo_minutos FROM matriz_distancia ORDER BY id_pabellon_origen, id_pabellon_destino`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Distancia, 0)
	for rows.Next() {
		var item Distancia
		if err := rows.Scan(&item.DesdeID, &item.HastaID, &item.Minutos); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListCargasAcademicas(ctx context.Context, periodo, escuela string) ([]CargaAcademica, error) {
	query := `SELECT id_carga, id_curso, id_periodo, id_escuela, estado, fecha_aprobacion FROM carga_academica WHERE 1=1`
	var args []interface{}
	argNum := 1

	if periodo != "" {
		query += ` AND id_periodo = $` + strconv.Itoa(argNum)
		args = append(args, periodo)
		argNum++
	}
	if escuela != "" {
		escuelaInt, err := strconv.Atoi(escuela)
		if err == nil {
			query += ` AND id_escuela = $` + strconv.Itoa(argNum)
			args = append(args, escuelaInt)
		}
	}

	query += ` ORDER BY id_periodo, id_escuela`

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]CargaAcademica, 0)
	for rows.Next() {
		var item CargaAcademica
		if err := rows.Scan(&item.IDCarga, &item.IDCurso, &item.IDPeriodo, &item.IDEscuela, &item.Estado, &item.FechaAprobacion); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListGrupos(ctx context.Context) ([]Grupo, error) {
	rows, err := r.db.Query(ctx, `SELECT id_grupo, id_carga, id_docente, id_grupo_teoria_ref, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados, matriculados_reales FROM grupo ORDER BY id_carga, codigo_grupo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Grupo, 0)
	for rows.Next() {
		var item Grupo
		if err := rows.Scan(&item.ID, &item.IDCarga, &item.IDDocente, &item.IDGrupoTeoriaRef, &item.CodigoGrupo, &item.TipoComponente, &item.EsNuevaNecesidad, &item.MatriculadosProyectados, &item.MatriculadosReales); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListHorarios(ctx context.Context) ([]Horario, error) {
	rows, err := r.db.Query(ctx, `SELECT id_horario, id_escuela, id_periodo, id_serie, semestre, estado::text, version_reajuste, fecha_actualizacion FROM horario ORDER BY id_escuela, id_periodo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Horario, 0)
	for rows.Next() {
		var item Horario
		if err := rows.Scan(&item.ID, &item.IDEscuela, &item.IDPeriodo, &item.IDSerie, &item.Semestre, &item.Estado, &item.VersionReajuste, &item.FechaActualizacion); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListBloquesHorario(ctx context.Context) ([]BloqueHorario, error) {
	rows, err := r.db.Query(ctx, `SELECT id_bloque, id_horario, id_grupo, id_aula, id_docente, dia_semana, slot_inicio, slot_fin FROM bloque_horario ORDER BY id_horario, dia_semana, slot_inicio`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]BloqueHorario, 0)
	for rows.Next() {
		var item BloqueHorario
		if err := rows.Scan(&item.ID, &item.IDHorario, &item.IDGrupo, &item.IDAula, &item.IDDocente, &item.DiaSemana, &item.SlotInicio, &item.SlotFin); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListBitacoraAuditoria(ctx context.Context) ([]BitacoraAuditoria, error) {
	rows, err := r.db.Query(ctx, `SELECT id_log, id_horario, id_usuario, accion, motivo_justificacion, version_resultante, fecha_hora FROM bitacora_auditoria ORDER BY fecha_hora DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]BitacoraAuditoria, 0)
	for rows.Next() {
		var item BitacoraAuditoria
		if err := rows.Scan(&item.ID, &item.IDHorario, &item.IDUsuario, &item.Accion, &item.MotivoJustificacion, &item.VersionResultante, &item.FechaHora); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ExistsHorario(ctx context.Context, idEscuela, idPeriodo int, idSerie *int, semestre *string) (bool, error) {
	var count int
	err := r.db.QueryRow(ctx, `
		SELECT COUNT(*) FROM horario
		WHERE id_escuela = $1 AND id_periodo = $2
		  AND ($3::int IS NULL AND id_serie IS NULL OR id_serie = $3)
		  AND ($4::text IS NULL AND semestre IS NULL OR semestre = $4)
	`, idEscuela, idPeriodo, idSerie, semestre).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r Repository) CreateHorario(ctx context.Context, input CreateHorarioInput) (*Horario, error) {
	var idHorario int
	err := r.db.QueryRow(ctx, `
		INSERT INTO horario (id_escuela, id_periodo, id_serie, semestre, estado, version_reajuste)
		VALUES ($1, $2, $3, $4, 'BORRADOR', 0)
		ON CONFLICT (id_escuela, id_periodo, id_serie, semestre) DO UPDATE SET id_escuela = EXCLUDED.id_escuela
		RETURNING id_horario
	`, input.IDEscuela, input.IDPeriodo, input.IDSerie, input.Semestre).Scan(&idHorario)
	if err != nil {
		return nil, err
	}

	var h Horario
	err = r.db.QueryRow(ctx, `
		SELECT id_horario, id_escuela, id_periodo, id_serie, semestre, estado::text, version_reajuste, fecha_actualizacion
		FROM horario WHERE id_horario = $1
	`, idHorario).Scan(&h.ID, &h.IDEscuela, &h.IDPeriodo, &h.IDSerie, &h.Semestre, &h.Estado, &h.VersionReajuste, &h.FechaActualizacion)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

func (r Repository) GenerateHorarios(ctx context.Context, input GenerateHorariosInput) (*GenerateHorariosResult, error) {
	subciclo := 1
	if input.Semestre == "II" {
		subciclo = 2
	}

	rows, err := r.db.Query(ctx, `
		SELECT s.id_serie, s.numero_ciclo
		FROM serie s
		JOIN plan_estudio p ON p.id_plan = s.id_plan
		WHERE p.id_escuela = $1 AND s.subciclo = $2
		ORDER BY s.numero_ciclo
	`, input.IDEscuela, subciclo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var series []struct {
		ID           int
		NumeroCiclo int
	}
	for rows.Next() {
		var s struct {
			ID           int
			NumeroCiclo int
		}
		if err := rows.Scan(&s.ID, &s.NumeroCiclo); err != nil {
			continue
		}
		series = append(series, s)
	}

	result := &GenerateHorariosResult{
		Horarios: []Horario{},
	}

	semestre := input.Semestre

	for _, serie := range series {
		idSerie := serie.ID

		existe, err := r.ExistsHorario(ctx, input.IDEscuela, input.IDPeriodo, &idSerie, &semestre)
		if err != nil {
			continue
		}

		if existe {
			result.Existentes++
			var h Horario
			err = r.db.QueryRow(ctx, `
				SELECT id_horario, id_escuela, id_periodo, id_serie, semestre, estado::text, version_reajuste, fecha_actualizacion
				FROM horario
				WHERE id_escuela = $1 AND id_periodo = $2 AND id_serie = $3 AND semestre = $4
			`, input.IDEscuela, input.IDPeriodo, idSerie, semestre).Scan(&h.ID, &h.IDEscuela, &h.IDPeriodo, &h.IDSerie, &h.Semestre, &h.Estado, &h.VersionReajuste, &h.FechaActualizacion)
			if err == nil {
				result.Horarios = append(result.Horarios, h)
			}
			continue
		}

		h, err := r.CreateHorario(ctx, CreateHorarioInput{
			IDEscuela: input.IDEscuela,
			IDPeriodo: input.IDPeriodo,
			IDSerie:   &idSerie,
			Semestre:  &semestre,
		})
		if err != nil {
			continue
		}

		result.Creados++
		result.Horarios = append(result.Horarios, *h)
	}

	return result, nil
}

func (r Repository) DeleteHorario(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM bloque_horario WHERE id_horario = $1`, id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, `DELETE FROM horario WHERE id_horario = $1`, id)
	return err
}

type ConflictoBloque struct {
	Tipo       string `json:"tipo"`
	Mensaje    string `json:"mensaje"`
	Detalle    string `json:"detalle,omitempty"`
}

func (r Repository) VerificarConflictoBloque(ctx context.Context, input CreateBloqueInput) ([]ConflictoBloque, error) {
	var conflictos []ConflictoBloque

	if input.IDDocente != nil {
		var idPeriodo int
		err := r.db.QueryRow(ctx, `SELECT id_periodo FROM horario WHERE id_horario = $1`, input.IDHorario).Scan(&idPeriodo)
		if err != nil {
			return nil, err
		}

		var count int
		err = r.db.QueryRow(ctx, `
			SELECT COUNT(*) FROM bloque_horario bh
			JOIN grupo g ON g.id_grupo = bh.id_grupo
			JOIN carga_academica ca ON ca.id_carga = g.id_carga
			JOIN horario h ON h.id_horario = bh.id_horario
			WHERE g.id_docente = $1 AND h.id_periodo = $2
			  AND bh.dia_semana = $3
			  AND bh.slot_inicio < $5 AND bh.slot_fin > $4
		`, *input.IDDocente, idPeriodo, input.DiaSemana, input.SlotInicio, input.SlotFin).Scan(&count)
		if err != nil {
			return nil, err
		}
		if count > 0 {
			var nombreDocente string
			r.db.QueryRow(ctx, `SELECT nombres || ' ' || apellidos FROM docente WHERE id_docente = $1`, *input.IDDocente).Scan(&nombreDocente)
			conflictos = append(conflictos, ConflictoBloque{
				Tipo:    "CRUCE_DOCENTE",
				Mensaje:  "El docente " + nombreDocente + " ya tiene una clase en ese horario",
			})
		}
	}

	var countAula int
	err := r.db.QueryRow(ctx, `
		SELECT COUNT(*) FROM bloque_horario
		WHERE id_horario = $1 AND dia_semana = $2
		  AND slot_inicio < $4 AND slot_fin > $3
		  AND id_aula = $5
	`, input.IDHorario, input.DiaSemana, input.SlotInicio, input.SlotFin, input.IDAula).Scan(&countAula)
	if err != nil {
		return nil, err
	}
	if countAula > 0 {
		var codigoAula string
		r.db.QueryRow(ctx, `SELECT codigo FROM aula WHERE id_aula = $1`, input.IDAula).Scan(&codigoAula)
		conflictos = append(conflictos, ConflictoBloque{
			Tipo:    "CRUCE_AULA",
			Mensaje:  "El aula " + codigoAula + " ya está ocupada en ese horario",
		})
	}

	return conflictos, nil
}

func (r Repository) CreateBloque(ctx context.Context, input CreateBloqueInput) (*BloqueHorario, error) {
	var idBloque int
	err := r.db.QueryRow(ctx, `
		INSERT INTO bloque_horario (id_horario, id_grupo, id_aula, id_docente, dia_semana, slot_inicio, slot_fin)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_bloque
	`, input.IDHorario, input.IDGrupo, input.IDAula, input.IDDocente, input.DiaSemana, input.SlotInicio, input.SlotFin).Scan(&idBloque)
	if err != nil {
		return nil, err
	}

	var b BloqueHorario
	err = r.db.QueryRow(ctx, `
		SELECT id_bloque, id_horario, id_grupo, id_aula, id_docente, dia_semana, slot_inicio, slot_fin
		FROM bloque_horario WHERE id_bloque = $1
	`, idBloque).Scan(&b.ID, &b.IDHorario, &b.IDGrupo, &b.IDAula, &b.IDDocente, &b.DiaSemana, &b.SlotInicio, &b.SlotFin)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r Repository) DeleteBloque(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM bloque_horario WHERE id_bloque = $1`, id)
	return err
}

func (r Repository) GetBloquesByHorario(ctx context.Context, idHorario int) ([]BloqueContexto, error) {
	rows, err := r.db.Query(ctx, `
		SELECT bh.id_bloque, bh.id_horario, bh.id_grupo, g.id_carga, e.id_escuela,
		       bh.id_aula, a.id_pabellon, bh.id_docente, COALESCE(d.id_departamento, 0),
		       a.codigo as codigo_aula,
		       COALESCE(d.nombres || ' ' || d.apellidos, '') as nombre_docente,
		       g.codigo_grupo, g.tipo_componente::text,
		       h.estado::text,
		       bh.dia_semana, bh.slot_inicio, bh.slot_fin,
		       e.nombre as escuela_nombre, c.codigo as curso_codigo, c.nombre as curso_nombre,
		       c.horas_teoria, c.horas_practica, c.id_curso
		FROM bloque_horario bh
		JOIN horario h ON h.id_horario = bh.id_horario
		JOIN grupo g ON g.id_grupo = bh.id_grupo
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		JOIN curso c ON c.id_curso = ca.id_curso
		JOIN escuela_profesional e ON e.id_escuela = h.id_escuela
		JOIN aula a ON a.id_aula = bh.id_aula
		LEFT JOIN docente d ON d.id_docente = bh.id_docente
		WHERE bh.id_horario = $1
		ORDER BY bh.dia_semana, bh.slot_inicio
	`, idHorario)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []BloqueContexto
	for rows.Next() {
		var b BloqueContexto
		if err := rows.Scan(&b.ID, &b.IDHorario, &b.IDGrupo, &b.IDCarga, &b.IDEscuela,
			&b.IDAula, &b.IDAulaPabellon, &b.IDDocente, &b.IDDocenteDepto,
			&b.CodigoAula, &b.NombreDocente,
			&b.CodigoGrupo, &b.TipoComponente, &b.EstadoHorario,
			&b.DiaSemana, &b.SlotInicio, &b.SlotFin,
			&b.NombreEscuela, &b.CodigoCurso, &b.NombreCurso,
			&b.HorasTeoria, &b.HorasPractica, &b.IDCurso); err != nil {
			return nil, err
		}
		items = append(items, b)
	}
	return items, rows.Err()
}

func (r Repository) GetGruposParaHorario(ctx context.Context, idEscuela int, idPeriodo int, idSerie *int, semestre *string) ([]GrupoInfo, error) {
	query := `
		SELECT g.id_grupo, g.id_carga, g.codigo_grupo, g.tipo_componente::text,
		       g.id_docente, COALESCE(d.nombres || ' ' || d.apellidos, '') as docente_nombre,
		       c.codigo, c.nombre, c.horas_teoria, c.horas_practica
		FROM grupo g
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		JOIN curso c ON c.id_curso = ca.id_curso
		LEFT JOIN docente d ON d.id_docente = g.id_docente
		WHERE ca.id_escuela = $1 AND ca.id_periodo = $2 AND ca.estado = 'AUTORIZADO'
	`
	args := []interface{}{idEscuela, idPeriodo}
	argIdx := 3

	if idSerie != nil {
		query += fmt.Sprintf(" AND c.id_serie = $%d", argIdx)
		args = append(args, *idSerie)
		argIdx++
	}
	if semestre != nil && *semestre != "" {
		if *semestre == "I" {
			query += " AND RIGHT(c.codigo, 1) ~ '[13579]'"
		} else if *semestre == "II" {
			query += " AND RIGHT(c.codigo, 1) ~ '[02468]'"
		}
	}

	query += " ORDER BY c.codigo, g.codigo_grupo"

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []GrupoInfo
	for rows.Next() {
		var g GrupoInfo
		if err := rows.Scan(&g.ID, &g.IDCarga, &g.CodigoGrupo, &g.TipoComponente,
			&g.IDDocente, &g.DocenteNombre, &g.CodigoCurso, &g.NombreCurso,
			&g.HorasTeoria, &g.HorasPractica); err != nil {
			return nil, err
		}
		items = append(items, g)
	}
	return items, rows.Err()
}

type GrupoInfo struct {
	ID             int     `json:"id_grupo"`
	IDCarga       int     `json:"id_carga"`
	CodigoGrupo   string  `json:"codigo_grupo"`
	TipoComponente string  `json:"tipo_componente"`
	IDDocente     *int    `json:"id_docente"`
	DocenteNombre string  `json:"docente_nombre"`
	CodigoCurso   string  `json:"codigo_curso"`
	NombreCurso   string  `json:"nombre_curso"`
	HorasTeoria   int     `json:"horas_teoria"`
	HorasPractica int     `json:"horas_practica"`
}

func (r Repository) CreateDepartamento(ctx context.Context, idFacultad int, nombre string) (*Departamento, error) {
	var id int
	err := r.db.QueryRow(ctx, `
		INSERT INTO departamento_academico (id_facultad, nombre)
		VALUES ($1, $2)
		RETURNING id_departamento
	`, idFacultad, nombre).Scan(&id)
	if err != nil {
		return nil, err
	}

	var d Departamento
	err = r.db.QueryRow(ctx, `
		SELECT id_departamento, id_facultad, nombre FROM departamento_academico WHERE id_departamento = $1
	`, id).Scan(&d.ID, &d.IDFacultad, &d.Nombre)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r Repository) UpdateDepartamento(ctx context.Context, id int, idFacultad int, nombre string) (*Departamento, error) {
	_, err := r.db.Exec(ctx, `
		UPDATE departamento_academico SET id_facultad = $1, nombre = $2 WHERE id_departamento = $3
	`, idFacultad, nombre, id)
	if err != nil {
		return nil, err
	}

	var d Departamento
	err = r.db.QueryRow(ctx, `
		SELECT id_departamento, id_facultad, nombre FROM departamento_academico WHERE id_departamento = $1
	`, id).Scan(&d.ID, &d.IDFacultad, &d.Nombre)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r Repository) DeleteDepartamento(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM departamento_academico WHERE id_departamento = $1`, id)
	return err
}
