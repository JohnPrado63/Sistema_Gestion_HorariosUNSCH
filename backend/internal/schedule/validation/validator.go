package validation

import "strings"

const slotMinutes = 60

func ValidatePlacement(input PlacementInput) []Finding {
	var findings []Finding
	proposed := input.Proposed

	for _, existing := range input.Existing {
		if proposed.Day != existing.Day || !slotsOverlap(proposed, existing) {
			findings = append(findings, transferFindings(proposed, existing, input.Distances)...)
			continue
		}

		if proposed.TeacherID != 0 && proposed.TeacherID == existing.TeacherID {
			findings = append(findings, Finding{
				Rule:     RuleTeacherConflict,
				Severity: SeverityBlocker,
				Message:  "el docente ya tiene una clase asignada en el mismo dia y rango horario",
			})
		}

		if proposed.RoomID != 0 && proposed.RoomID == existing.RoomID {
			if proposed.RoomShared && proposed.SchoolID != existing.SchoolID {
				findings = append(findings, Finding{
					Rule:     RuleSharedRoomReservation,
					Severity: SeverityBlocker,
					Message:  "el aula compartida ya fue reservada por otra escuela en ese horario",
				})
			} else {
				findings = append(findings, Finding{
					Rule:     RuleRoomConflict,
					Severity: SeverityBlocker,
					Message:  "el aula ya tiene una clase asignada en el mismo horario",
				})
			}
		}

		if proposed.SeriesID != 0 && proposed.SeriesID == existing.SeriesID && proposed.CourseID != existing.CourseID {
			findings = append(findings, Finding{
				Rule:     RuleSameSeriesConflict,
				Severity: SeverityInfo,
				Message:  "dos asignaturas de la misma serie coinciden en el mismo horario",
			})
		}
	}

	for _, session := range input.DepartmentSessions {
		if proposed.DepartmentID == session.DepartmentID && proposed.Day == session.Day && rangesOverlap(proposed.StartSlot, proposed.EndSlot, session.StartSlot, session.EndSlot) {
			findings = append(findings, Finding{
				Rule:     RuleDepartmentSession,
				Severity: SeverityBlocker,
				Message:  "la clase coincide con la sesion semanal reservada del departamento",
			})
		}
	}

	if proposed.Enrollment > proposed.RoomCapacity && proposed.RoomCapacity > 0 {
		switch input.State {
		case StatePreliminary:
			findings = append(findings, Finding{
				Rule:     RuleCapacityReadjustment,
				Severity: SeverityInfo,
				Message:  "la matricula proyectada excede el aforo del aula en estado preliminar",
			})
		case StateReadjusting:
			findings = append(findings, Finding{
				Rule:     RuleCapacityReadjustment,
				Severity: SeverityWarning,
				Message:  "la matricula real excede el aforo del aula; se recomienda cambiar de aula o abrir otro grupo",
			})
		}
	}

	return findings
}

func ValidateAuditChange(input AuditChangeInput) []Finding {
	if (input.State == StateOfficial || input.State == StateReadjusted) && strings.TrimSpace(input.Justification) == "" {
		return []Finding{{
			Rule:     RuleAuditJustification,
			Severity: SeverityBlocker,
			Message:  "los cambios sobre horarios oficiales o reajustados requieren justificacion obligatoria",
		}}
	}
	return nil
}

func ValidateTeachingLoadApproval(input TeachingLoadInput) []Finding {
	if input.WeeklyHours <= 16 || input.Confirmed {
		return nil
	}

	return []Finding{{
		Rule:     RuleTeachingLoadLimit,
		Severity: SeverityWarning,
		Message:  "la carga lectiva semanal del docente supera 16 horas y requiere confirmacion documentada de la DGA",
	}}
}

func transferFindings(proposed Block, existing Block, distances []Distance) []Finding {
	if proposed.Day != existing.Day || proposed.TeacherID == 0 || proposed.TeacherID != existing.TeacherID {
		return nil
	}
	if proposed.PavilionID == 0 || existing.PavilionID == 0 || proposed.PavilionID == existing.PavilionID {
		return nil
	}

	gapSlots, from, to, ok := gapBetween(existing, proposed)
	if !ok {
		return nil
	}

	distance := distanceMinutes(from, to, distances)
	if distance <= 0 {
		return nil
	}

	if gapSlots == 0 {
		return []Finding{{
			Rule:     RuleNoTransferTime,
			Severity: SeverityBlocker,
			Message:  "el docente tiene clases consecutivas en pabellones distintos sin intervalo libre",
		}}
	}

	if gapSlots*slotMinutes < distance {
		return []Finding{{
			Rule:     RuleInsufficientTransfer,
			Severity: SeverityWarning,
			Message:  "el intervalo libre entre pabellones es menor al tiempo estimado de desplazamiento",
		}}
	}

	return nil
}

func slotsOverlap(a Block, b Block) bool {
	return rangesOverlap(a.StartSlot, a.EndSlot, b.StartSlot, b.EndSlot)
}

func rangesOverlap(aStart, aEnd, bStart, bEnd int) bool {
	return aStart <= bEnd && bStart <= aEnd
}

func gapBetween(existing Block, proposed Block) (gapSlots int, fromPavilionID int, toPavilionID int, ok bool) {
	if existing.EndSlot < proposed.StartSlot {
		return proposed.StartSlot - existing.EndSlot - 1, existing.PavilionID, proposed.PavilionID, true
	}
	if proposed.EndSlot < existing.StartSlot {
		return existing.StartSlot - proposed.EndSlot - 1, proposed.PavilionID, existing.PavilionID, true
	}
	return 0, 0, 0, false
}

func distanceMinutes(from int, to int, distances []Distance) int {
	for _, distance := range distances {
		if distance.FromPavilionID == from && distance.ToPavilionID == to {
			return distance.Minutes
		}
	}
	return 0
}
