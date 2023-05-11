package utils

import (
	"math/rand"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func Paginate(offset, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := offset

		if page <= 0 {
			page = 1
		}

		pageSize := limit

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(pageSize)

	}
}

func GetCurrentUser(ctx *gin.Context) *oauthDto.MapClaimsResponse {
	user, _ := ctx.Get("user")

	return user.(*oauthDto.MapClaimsResponse)
}

func GetFileName(filename string) string {
	file := filepath.Base(filename)

	return file[:len(file)-len(filepath.Ext(file))]
}
