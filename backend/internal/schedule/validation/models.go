package validation

type Severity string

const (
	SeverityBlocker Severity = "BLOCKER"
	SeverityWarning Severity = "WARNING"
	SeverityInfo    Severity = "INFO"
)

type RuleID string

const (
	RuleTeacherConflict       RuleID = "RV-01"
	RuleRoomConflict          RuleID = "RV-02"
	RuleDepartmentSession     RuleID = "RV-03"
	RuleNoTransferTime        RuleID = "RV-04a"
	RuleInsufficientTransfer  RuleID = "RV-04b"
	RuleTeachingLoadLimit     RuleID = "RV-05"
	RuleSameSeriesConflict    RuleID = "RV-06"
	RuleSharedRoomReservation RuleID = "RV-07"
	RuleAuditJustification    RuleID = "RV-08"
	RuleCapacityReadjustment  RuleID = "RV-09"
)

type Finding struct {
	Rule     RuleID   `json:"rule"`
	Severity Severity `json:"severity"`
	Message  string   `json:"message"`
}

func (f Finding) Blocks() bool {
	return f.Severity == SeverityBlocker
}

type ScheduleState string

const (
	StateDraft       ScheduleState = "BORRADOR"
	StatePreliminary ScheduleState = "PRELIMINAR"
	StateReadjusting ScheduleState = "EN_REAJUSTE"
	StateOfficial    ScheduleState = "OFICIAL"
	StateReadjusted  ScheduleState = "REAJUSTADO"
)

type Block struct {
	ID           int
	ScheduleID   int
	SchoolID     int
	GroupID      int
	CourseID     int
	SeriesID     int
	TeacherID    int
	DepartmentID int
	RoomID       int
	RoomShared   bool
	PavilionID   int
	Day          int
	StartSlot    int
	EndSlot      int
	Enrollment   int
	RoomCapacity int
}

type DepartmentSession struct {
	DepartmentID int
	Day          int
	StartSlot    int
	EndSlot      int
}

type Distance struct {
	FromPavilionID int
	ToPavilionID   int
	Minutes        int
}

type PlacementInput struct {
	Proposed           Block
	Existing           []Block
	DepartmentSessions []DepartmentSession
	Distances          []Distance
	State              ScheduleState
}

type AuditChangeInput struct {
	State         ScheduleState
	Justification string
}

type TeachingLoadInput struct {
	TeacherID   int
	WeeklyHours int
	Confirmed   bool
}
