//go:build wireinject
// +build wireinject

package register

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course.mifwar.com/internal/register/delivery/http"
	useCase "online-course.mifwar.com/internal/register/usecase"
	userRepository "online-course.mifwar.com/internal/user/repository"
	userUseCase "online-course.mifwar.com/internal/user/usecase"
	mail "online-course.mifwar.com/pkg/mail/sendgrid"
)

func InitializedService(db *gorm.DB) *handler.RegisterHandler {
	wire.Build(
		handler.NewRegisterHandler,
		useCase.NewRegisterUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
		mail.NewMail,
	)

	return &handler.RegisterHandler{}
}
