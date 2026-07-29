package catalog

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

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
