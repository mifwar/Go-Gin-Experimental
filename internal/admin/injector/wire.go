//go:build wireinject
// +build wireinject

package admin

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course.mifwar.com/internal/admin/delivery/http"
	repository "online-course.mifwar.com/internal/admin/repository"
	usecase "online-course.mifwar.com/internal/admin/usecase"
)

func InitializedService(db *gorm.DB) *handler.AdminHandler {
	wire.Build(
		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
		handler.NewAdminHandler,
	)

	return &handler.AdminHandler{}
}
