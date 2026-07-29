package schedule

import (
	"encoding/json"
	"net/http"

	"unsch-horarios/backend/internal/schedule/validation"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) ValidatePlacement(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input validation.PlacementInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	findings := validation.ValidatePlacement(input)
	if findings == nil {
		findings = []validation.Finding{}
	}
	respond(w, findings)
}

func (h Handler) ValidateAuditChange(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input validation.AuditChangeInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	findings := validation.ValidateAuditChange(input)
	if findings == nil {
		findings = []validation.Finding{}
	}
	respond(w, findings)
}

func (h Handler) ValidateTeachingLoad(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input validation.TeachingLoadInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	findings := validation.ValidateTeachingLoadApproval(input)
	if findings == nil {
		findings = []validation.Finding{}
	}
	respond(w, findings)
}

func respond[T any](w http.ResponseWriter, data T) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
