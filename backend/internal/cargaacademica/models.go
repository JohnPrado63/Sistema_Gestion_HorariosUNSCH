package cargaacademica

type CargaAcademica struct {
	IDCarga         int      `json:"id_carga"`
	IDCurso         int      `json:"id_curso"`
	IDPeriodo       int      `json:"id_periodo"`
	IDEscuela       int      `json:"id_escuela"`
	Estado          string   `json:"estado"`
	FechaAprobacion *string  `json:"fecha_aprobacion,omitempty"`
	Curso           *CursoInfo `json:"curso,omitempty"`
	Escuela         string   `json:"escuela,omitempty"`
	Grupos          []Grupo  `json:"grupos,omitempty"`
}

type CursoInfo struct {
	ID            int    `json:"id_curso"`
	Codigo        string `json:"codigo"`
	Nombre        string `json:"nombre"`
	HorasTeoria   int    `json:"horas_teoria"`
	HorasPractica int    `json:"horas_practica"`
	Creditos      int    `json:"creditos"`
}

type Grupo struct {
	ID                    int     `json:"id_grupo"`
	IDCarga               int     `json:"id_carga"`
	IDDocente             *int    `json:"id_docente,omitempty"`
	Docente               string  `json:"docente,omitempty"`
	IDGrupoTeoriaRef      *int    `json:"id_grupo_teoria_ref,omitempty"`
	CodigoGrupo           string  `json:"codigo_grupo"`
	TipoComponente        string  `json:"tipo_componente"`
	EsNuevaNecesidad      bool    `json:"es_nueva_necesidad"`
	MatriculadosProyectados int   `json:"matriculados_proyectados"`
	MatriculadosReales    int     `json:"matriculados_reales"`
	HorasSemanales        int     `json:"horas_semanales"`
}

type CreateCargaInput struct {
	IDCurso   int `json:"id_curso" binding:"required"`
	IDPeriodo int `json:"id_periodo" binding:"required"`
	IDEscuela int `json:"id_escuela" binding:"required"`
}

type CreateGrupoInput struct {
	IDDocente              *int   `json:"id_docente"`
	CodigoGrupo            string `json:"codigo_grupo" binding:"required"`
	TipoComponente         string `json:"tipo_componente" binding:"required,oneof=TEORIA PRACTICA"`
	EsNuevaNecesidad       bool   `json:"es_nueva_necesidad"`
	MatriculadosProyectados int   `json:"matriculados_proyectados"`
}

type UpdateGrupoInput struct {
	IDDocente              *int   `json:"id_docente"`
	CodigoGrupo            string `json:"codigo_grupo"`
	EsNuevaNecesidad       *bool  `json:"es_nueva_necesidad"`
	MatriculadosProyectados *int  `json:"matriculados_proyectados"`
}

type AprobarCargaInput struct {
	IDUsuario int    `json:"id_usuario" binding:"required"`
	Justificacion string `json:"justificacion"`
}

type ResumenDocente struct {
	IDDocente       int    `json:"id_docente"`
	Nombre          string `json:"nombre"`
	CodigoPlaza     string `json:"codigo_plaza"`
	IDEscuela       int    `json:"id_escuela"`
	Escuela         string `json:"escuela"`
	HorasAsignadas  int    `json:"horas_asignadas"`
	HorasRestantes  int    `json:"horas_restantes"`
	Cursos          []CursoDocente `json:"cursos"`
}

type CursoDocente struct {
	IDCarga      int    `json:"id_carga"`
	CodigoCurso  string `json:"codigo_curso"`
	NombreCurso  string `json:"nombre_curso"`
	CodigoGrupo  string `json:"codigo_grupo"`
	HorasSemanales int  `json:"horas_semanales"`
	Estado       string `json:"estado"`
}

type HorarioSugerido struct {
	DiaSemana  int `json:"dia_semana"`
	SlotInicio int `json:"slot_inicio"`
	SlotFin    int `json:"slot_fin"`
	IDAula     int `json:"id_aula"`
	EsValido   bool `json:"es_valido"`
	Reason     string `json:"reason,omitempty"`
}

type PropuestaBloque struct {
	IDGrupo    int             `json:"id_grupo"`
	Horarios   []HorarioSugerido `json:"horarios"`
}

type DisponibilidadDocente struct {
	DocenteID       int               `json:"docente_id"`
	Nombre          string            `json:"nombre"`
	BloquesAsignados []BloqueAsignado `json:"bloques_asignados"`
}

type BloqueAsignado struct {
	ID           int    `json:"id_bloque"`
	IDHorario    int    `json:"id_horario"`
	IDGrupo      int    `json:"id_grupo"`
	IDCarga      int    `json:"id_carga"`
	IDAula       int    `json:"id_aula"`
	DocenteID    int    `json:"docente_id"`
	EscuelaID    int    `json:"escuela_id"`
	EscuelaNombre string `json:"escuela_nombre"`
	CursoCodigo  string `json:"curso_codigo"`
	CursoNombre  string `json:"curso_nombre"`
	GrupoCodigo  string `json:"grupo_codigo"`
	DiaSemana    int    `json:"dia_semana"`
	SlotInicio   int    `json:"slot_inicio"`
	SlotFin      int    `json:"slot_fin"`
}

type VerificarDisponibilidadInput struct {
	DocenteID   int    `json:"docente_id" binding:"required"`
	DiaSemana   int    `json:"dia_semana" binding:"required,min=1,max=6"`
	SlotInicio  int    `json:"slot_inicio" binding:"required,min=1,max=14"`
	SlotFin     int    `json:"slot_fin" binding:"required,min=1,max=14"`
}

type VerificarDisponibilidadResponse struct {
	TieneConflicto bool              `json:"tiene_conflicto"`
	Conflictos     []ConflictoDetalle `json:"conflictos,omitempty"`
	SlotsLibres    []SlotLibre       `json:"slots_libres,omitempty"`
}

type ConflictoDetalle struct {
	Tipo           string `json:"tipo"`
	DiaSemana      int    `json:"dia_semana"`
	SlotSolicitado string `json:"slot_solicitado"`
	SlotOcupado    string `json:"slot_ocupado"`
	Escuela        string `json:"escuela"`
	CursoCodigo    string `json:"curso_codigo"`
	CursoNombre    string `json:"curso_nombre"`
	GrupoCodigo    string `json:"grupo_codigo"`
}

type SlotLibre struct {
	DiaSemana  int    `json:"dia_semana"`
	SlotInicio int    `json:"slot_inicio"`
	SlotFin    int    `json:"slot_fin"`
	HoraInicio string `json:"hora_inicio"`
	HoraFin    string `json:"hora_fin"`
}
