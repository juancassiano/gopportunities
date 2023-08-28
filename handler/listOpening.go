package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juancassiano/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List opening
// @Description List all jobs opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSucess(ctx, "list-openings", openings)
}
