package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "online-course.mifwar.com/internal/oauth/dto"
	usecase "online-course.mifwar.com/internal/oauth/usecase"
	"online-course.mifwar.com/pkg/utils"
)

type OauthHandler struct {
	usecase usecase.OauthUseCase
}

func NewOauthHandler(usecase usecase.OauthUseCase) *OauthHandler {
	return &OauthHandler{usecase}
}

func (handler *OauthHandler) Route(r *gin.RouterGroup) {
	oauthRouter := r.Group("/api/v1")

	oauthRouter.POST("/oauth", handler.Login)
}

func (handler *OauthHandler) Login(ctx *gin.Context) {
	var input dto.LoginRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	result, err := handler.usecase.Login(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "success", result))
}
