package catalog

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryInterface interface {
	ListFacultades(ctx context.Context) ([]Facultad, error)
	ListDepartamentos(ctx context.Context) ([]Departamento, error)
	ListEscuelas(ctx context.Context) ([]Escuela, error)
	ListAulas(ctx context.Context) ([]Aula, error)
	ListUsuarios(ctx context.Context) ([]Usuario, error)
	ListPlanesEstudio(ctx context.Context) ([]PlanEstudio, error)
	ListSeries(ctx context.Context) ([]Serie, error)
	ListCursos(ctx context.Context) ([]Curso, error)
	ListDocentes(ctx context.Context) ([]Docente, error)
	ListPeriodos(ctx context.Context) ([]PeriodoAcademico, error)
	ListSesionesDepartamento(ctx context.Context) ([]SesionDepartamento, error)
	ListLocales(ctx context.Context) ([]Local, error)
	ListPabellones(ctx context.Context) ([]Pabellon, error)
	ListDistancias(ctx context.Context) ([]Distancia, error)
	ListCargasAcademicas(ctx context.Context) ([]CargaAcademica, error)
	ListGrupos(ctx context.Context) ([]Grupo, error)
	ListHorarios(ctx context.Context) ([]Horario, error)
	ListBloquesHorario(ctx context.Context) ([]BloqueHorario, error)
	ListBitacoraAuditoria(ctx context.Context) ([]BitacoraAuditoria, error)
	CreateHorario(ctx context.Context, input CreateHorarioInput) (*Horario, error)
	VerificarConflictoBloque(ctx context.Context, input CreateBloqueInput) ([]ConflictoBloque, error)
	CreateBloque(ctx context.Context, input CreateBloqueInput) (*BloqueHorario, error)
	GetBloquesByHorario(ctx context.Context, idHorario int) ([]BloqueContexto, error)
	GetGruposParaHorario(ctx context.Context, idEscuela int, idPeriodo int) ([]GrupoInfo, error)
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

func (r Repository) ListAulas(ctx context.Context) ([]Aula, error) {
	rows, err := r.db.Query(ctx, `SELECT id_aula, id_pabellon, id_escuela, codigo, tipo::text, aforo, es_compartida FROM aula ORDER BY codigo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Aula, 0)
	for rows.Next() {
		var item Aula
		if err := rows.Scan(&item.ID, &item.IDPabellon, &item.IDEscuela, &item.Codigo, &item.Tipo, &item.Aforo, &item.EsCompartida); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
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

func (r Repository) ListSeries(ctx context.Context) ([]Serie, error) {
	rows, err := r.db.Query(ctx, `SELECT id_serie, id_plan, numero_ciclo FROM serie ORDER BY id_plan, numero_ciclo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Serie, 0)
	for rows.Next() {
		var item Serie
		if err := rows.Scan(&item.ID, &item.IDPlan, &item.NumeroCiclo); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r Repository) ListCursos(ctx context.Context) ([]Curso, error) {
	rows, err := r.db.Query(ctx, `SELECT id_curso, id_serie, codigo, nombre, creditos, horas_teoria, horas_practica FROM curso ORDER BY codigo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Curso, 0)
	for rows.Next() {
		var item Curso
		if err := rows.Scan(&item.ID, &item.IDSerie, &item.Codigo, &item.Nombre, &item.Creditos, &item.HorasTeoria, &item.HorasPractica); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
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

func (r Repository) ListCargasAcademicas(ctx context.Context) ([]CargaAcademica, error) {
	rows, err := r.db.Query(ctx, `SELECT id_carga, id_curso, id_periodo, id_escuela, estado, fecha_aprobacion FROM carga_academica ORDER BY id_periodo, id_escuela`)
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
	rows, err := r.db.Query(ctx, `SELECT id_horario, id_escuela, id_periodo, estado::text, version_reajuste, fecha_actualizacion FROM horario ORDER BY id_escuela, id_periodo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Horario, 0)
	for rows.Next() {
		var item Horario
		if err := rows.Scan(&item.ID, &item.IDEscuela, &item.IDPeriodo, &item.Estado, &item.VersionReajuste, &item.FechaActualizacion); err != nil {
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

func (r Repository) CreateHorario(ctx context.Context, input CreateHorarioInput) (*Horario, error) {
	var idHorario int
	err := r.db.QueryRow(ctx, `
		INSERT INTO horario (id_escuela, id_periodo, estado, version_reajuste)
		VALUES ($1, $2, 'BORRADOR', 0)
		ON CONFLICT (id_escuela, id_periodo) DO UPDATE SET id_escuela = EXCLUDED.id_escuela
		RETURNING id_horario
	`, input.IDEscuela, input.IDPeriodo).Scan(&idHorario)
	if err != nil {
		return nil, err
	}

	var h Horario
	err = r.db.QueryRow(ctx, `
		SELECT id_horario, id_escuela, id_periodo, estado::text, version_reajuste, fecha_actualizacion
		FROM horario WHERE id_horario = $1
	`, idHorario).Scan(&h.ID, &h.IDEscuela, &h.IDPeriodo, &h.Estado, &h.VersionReajuste, &h.FechaActualizacion)
	if err != nil {
		return nil, err
	}
	return &h, nil
}

type ConflictoBloque struct {
	Tipo       string `json:"tipo"`
	Mensaje    string `json:"mensaje"`
	Detalle    string `json:"detalle,omitempty"`
}

func (r Repository) VerificarConflictoBloque(ctx context.Context, input CreateBloqueInput) ([]ConflictoBloque, error) {
	var conflictos []ConflictoBloque

	if input.IDDocente != nil {
		var count int
		err := r.db.QueryRow(ctx, `
			SELECT COUNT(*) FROM bloque_horario bh
			JOIN grupo g ON g.id_grupo = bh.id_grupo
			JOIN carga_academica ca ON ca.id_carga = g.id_carga
			JOIN horario h ON h.id_horario = bh.id_horario
			WHERE g.id_docente = $1 AND h.id_periodo = $2
			  AND bh.dia_semana = $3
			  AND bh.slot_inicio < $5 AND bh.slot_fin > $4
		`, *input.IDDocente, input.IDHorario, input.DiaSemana, input.SlotInicio, input.SlotFin).Scan(&count)
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

func (r Repository) GetBloquesByHorario(ctx context.Context, idHorario int) ([]BloqueContexto, error) {
	rows, err := r.db.Query(ctx, `
		SELECT bh.id_bloque, bh.id_horario, bh.id_grupo, g.id_carga, e.id_escuela,
		       bh.id_aula, a.id_pabellon, bh.id_docente, COALESCE(d.id_departamento, 0),
		       a.codigo as codigo_aula,
		       COALESCE(d.nombres || ' ' || d.apellidos, '') as nombre_docente,
		       g.codigo_grupo, g.tipo_componente::text,
		       h.estado::text,
		       bh.dia_semana, bh.slot_inicio, bh.slot_fin,
		       e.nombre as escuela_nombre, c.codigo as curso_codigo, c.nombre as curso_nombre
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
			&b.NombreEscuela, &b.CodigoCurso, &b.NombreCurso); err != nil {
			return nil, err
		}
		items = append(items, b)
	}
	return items, rows.Err()
}

func (r Repository) GetGruposParaHorario(ctx context.Context, idEscuela int, idPeriodo int) ([]GrupoInfo, error) {
	rows, err := r.db.Query(ctx, `
		SELECT g.id_grupo, g.id_carga, g.codigo_grupo, g.tipo_componente::text,
		       g.id_docente, COALESCE(d.nombres || ' ' || d.apellidos, '') as docente_nombre,
		       c.codigo, c.nombre, c.horas_teoria, c.horas_practica
		FROM grupo g
		JOIN carga_academica ca ON ca.id_carga = g.id_carga
		JOIN curso c ON c.id_curso = ca.id_curso
		LEFT JOIN docente d ON d.id_docente = g.id_docente
		WHERE ca.id_escuela = $1 AND ca.id_periodo = $2 AND ca.estado = 'AUTORIZADO'
		ORDER BY c.codigo, g.codigo_grupo
	`, idEscuela, idPeriodo)
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
