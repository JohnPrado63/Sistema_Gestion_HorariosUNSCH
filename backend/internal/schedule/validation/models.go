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
	ID           int  `json:"id"`
	ScheduleID   int  `json:"schedule_id"`
	SchoolID     int  `json:"school_id"`
	GroupID      int  `json:"group_id"`
	CourseID     int  `json:"course_id"`
	SeriesID     int  `json:"series_id"`
	TeacherID    int  `json:"teacher_id"`
	DepartmentID int  `json:"department_id"`
	RoomID       int  `json:"room_id"`
	RoomShared   bool `json:"room_shared"`
	PavilionID   int  `json:"pavilion_id"`
	Day          int  `json:"day"`
	StartSlot    int  `json:"start_slot"`
	EndSlot      int  `json:"end_slot"`
	Enrollment   int  `json:"enrollment"`
	RoomCapacity int  `json:"room_capacity"`
}

type DepartmentSession struct {
	DepartmentID int `json:"department_id"`
	Day          int `json:"day"`
	StartSlot    int `json:"start_slot"`
	EndSlot      int `json:"end_slot"`
}

type Distance struct {
	FromPavilionID int `json:"from_pavilion_id"`
	ToPavilionID   int `json:"to_pavilion_id"`
	Minutes        int `json:"minutes"`
}

type PlacementInput struct {
	Proposed           Block               `json:"proposed"`
	Existing           []Block             `json:"existing"`
	DepartmentSessions []DepartmentSession `json:"department_sessions"`
	Distances          []Distance          `json:"distances"`
	State              ScheduleState       `json:"state"`
}

type AuditChangeInput struct {
	State         ScheduleState `json:"state"`
	Justification string        `json:"justification"`
}

type TeachingLoadInput struct {
	TeacherID   int  `json:"teacher_id"`
	WeeklyHours int  `json:"weekly_hours"`
	Confirmed   bool `json:"confirmed"`
}
