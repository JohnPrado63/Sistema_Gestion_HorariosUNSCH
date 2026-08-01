package schedule

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"unsch-horarios/backend/internal/schedule/validation"
)

type Handler struct{}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) ValidatePlacement(c *gin.Context) {
	var input validation.PlacementInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findings := validation.ValidatePlacement(input)
	if findings == nil {
		findings = []validation.Finding{}
	}
	c.JSON(http.StatusOK, findings)
}

func (h Handler) ValidateAuditChange(c *gin.Context) {
	var input validation.AuditChangeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findings := validation.ValidateAuditChange(input)
	if findings == nil {
		findings = []validation.Finding{}
	}
	c.JSON(http.StatusOK, findings)
}

func (h Handler) ValidateTeachingLoad(c *gin.Context) {
	var input validation.TeachingLoadInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findings := validation.ValidateTeachingLoadApproval(input)
	if findings == nil {
		findings = []validation.Finding{}
	}
	c.JSON(http.StatusOK, findings)
}
