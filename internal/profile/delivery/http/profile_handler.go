package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middleware "online-course.mifwar.com/internal/middleware"
	useCase "online-course.mifwar.com/internal/profile/usecase"
	"online-course.mifwar.com/pkg/utils"
)

type ProfileHandler struct {
	useCase useCase.ProfileUseCase
}

func NewProfileHandler(useCase useCase.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{useCase}
}

func (handler *ProfileHandler) Route(r *gin.RouterGroup) {
	authorized := r.Group("/api/v1")

	authorized.Use(middleware.AuthJwt)
	{
		authorized.GET("/profile", handler.GetProfile)
	}
}

func (handler *ProfileHandler) GetProfile(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	// Get profile
	profile, err := handler.useCase.GetProfile(int(user.ID))

	if err != nil {

		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success", profile))
}
