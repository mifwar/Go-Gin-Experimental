//go:build wireinject
// +build wireinject

package product_category

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course.mifwar.com/internal/product_category/delivery/http"
	repository "online-course.mifwar.com/internal/product_category/repository"
	usecase "online-course.mifwar.com/internal/product_category/usecase"
)

func InitializedService(db *gorm.DB) *handler.ProductCategoryHandler {
	wire.Build(
		handler.NewProductCategoryHandler,
		repository.NewProductCategoryRepository,
		usecase.NewProductCategoryUseCase,
	)

	return &handler.ProductCategoryHandler{}
}
