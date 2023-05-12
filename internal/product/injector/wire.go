//go:build wireinject
// +build wireinject

package product

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course.mifwar.com/internal/product/delivery/http"
	repository "online-course.mifwar.com/internal/product/repository"
	usecase "online-course.mifwar.com/internal/product/usecase"
	fileUpload "online-course.mifwar.com/pkg/fileupload/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.ProductHandler {
	wire.Build(
		handler.NewProductHandler,
		usecase.NewProductUseCase,
		repository.NewProductRepository,
		fileUpload.NewFileUpload,
	)

	return &handler.ProductHandler{}
}
