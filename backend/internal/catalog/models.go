package catalog

import "time"

type Facultad struct {
	ID     int    `json:"id_facultad"`
	Nombre string `json:"nombre"`
}

type Departamento struct {
	ID         int    `json:"id_departamento"`
	IDFacultad int    `json:"id_facultad"`
	Nombre     string `json:"nombre"`
}

type Escuela struct {
	ID             int    `json:"id_escuela"`
	IDFacultad     int    `json:"id_facultad"`
	IDDepartamento int    `json:"id_departamento"`
	Nombre         string `json:"nombre"`
}

type Aula struct {
	ID           int    `json:"id_aula"`
	IDPabellon   int    `json:"id_pabellon"`
	IDEscuela    *int   `json:"id_escuela"`
	Codigo       string `json:"codigo"`
	Tipo         string `json:"tipo"`
	Aforo        int    `json:"aforo"`
	EsCompartida bool   `json:"es_compartida"`
}

type Usuario struct {
	ID     int    `json:"id_usuario"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Rol    string `json:"rol"`
}

type PlanEstudio struct {
	ID        int    `json:"id_plan"`
	IDEscuela int    `json:"id_escuela"`
	Codigo    string `json:"codigo_plan"`
	Nombre    string `json:"nombre"`
}

type Serie struct {
	ID          int `json:"id_serie"`
	IDPlan      int `json:"id_plan"`
	NumeroCiclo int `json:"numero_ciclo"`
}

type Curso struct {
	ID            int    `json:"id_curso"`
	IDSerie       int    `json:"id_serie"`
	Codigo        string `json:"codigo"`
	Nombre        string `json:"nombre"`
	Creditos      int    `json:"creditos"`
	HorasTeoria   int    `json:"horas_teoria"`
	HorasPractica int    `json:"horas_practica"`
}

type Docente struct {
	ID             int    `json:"id_docente"`
	IDDepartamento int    `json:"id_departamento"`
	CodigoPlaza    string `json:"codigo_plaza"`
	Nombres        string `json:"nombres"`
	Apellidos      string `json:"apellidos"`
	Email          string `json:"email"`
}

type PeriodoAcademico struct {
	ID     int    `json:"id_periodo"`
	Codigo string `json:"codigo"`
	Activo bool   `json:"activo"`
}

type SesionDepartamento struct {
	ID             int    `json:"id_sesion"`
	IDDepartamento int    `json:"id_departamento"`
	IDPeriodo      int    `json:"id_periodo"`
	DiaSemana      int    `json:"dia_semana"`
	HoraInicio     string `json:"hora_inicio"`
	HoraFin        string `json:"hora_fin"`
}

type Local struct {
	ID     int    `json:"id_local"`
	Nombre string `json:"nombre"`
}

type Pabellon struct {
	ID      int    `json:"id_pabellon"`
	IDLocal int    `json:"id_local"`
	Codigo  string `json:"codigo"`
	Nombre  string `json:"nombre"`
}

type Distancia struct {
	DesdeID int `json:"id_pabellon_origen"`
	HastaID int `json:"id_pabellon_destino"`
	Minutos int `json:"tiempo_minutos"`
}

type CargaAcademica struct {
	IDPeriodo       int     `json:"id_periodo"`
	IDEscuela       int     `json:"id_escuela"`
	IDCurso         int     `json:"id_curso"`
	IDCarga         int     `json:"id_carga"`
	Estado          string  `json:"estado"`
	FechaAprobacion *string `json:"fecha_aprobacion"`
}

type Grupo struct {
	ID                      int    `json:"id_grupo"`
	IDCarga                 int    `json:"id_carga"`
	IDDocente               *int   `json:"id_docente"`
	IDGrupoTeoriaRef        *int   `json:"id_grupo_teoria_ref"`
	CodigoGrupo             string `json:"codigo_grupo"`
	TipoComponente          string `json:"tipo_componente"`
	EsNuevaNecesidad        bool   `json:"es_nueva_necesidad"`
	MatriculadosProyectados int    `json:"matriculados_proyectados"`
	MatriculadosReales      int    `json:"matriculados_reales"`
}

type Horario struct {
	ID                 int       `json:"id_horario"`
	IDEscuela          int       `json:"id_escuela"`
	IDPeriodo          int       `json:"id_periodo"`
	Estado             string    `json:"estado"`
	VersionReajuste    int       `json:"version_reajuste"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

type BloqueHorario struct {
	ID         int  `json:"id_bloque"`
	IDHorario  int  `json:"id_horario"`
	IDGrupo    int  `json:"id_grupo"`
	IDAula     int  `json:"id_aula"`
	IDDocente  *int `json:"id_docente"`
	DiaSemana  int  `json:"dia_semana"`
	SlotInicio int  `json:"slot_inicio"`
	SlotFin    int  `json:"slot_fin"`
}

type BitacoraAuditoria struct {
	ID                  int    `json:"id_log"`
	IDHorario           int    `json:"id_horario"`
	IDUsuario           int    `json:"id_usuario"`
	Accion              string `json:"accion"`
	MotivoJustificacion string `json:"motivo_justificacion"`
	VersionResultante   int    `json:"version_resultante"`
	FechaHora           string `json:"fecha_hora"`
}

type CreateHorarioInput struct {
	IDEscuela  int `json:"id_escuela" binding:"required"`
	IDPeriodo  int `json:"id_periodo" binding:"required"`
}

type CreateBloqueInput struct {
	IDHorario  int    `json:"id_horario" binding:"required"`
	IDGrupo    int    `json:"id_grupo" binding:"required"`
	IDAula     int    `json:"id_aula" binding:"required"`
	IDDocente  *int   `json:"id_docente"`
	DiaSemana  int    `json:"dia_semana" binding:"required,min=1,max=6"`
	SlotInicio int    `json:"slot_inicio" binding:"required,min=1,max=14"`
	SlotFin    int    `json:"slot_fin" binding:"required,min=1,max=14"`
}

type UpdateBloqueInput struct {
	IDAula     *int  `json:"id_aula"`
	IDDocente  *int  `json:"id_docente"`
	DiaSemana  *int  `json:"dia_semana"`
	SlotInicio *int  `json:"slot_inicio"`
	SlotFin    *int  `json:"slot_fin"`
}

type BloqueContexto struct {
	ID               int    `json:"id_bloque"`
	IDHorario        int    `json:"id_horario"`
	IDGrupo          int    `json:"id_grupo"`
	IDCarga          int    `json:"id_carga"`
	IDEscuela        int    `json:"id_escuela"`
	IDAula           int    `json:"id_aula"`
	IDAulaPabellon   int    `json:"id_aula_pabellon"`
	IDDocente        *int   `json:"id_docente"`
	IDDocenteDepto   int    `json:"id_docente_departamento"`
	CodigoAula       string `json:"codigo_aula"`
	NombreDocente    string `json:"nombre_docente"`
	CodigoGrupo      string `json:"codigo_grupo"`
	TipoComponente   string `json:"tipo_componente"`
	EstadoHorario    string `json:"estado_horario"`
	DiaSemana        int    `json:"dia_semana"`
	SlotInicio       int    `json:"slot_inicio"`
	SlotFin          int    `json:"slot_fin"`
	NombreEscuela    string `json:"nombre_escuela"`
	CodigoCurso      string `json:"codigo_curso"`
	NombreCurso      string `json:"nombre_curso"`
}
