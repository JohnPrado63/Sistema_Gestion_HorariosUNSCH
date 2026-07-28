package validation

import "testing"

func TestValidatePlacementBlocksTeacherConflictRV01(t *testing.T) {
	findings := ValidatePlacement(PlacementInput{
		Proposed: block(1, 10, 1, 1, 1, 100, 1, 3, 4),
		Existing: []Block{
			block(2, 10, 2, 2, 2, 200, 1, 4, 5),
		},
	})

	assertFinding(t, findings, RuleTeacherConflict, SeverityBlocker)
}

func TestValidatePlacementBlocksRoomConflictRV02(t *testing.T) {
	findings := ValidatePlacement(PlacementInput{
		Proposed: block(1, 10, 1, 1, 5, 100, 1, 3, 4),
		Existing: []Block{
			block(2, 20, 2, 2, 5, 200, 1, 4, 5),
		},
	})

	assertFinding(t, findings, RuleRoomConflict, SeverityBlocker)
}

func TestValidatePlacementBlocksDepartmentSessionRV03(t *testing.T) {
	findings := ValidatePlacement(PlacementInput{
		Proposed: blockWithDepartment(1, 10, 1, 7, 1, 3, 4),
		DepartmentSessions: []DepartmentSession{
			{DepartmentID: 7, Day: 1, StartSlot: 4, EndSlot: 5},
		},
	})

	assertFinding(t, findings, RuleDepartmentSession, SeverityBlocker)
}

func TestValidatePlacementBlocksNoTransferTimeRV04a(t *testing.T) {
	findings := ValidatePlacement(PlacementInput{
		Proposed: blockWithPavilion(1, 10, 1, 2, 1, 4, 5),
		Existing: []Block{
			blockWithPavilion(2, 10, 1, 1, 1, 2, 3),
		},
		Distances: []Distance{{FromPavilionID: 1, ToPavilionID: 2, Minutes: 10}},
	})

	assertFinding(t, findings, RuleNoTransferTime, SeverityBlocker)
}

func TestValidatePlacementWarnsInsufficientTransferRV04b(t *testing.T) {
	findings := ValidatePlacement(PlacementInput{
		Proposed: blockWithPavilion(1, 10, 1, 2, 1, 5, 6),
		Existing: []Block{
			blockWithPavilion(2, 10, 1, 1, 1, 2, 3),
		},
		Distances: []Distance{{FromPavilionID: 1, ToPavilionID: 2, Minutes: 90}},
	})

	assertFinding(t, findings, RuleInsufficientTransfer, SeverityWarning)
}

func TestValidateTeachingLoadApprovalWarnsRV05(t *testing.T) {
	findings := ValidateTeachingLoadApproval(TeachingLoadInput{TeacherID: 10, WeeklyHours: 18})

	assertFinding(t, findings, RuleTeachingLoadLimit, SeverityWarning)
}

func TestValidateTeachingLoadApprovalAllowsConfirmedExceptionRV05(t *testing.T) {
	findings := ValidateTeachingLoadApproval(TeachingLoadInput{TeacherID: 10, WeeklyHours: 18, Confirmed: true})

	if len(findings) != 0 {
		t.Fatalf("expected no findings for confirmed exception, got %+v", findings)
	}
}

func TestValidatePlacementInformsSameSeriesConflictRV06(t *testing.T) {
	findings := ValidatePlacement(PlacementInput{
		Proposed: blockWithSeriesCourse(1, 10, 1, 8, 101, 1, 3, 4),
		Existing: []Block{
			blockWithSeriesCourse(2, 20, 2, 8, 102, 1, 4, 5),
		},
	})

	assertFinding(t, findings, RuleSameSeriesConflict, SeverityInfo)
}

func TestValidatePlacementBlocksSharedRoomReservedByAnotherSchoolRV07(t *testing.T) {
	proposed := block(1, 10, 1, 1, 5, 100, 1, 3, 4)
	proposed.RoomShared = true

	existing := block(2, 20, 2, 2, 5, 200, 1, 4, 5)
	existing.RoomShared = true

	findings := ValidatePlacement(PlacementInput{
		Proposed: proposed,
		Existing: []Block{existing},
	})

	assertFinding(t, findings, RuleSharedRoomReservation, SeverityBlocker)
}

func TestValidateAuditChangeRequiresJustificationRV08(t *testing.T) {
	findings := ValidateAuditChange(AuditChangeInput{State: StateOfficial})

	assertFinding(t, findings, RuleAuditJustification, SeverityBlocker)
}

func TestValidateAuditChangeAllowsDraftWithoutJustificationRV08(t *testing.T) {
	findings := ValidateAuditChange(AuditChangeInput{State: StateDraft})

	if len(findings) != 0 {
		t.Fatalf("expected no findings for draft change, got %+v", findings)
	}
}

func TestValidatePlacementCapacityInfoInPreliminaryRV09(t *testing.T) {
	proposed := block(1, 10, 1, 1, 5, 100, 1, 3, 4)
	proposed.Enrollment = 55
	proposed.RoomCapacity = 40

	findings := ValidatePlacement(PlacementInput{Proposed: proposed, State: StatePreliminary})

	assertFinding(t, findings, RuleCapacityReadjustment, SeverityInfo)
}

func TestValidatePlacementCapacityWarningInReadjustmentRV09(t *testing.T) {
	proposed := block(1, 10, 1, 1, 5, 100, 1, 3, 4)
	proposed.Enrollment = 55
	proposed.RoomCapacity = 40

	findings := ValidatePlacement(PlacementInput{Proposed: proposed, State: StateReadjusting})

	assertFinding(t, findings, RuleCapacityReadjustment, SeverityWarning)
}

func TestValidatePlacementAcceptsCleanPlacement(t *testing.T) {
	findings := ValidatePlacement(PlacementInput{
		Proposed: block(1, 10, 1, 1, 1, 100, 1, 6, 7),
		Existing: []Block{
			block(2, 20, 2, 2, 2, 200, 1, 3, 4),
		},
		State: StateDraft,
	})

	if len(findings) != 0 {
		t.Fatalf("expected clean placement, got %+v", findings)
	}
}

func block(id int, teacherID int, schoolID int, groupID int, roomID int, seriesID int, day int, start int, end int) Block {
	return Block{
		ID:        id,
		TeacherID: teacherID,
		SchoolID:  schoolID,
		GroupID:   groupID,
		RoomID:    roomID,
		SeriesID:  seriesID,
		Day:       day,
		StartSlot: start,
		EndSlot:   end,
	}
}

func blockWithDepartment(id int, teacherID int, schoolID int, departmentID int, day int, start int, end int) Block {
	item := block(id, teacherID, schoolID, id, id, id, day, start, end)
	item.DepartmentID = departmentID
	return item
}

func blockWithPavilion(id int, teacherID int, schoolID int, pavilionID int, day int, start int, end int) Block {
	item := block(id, teacherID, schoolID, id, id, id, day, start, end)
	item.PavilionID = pavilionID
	return item
}

func blockWithSeriesCourse(id int, teacherID int, schoolID int, seriesID int, courseID int, day int, start int, end int) Block {
	item := block(id, teacherID, schoolID, id, id, seriesID, day, start, end)
	item.CourseID = courseID
	return item
}

func assertFinding(t *testing.T, findings []Finding, rule RuleID, severity Severity) {
	t.Helper()
	for _, finding := range findings {
		if finding.Rule == rule && finding.Severity == severity {
			return
		}
	}
	t.Fatalf("expected finding %s/%s, got %+v", rule, severity, findings)
}
