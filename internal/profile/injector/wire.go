//go:build wireinject
// +build wireinject

package profile

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course.mifwar.com/internal/profile/delivery/http"
	useCase "online-course.mifwar.com/internal/profile/usecase"
	userRepository "online-course.mifwar.com/internal/user/repository"
	userUseCase "online-course.mifwar.com/internal/user/usecase"
)

func InitializedService(db *gorm.DB) *handler.ProfileHandler {
	wire.Build(
		handler.NewProfileHandler,
		useCase.NewProfileUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &handler.ProfileHandler{}
}
