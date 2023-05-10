package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"online-course.mifwar.com/pkg/utils"
)

func AuthAdmin(ctx *gin.Context) {
	admin := utils.GetCurrentUser(ctx)

	if !admin.IsAdmin {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "Unauthorized", "Unauthorized"))
		ctx.Abort()
		return
	}

	ctx.Next()
}
