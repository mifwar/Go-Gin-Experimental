package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
	registerUseCase "online-course.mifwar.com/internal/register/usecase"
	userDto "online-course.mifwar.com/internal/user/dto"
	"online-course.mifwar.com/pkg/utils"
)

type RegisterHandler struct {
	registerUseCase registerUseCase.RegisterUseCase
}

func NewRegisterHandler(registerUseCase registerUseCase.RegisterUseCase) *RegisterHandler {
	return &RegisterHandler{registerUseCase}
}

func (rh *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("/api/v1/register", rh.Register)
}

func (rh *RegisterHandler) Register(ctx *gin.Context) {
	var registerRequestInput userDto.UserRequestBody

	if err := ctx.ShouldBindJSON(&registerRequestInput); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(400, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	err := rh.registerUseCase.Register(registerRequestInput)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(500, "internal server error", err))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(201, "created", "Success, please check your email"))

}
