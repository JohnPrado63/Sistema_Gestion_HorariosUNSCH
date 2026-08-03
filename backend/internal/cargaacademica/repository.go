package cargaacademica

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetCargaByID(ctx context.Context, idCarga int) (*CargaAcademica, error) {
	var c CargaAcademica
	var escuelaNombre string
	var cursoInfo CursoInfo
	err := r.db.QueryRow(ctx, `
		SELECT ca.id_carga, ca.id_curso, ca.id_periodo, ca.id_escuela, ca.estado, COALESCE(ca.fecha_aprobacion::text, ''),
		       c.codigo, c.nombre, c.horas_teoria, c.horas_practica, c.creditos,
		       e.nombre
		FROM carga_academica ca
		JOIN curso c ON c.id_curso = ca.id_curso
		JOIN escuela_profesional e ON e.id_escuela = ca.id_escuela
		WHERE ca.id_carga = $1
	`, idCarga).Scan(
		&c.IDCarga, &c.IDCurso, &c.IDPeriodo, &c.IDEscuela, &c.Estado, &c.FechaAprobacion,
		&cursoInfo.Codigo, &cursoInfo.Nombre, &cursoInfo.HorasTeoria, &cursoInfo.HorasPractica, &cursoInfo.Creditos,
		&escuelaNombre,
	)
	if err != nil {
		return nil, err
	}
	cursoInfo.ID = c.IDCurso
	c.Curso = &cursoInfo
	c.Escuela = escuelaNombre
	return &c, nil
}

func (r *Repository) ListCargasByPeriodo(ctx context.Context, idPeriodo int) ([]CargaAcademica, error) {
	rows, err := r.db.Query(ctx, `
		SELECT ca.id_carga, ca.id_curso, ca.id_periodo, ca.id_escuela, ca.estado, COALESCE(ca.fecha_aprobacion::text, ''),
		       c.codigo, c.nombre, c.horas_teoria, c.horas_practica, c.creditos,
		       e.nombre
		FROM carga_academica ca
		JOIN curso c ON c.id_curso = ca.id_curso
		JOIN escuela_profesional e ON e.id_escuela = ca.id_escuela
		WHERE ca.id_periodo = $1
		ORDER BY e.nombre, c.codigo
	`, idPeriodo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cargas []CargaAcademica
	for rows.Next() {
		var c CargaAcademica
		var escuelaNombre string
		var cursoInfo CursoInfo
		if err := rows.Scan(
			&c.IDCarga, &c.IDCurso, &c.IDPeriodo, &c.IDEscuela, &c.Estado, &c.FechaAprobacion,
			&cursoInfo.Codigo, &cursoInfo.Nombre, &cursoInfo.HorasTeoria, &cursoInfo.HorasPractica, &cursoInfo.Creditos,
			&escuelaNombre,
		); err != nil {
			return nil, err
		}
		cursoInfo.ID = c.IDCurso
		c.Curso = &cursoInfo
		c.Escuela = escuelaNombre
		cargas = append(cargas, c)
	}
	return cargas, rows.Err()
}

func (r *Repository) ListCargasByEscuela(ctx context.Context, idEscuela, idPeriodo int) ([]CargaAcademica, error) {
	rows, err := r.db.Query(ctx, `
		SELECT ca.id_carga, ca.id_curso, ca.id_periodo, ca.id_escuela, ca.estado, COALESCE(ca.fecha_aprobacion::text, ''),
		       c.codigo, c.nombre, c.horas_teoria, c.horas_practica, c.creditos,
		       e.nombre
		FROM carga_academica ca
		JOIN curso c ON c.id_curso = ca.id_curso
		JOIN escuela_profesional e ON e.id_escuela = ca.id_escuela
		WHERE ca.id_escuela = $1 AND ca.id_periodo = $2
		ORDER BY c.codigo
	`, idEscuela, idPeriodo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cargas []CargaAcademica
	for rows.Next() {
		var c CargaAcademica
		var escuelaNombre string
		var cursoInfo CursoInfo
		if err := rows.Scan(
			&c.IDCarga, &c.IDCurso, &c.IDPeriodo, &c.IDEscuela, &c.Estado, &c.FechaAprobacion,
			&cursoInfo.Codigo, &cursoInfo.Nombre, &cursoInfo.HorasTeoria, &cursoInfo.HorasPractica, &cursoInfo.Creditos,
			&escuelaNombre,
		); err != nil {
			return nil, err
		}
		cursoInfo.ID = c.IDCurso
		c.Curso = &cursoInfo
		c.Escuela = escuelaNombre
		cargas = append(cargas, c)
	}
	return cargas, rows.Err()
}

func (r *Repository) CreateCarga(ctx context.Context, input CreateCargaInput) (*CargaAcademica, error) {
	var idCarga int
	err := r.db.QueryRow(ctx, `
		INSERT INTO carga_academica (id_curso, id_periodo, id_escuela, estado)
		VALUES ($1, $2, $3, 'BORRADOR')
		ON CONFLICT (id_curso, id_periodo, id_escuela) DO UPDATE SET id_curso = EXCLUDED.id_curso
		RETURNING id_carga
	`, input.IDCurso, input.IDPeriodo, input.IDEscuela).Scan(&idCarga)
	if err != nil {
		return nil, err
	}
	return r.GetCargaByID(ctx, idCarga)
}

func (r *Repository) ApproveCarga(ctx context.Context, idCarga int, idUsuario int, justificacion string) error {
	result, err := r.db.Exec(ctx, `
		UPDATE carga_academica
		SET estado = 'AUTORIZADO', fecha_aprobacion = CURRENT_TIMESTAMP
		WHERE id_carga = $1
	`, idCarga)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("carga academica no encontrada")
	}
	return nil
}

func (r *Repository) GetGruposByCarga(ctx context.Context, idCarga int) ([]Grupo, error) {
	rows, err := r.db.Query(ctx, `
		SELECT g.id_grupo, g.id_carga, g.id_docente, g.id_grupo_teoria_ref,
		       g.codigo_grupo, g.tipo_componente, g.es_nueva_necesidad,
		       g.matriculados_proyectados, g.matriculados_reales,
		       c.horas_teoria + c.horas_practica as horas_semanales,
		       COALESCE(d.nombres || ' ' || d.apellidos, '') as docente_nombre
		FROM grupo g
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		JOIN curso c ON c.id_curso = ca.id_curso
		LEFT JOIN docente d ON d.id_docente = g.id_docente
		WHERE g.id_carga = $1
		ORDER BY g.codigo_grupo, g.tipo_componente
	`, idCarga)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grupos []Grupo
	for rows.Next() {
		var g Grupo
		var docenteNombre string
		if err := rows.Scan(
			&g.ID, &g.IDCarga, &g.IDDocente, &g.IDGrupoTeoriaRef,
			&g.CodigoGrupo, &g.TipoComponente, &g.EsNuevaNecesidad,
			&g.MatriculadosProyectados, &g.MatriculadosReales,
			&g.HorasSemanales, &docenteNombre,
		); err != nil {
			return nil, err
		}
		g.Docente = docenteNombre
		grupos = append(grupos, g)
	}
	return grupos, rows.Err()
}

func (r *Repository) CreateGrupo(ctx context.Context, idCarga int, input CreateGrupoInput) (*Grupo, error) {
	var idGrupo int
	err := r.db.QueryRow(ctx, `
		INSERT INTO grupo (id_carga, id_docente, codigo_grupo, tipo_componente, es_nueva_necesidad, matriculados_proyectados)
		VALUES ($1, $2, $3, $4::tipo_componente_enum, $5, $6)
		RETURNING id_grupo
	`, idCarga, input.IDDocente, input.CodigoGrupo, input.TipoComponente, input.EsNuevaNecesidad, input.MatriculadosProyectados).Scan(&idGrupo)
	if err != nil {
		return nil, err
	}

	grupos, err := r.GetGruposByCarga(ctx, idCarga)
	if err != nil {
		return nil, err
	}
	for i := range grupos {
		if grupos[i].ID == idGrupo {
			return &grupos[i], nil
		}
	}
	return nil, nil
}

func (r *Repository) UpdateGrupo(ctx context.Context, idGrupo int, input UpdateGrupoInput) (*Grupo, error) {
	_, err := r.db.Exec(ctx, `
		UPDATE grupo SET
			id_docente = COALESCE($2, id_docente),
			codigo_grupo = COALESCE($3, codigo_grupo),
			es_nueva_necesidad = COALESCE($4, es_nueva_necesidad),
			matriculados_proyectados = COALESCE($5, matriculados_proyectados)
		WHERE id_grupo = $1
	`, idGrupo, input.IDDocente, input.CodigoGrupo, input.EsNuevaNecesidad, input.MatriculadosProyectados)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, `
		SELECT g.id_grupo, g.id_carga, g.id_docente, g.id_grupo_teoria_ref,
		       g.codigo_grupo, g.tipo_componente, g.es_nueva_necesidad,
		       g.matriculados_proyectados, g.matriculados_reales,
		       c.horas_teoria + c.horas_practica as horas_semanales,
		       COALESCE(d.nombres || ' ' || d.apellidos, '') as docente_nombre
		FROM grupo g
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		JOIN curso c ON c.id_curso = ca.id_curso
		LEFT JOIN docente d ON d.id_docente = g.id_docente
		WHERE g.id_grupo = $1
	`, idGrupo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var g Grupo
		var docenteNombre string
		if err := rows.Scan(
			&g.ID, &g.IDCarga, &g.IDDocente, &g.IDGrupoTeoriaRef,
			&g.CodigoGrupo, &g.TipoComponente, &g.EsNuevaNecesidad,
			&g.MatriculadosProyectados, &g.MatriculadosReales,
			&g.HorasSemanales, &docenteNombre,
		); err != nil {
			return nil, err
		}
		g.Docente = docenteNombre
		return &g, nil
	}
	return nil, errors.New("grupo no encontrado")
}

func (r *Repository) GetHorasDocente(ctx context.Context, idDocente, idPeriodo int) (int, error) {
	var totalHoras int
	err := r.db.QueryRow(ctx, `
		SELECT COALESCE(SUM(c.horas_teoria + c.horas_practica), 0)
		FROM grupo g
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		JOIN curso c ON c.id_curso = ca.id_curso
		WHERE g.id_docente = $1 AND ca.id_periodo = $2 AND ca.estado = 'AUTORIZADO'
	`, idDocente, idPeriodo).Scan(&totalHoras)
	return totalHoras, err
}

func (r *Repository) GetDocentesResumen(ctx context.Context, idEscuela, idPeriodo int) ([]ResumenDocente, error) {
	rows, err := r.db.Query(ctx, `
		SELECT DISTINCT d.id_docente, d.codigo_plaza, d.nombres, d.apellidos, e.id_escuela, e.nombre as escuela
		FROM docente d
		JOIN departamento_academico dep ON dep.id_departamento = d.id_departamento
		JOIN escuela_profesional e ON e.id_departamento = dep.id_departamento
		JOIN grupo g ON g.id_docente = d.id_docente
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		WHERE ca.id_periodo = $1 AND g.id_docente IS NOT NULL
		ORDER BY d.apellidos, d.nombres
	`, idPeriodo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resumenes []ResumenDocente
	for rows.Next() {
		var res ResumenDocente
		if err := rows.Scan(&res.IDDocente, &res.CodigoPlaza, &res.Nombre, &res.IDEscuela, &res.Escuela); err != nil {
			return nil, err
		}

		cursos, err := r.GetCursosDocente(ctx, res.IDDocente, idPeriodo)
		if err != nil {
			return nil, err
		}
		res.Cursos = cursos

		var total int
		for _, c := range cursos {
			total += c.HorasSemanales
		}
		res.HorasAsignadas = total
		res.HorasRestantes = 16 - total
		if res.HorasRestantes < 0 {
			res.HorasRestantes = 0
		}

		resumenes = append(resumenes, res)
	}
	return resumenes, rows.Err()
}

func (r *Repository) GetCursosDocente(ctx context.Context, idDocente, idPeriodo int) ([]CursoDocente, error) {
	rows, err := r.db.Query(ctx, `
		SELECT ca.id_carga, c.codigo, c.nombre, g.codigo_grupo,
		       c.horas_teoria + c.horas_practica as horas_semanales, ca.estado
		FROM grupo g
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		JOIN curso c ON c.id_curso = ca.id_curso
		WHERE g.id_docente = $1 AND ca.id_periodo = $2
		ORDER BY c.codigo
	`, idDocente, idPeriodo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cursos []CursoDocente
	for rows.Next() {
		var c CursoDocente
		if err := rows.Scan(&c.IDCarga, &c.CodigoCurso, &c.NombreCurso, &c.CodigoGrupo, &c.HorasSemanales, &c.Estado); err != nil {
			return nil, err
		}
		cursos = append(cursos, c)
	}
	return cursos, rows.Err()
}

func (r *Repository) GetPeriodoActivo(ctx context.Context) (*int, error) {
	var idPeriodo int
	err := r.db.QueryRow(ctx, `
		SELECT id_periodo FROM periodo_academico WHERE activo = true LIMIT 1
	`).Scan(&idPeriodo)
	if err != nil {
		return nil, err
	}
	return &idPeriodo, nil
}

func (r *Repository) GetBloquesDocente(ctx context.Context, docenteID int, idPeriodo int) ([]BloqueAsignado, error) {
	rows, err := r.db.Query(ctx, `
		SELECT bh.id_bloque, bh.id_horario, bh.id_grupo, g.id_carga, bh.id_aula,
		       g.id_docente, e.id_escuela, e.nombre as escuela,
		       c.codigo as curso_codigo, c.nombre as curso_nombre, g.codigo_grupo,
		       bh.dia_semana, bh.slot_inicio, bh.slot_fin
		FROM bloque_horario bh
		JOIN grupo g ON g.id_grupo = bh.id_grupo
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		JOIN curso c ON c.id_curso = ca.id_curso
		JOIN escuela_profesional e ON e.id_escuela = ca.id_escuela
		JOIN horario h ON h.id_horario = bh.id_horario
		WHERE g.id_docente = $1 AND ca.id_periodo = $2 AND h.id_periodo = $2
		ORDER BY bh.dia_semana, bh.slot_inicio
	`, docenteID, idPeriodo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bloques []BloqueAsignado
	for rows.Next() {
		var b BloqueAsignado
		if err := rows.Scan(&b.ID, &b.IDHorario, &b.IDGrupo, &b.IDCarga, &b.IDAula,
			&b.DocenteID, &b.EscuelaID, &b.EscuelaNombre,
			&b.CursoCodigo, &b.CursoNombre, &b.GrupoCodigo,
			&b.DiaSemana, &b.SlotInicio, &b.SlotFin); err != nil {
			return nil, err
		}
		bloques = append(bloques, b)
	}
	return bloques, rows.Err()
}

type ConflictoHorario struct {
	Tipo            string `json:"tipo"`
	DiaSemana       int    `json:"dia_semana"`
	SlotInicio      int    `json:"slot_inicio"`
	SlotFin         int    `json:"slot_fin"`
	Escuela         string `json:"escuela"`
	CursoCodigo     string `json:"curso_codigo"`
	CursoNombre     string `json:"curso_nombre"`
	GrupoCodigo     string `json:"grupo_codigo"`
}

func (r *Repository) VerificarConflictoDocente(ctx context.Context, docenteID int, diaSemana int, slotInicio int, slotFin int, idPeriodo int) (*ConflictoHorario, error) {
	rows, err := r.db.Query(ctx, `
		SELECT e.nombre, c.codigo, c.nombre, g.codigo_grupo,
		       bh.dia_semana, bh.slot_inicio, bh.slot_fin
		FROM bloque_horario bh
		JOIN grupo g ON g.id_grupo = bh.id_grupo
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		JOIN curso c ON c.id_curso = ca.id_curso
		JOIN escuela_profesional e ON e.id_escuela = ca.id_escuela
		JOIN horario h ON h.id_horario = bh.id_horario
		WHERE g.id_docente = $1 AND ca.id_periodo = $2 AND h.id_periodo = $2
		  AND bh.dia_semana = $3
		  AND bh.slot_inicio < $5 AND bh.slot_fin > $4
	`, docenteID, idPeriodo, diaSemana, slotInicio, slotFin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var b ConflictoHorario
		if err := rows.Scan(&b.Escuela, &b.CursoCodigo, &b.CursoNombre, &b.GrupoCodigo,
			&b.DiaSemana, &b.SlotInicio, &b.SlotFin); err != nil {
			return nil, err
		}
		b.Tipo = "CRUCE_HORARIO"
		return &b, nil
	}
	return nil, nil
}

func formatSlotRange(inicio, fin int) string {
	horas := map[int]string{
		1:  "07:00", 2:  "08:00", 3:  "09:00", 4:  "10:00", 5:  "11:00", 6:  "12:00",
		7:  "13:00", 8:  "14:00", 9:  "15:00", 10: "16:00", 11: "17:00", 12: "18:00",
		13: "19:00", 14: "20:00",
	}
	start := horas[inicio]
	end := horas[fin]
	if start == "" {
		start = "??:??"
	}
	if end == "" {
		end = "??:??"
	}
	return start + "-" + end
}
