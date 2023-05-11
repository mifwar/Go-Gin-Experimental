//go:build wireinject
// +build wireinject

package Cart

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course.mifwar.com/internal/cart/delivery/http"
	repository "online-course.mifwar.com/internal/cart/repository"
	usecase "online-course.mifwar.com/internal/cart/usecase"
)

func InitiliazedService(db *gorm.DB) *handler.CartHandler {
	wire.Build(
		handler.NewCartHandler,
		repository.NewCartRepository,
		usecase.NewCartUseCase,
	)

	return &handler.CartHandler{}
}
