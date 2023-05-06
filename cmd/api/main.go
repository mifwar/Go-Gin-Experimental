package main

import (
	"github.com/gin-gonic/gin"
	mysql "online-course.mifwar.com/pkg/db/mysql"

	registerHandler "online-course.mifwar.com/internal/register/delivery/http"
	registerUseCase "online-course.mifwar.com/internal/register/usecase"
	userRepository "online-course.mifwar.com/internal/user/repository"
	userUseCase "online-course.mifwar.com/internal/user/usecase"
	mail "online-course.mifwar.com/pkg/mail/sendgrid"
)

func main() {

	db := mysql.DB()

	r := gin.Default()

	mail := mail.NewMail()
	userRepository := userRepository.NewUserRepositoryImpl(db)
	userUseCase := userUseCase.NewUserUseCase(userRepository)
	registerUseCase := registerUseCase.NewRegisterUseCase(userUseCase, mail)
	registerHandler.NewRegisterHandler(registerUseCase).Route(&r.RouterGroup)

	r.Run("localhost:8080")
}
