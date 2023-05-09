package utils

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	oauthDto "online-course.mifwar.com/internal/oauth/dto"
)

func GenerateRandomString(number int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, number)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetCurrentUser(ctx *gin.Context) *oauthDto.MapClaimsResponse {
	user := ctx.MustGet("user").(*oauthDto.MapClaimsResponse)
	return user
}
