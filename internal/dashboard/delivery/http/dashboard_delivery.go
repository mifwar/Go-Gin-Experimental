package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	useCase "online-course.mifwar.com/internal/dashboard/usecase"
	"online-course.mifwar.com/internal/middleware"
	"online-course.mifwar.com/pkg/utils"
)

type DashboardHandler struct {
	useCase useCase.DashboardUseCase
}

func NewDashboardHandler(useCase useCase.DashboardUseCase) *DashboardHandler {
	return &DashboardHandler{useCase}
}

func (handler *DashboardHandler) Route(r *gin.RouterGroup) {
	dashboardHandler := r.Group("/api/v1")

	dashboardHandler.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		dashboardHandler.GET("/dashboards", handler.GetDataDashboard)
	}
}

func (handler *DashboardHandler) GetDataDashboard(ctx *gin.Context) {
	data := handler.useCase.GetDataDashboard()

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
