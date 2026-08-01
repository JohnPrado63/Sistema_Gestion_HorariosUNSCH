package schedule

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestValidatePlacementHandlerReturnsFindings(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewHandler()
	body := `{"proposed":{"teacher_id":10,"day":1,"start_slot":3,"end_slot":4,"room_id":5,"school_id":1},"existing":[{"teacher_id":10,"day":1,"start_slot":4,"end_slot":5,"room_id":6,"school_id":1}],"state":"BORRADOR"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/validaciones/placement", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	h.ValidatePlacement(c)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "RV-01") {
		t.Fatalf("expected RV-01 finding, got %s", rr.Body.String())
	}
}

func TestValidateAuditChangeHandlerReturnsEmptyArrayWhenNoFindings(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewHandler()
	body := `{"state":"BORRADOR","justification":""}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/validaciones/audit", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	h.ValidateAuditChange(c)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	if strings.TrimSpace(rr.Body.String()) != "[]" {
		t.Fatalf("expected empty array, got %s", rr.Body.String())
	}
}

func TestValidateTeachingLoadHandlerWarnsOverload(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewHandler()
	body := `{"teacher_id":10,"weekly_hours":18,"confirmed":false}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/validaciones/carga", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	h.ValidateTeachingLoad(c)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "RV-05") {
		t.Fatalf("expected RV-05 finding, got %s", rr.Body.String())
	}
}