package validationui

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"unsch-horarios/backend/internal/schedule/validation"
)

type Scenario struct {
	ID          string               `json:"id"`
	Rule        validation.RuleID    `json:"rule"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Expected    validation.Severity  `json:"expected"`
	Findings    []validation.Finding `json:"findings"`
}

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Page(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, pageHTML)
}

func (h Handler) Scenarios(c *gin.Context) {
	c.JSON(http.StatusOK, buildScenarios())
}

func buildScenarios() []Scenario {
	return []Scenario{
		placementScenario(
			"rv01-docente",
			validation.RuleTeacherConflict,
			"Cruce de docente",
			"El docente 10 ya dicta otra clase el lunes entre los slots 4 y 5.",
			validation.SeverityBlocker,
			validation.PlacementInput{
				Proposed: demoBlock(1, 10, 1, 1, 100, 1, 3, 4),
				Existing: []validation.Block{demoBlock(2, 10, 2, 2, 200, 1, 4, 5)},
			},
		),
		placementScenario(
			"rv02-aula",
			validation.RuleRoomConflict,
			"Cruce de aula",
			"El aula 5 ya esta ocupada en el rango elegido.",
			validation.SeverityBlocker,
			validation.PlacementInput{
				Proposed: demoBlock(1, 10, 1, 1, 100, 1, 3, 4),
				Existing: []validation.Block{demoBlock(2, 20, 2, 2, 200, 1, 4, 5)},
			},
		),
		placementScenario(
			"rv03-sesion",
			validation.RuleDepartmentSession,
			"Sesion de departamento",
			"La clase cae dentro de la franja semanal reservada del departamento 7.",
			validation.SeverityBlocker,
			validation.PlacementInput{
				Proposed:           withDepartment(demoBlock(1, 10, 1, 1, 100, 1, 3, 4), 7),
				DepartmentSessions: []validation.DepartmentSession{{DepartmentID: 7, Day: 1, StartSlot: 4, EndSlot: 5}},
			},
		),
		placementScenario(
			"rv04a-traslado-sin-tiempo",
			validation.RuleNoTransferTime,
			"Traslado sin tiempo",
			"El docente termina en un pabellon y empieza inmediatamente en otro.",
			validation.SeverityBlocker,
			validation.PlacementInput{
				Proposed:  withPavilion(demoBlock(1, 10, 1, 1, 100, 1, 4, 5), 2),
				Existing:  []validation.Block{withPavilion(demoBlock(2, 10, 1, 2, 100, 1, 2, 3), 1)},
				Distances: []validation.Distance{{FromPavilionID: 1, ToPavilionID: 2, Minutes: 10}},
			},
		),
		placementScenario(
			"rv04b-traslado-ajustado",
			validation.RuleInsufficientTransfer,
			"Traslado con tiempo insuficiente",
			"Hay una hora libre, pero la matriz estima 90 minutos de traslado.",
			validation.SeverityWarning,
			validation.PlacementInput{
				Proposed:  withPavilion(demoBlock(1, 10, 1, 1, 100, 1, 5, 6), 2),
				Existing:  []validation.Block{withPavilion(demoBlock(2, 10, 1, 2, 100, 1, 2, 3), 1)},
				Distances: []validation.Distance{{FromPavilionID: 1, ToPavilionID: 2, Minutes: 90}},
			},
		),
		teachingLoadScenario(),
		placementScenario(
			"rv06-serie",
			validation.RuleSameSeriesConflict,
			"Cruce en misma serie",
			"Dos cursos obligatorios de la serie 8 coinciden en el mismo horario.",
			validation.SeverityInfo,
			validation.PlacementInput{
				Proposed: withCourse(demoBlock(1, 10, 1, 1, 8, 1, 3, 4), 101),
				Existing: []validation.Block{withCourse(demoBlock(2, 20, 2, 2, 8, 1, 4, 5), 102)},
			},
		),
		placementScenario(
			"rv07-aula-compartida",
			validation.RuleSharedRoomReservation,
			"Aula compartida reservada",
			"Otra escuela ya reservo el aula compartida 5 en el mismo instante.",
			validation.SeverityBlocker,
			validation.PlacementInput{
				Proposed: withSharedRoom(demoBlock(1, 10, 1, 1, 100, 1, 3, 4)),
				Existing: []validation.Block{withSharedRoom(demoBlock(2, 20, 2, 2, 200, 1, 4, 5))},
			},
		),
		auditScenario(),
		capacityScenario("rv09-aforo-preliminar", "Aforo excedido en preliminar", validation.StatePreliminary, validation.SeverityInfo),
		capacityScenario("rv09-aforo-reajuste", "Aforo excedido en reajuste", validation.StateReadjusting, validation.SeverityWarning),
		placementScenario(
			"ok-limpio",
			"OK",
			"Caso valido",
			"No hay cruces ni advertencias; la asignacion puede continuar.",
			"OK",
			validation.PlacementInput{
				Proposed: demoBlock(1, 10, 1, 1, 100, 1, 6, 7),
				Existing: []validation.Block{demoBlock(2, 20, 2, 2, 200, 1, 3, 4)},
				State:    validation.StateDraft,
			},
		),
	}
}

func placementScenario(id string, rule validation.RuleID, title string, description string, expected validation.Severity, input validation.PlacementInput) Scenario {
	return Scenario{ID: id, Rule: rule, Title: title, Description: description, Expected: expected, Findings: validation.ValidatePlacement(input)}
}

func teachingLoadScenario() Scenario {
	findings := validation.ValidateTeachingLoadApproval(validation.TeachingLoadInput{TeacherID: 10, WeeklyHours: 18})
	return Scenario{ID: "rv05-carga", Rule: validation.RuleTeachingLoadLimit, Title: "Tope de carga lectiva", Description: "El docente supera 16 horas semanales y requiere confirmacion documentada de DGA.", Expected: validation.SeverityWarning, Findings: findings}
}

func auditScenario() Scenario {
	findings := validation.ValidateAuditChange(validation.AuditChangeInput{State: validation.StateOfficial})
	return Scenario{ID: "rv08-auditoria", Rule: validation.RuleAuditJustification, Title: "Justificacion obligatoria", Description: "Un cambio sobre horario OFICIAL no puede guardarse sin motivo.", Expected: validation.SeverityBlocker, Findings: findings}
}

func capacityScenario(id string, title string, state validation.ScheduleState, expected validation.Severity) Scenario {
	proposed := demoBlock(1, 10, 1, 1, 100, 1, 3, 4)
	proposed.Enrollment = 55
	proposed.RoomCapacity = 40
	return Scenario{ID: id, Rule: validation.RuleCapacityReadjustment, Title: title, Description: "La matricula supera el aforo del aula asignada.", Expected: expected, Findings: validation.ValidatePlacement(validation.PlacementInput{Proposed: proposed, State: state})}
}

func demoBlock(id int, teacherID int, schoolID int, groupID int, seriesID int, day int, start int, end int) validation.Block {
	return validation.Block{ID: id, TeacherID: teacherID, SchoolID: schoolID, GroupID: groupID, SeriesID: seriesID, Day: day, StartSlot: start, EndSlot: end}
}

func withDepartment(block validation.Block, departmentID int) validation.Block {
	block.DepartmentID = departmentID
	return block
}

func withPavilion(block validation.Block, pavilionID int) validation.Block {
	block.PavilionID = pavilionID
	return block
}

func withCourse(block validation.Block, courseID int) validation.Block {
	block.CourseID = courseID
	return block
}

func withSharedRoom(block validation.Block) validation.Block {
	block.RoomShared = true
	return block
}
