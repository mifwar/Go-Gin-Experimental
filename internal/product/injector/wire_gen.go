// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package product

import (
	"gorm.io/gorm"
	"online-course.mifwar.com/internal/product/delivery/http"
	product2 "online-course.mifwar.com/internal/product/repository"
	product3 "online-course.mifwar.com/internal/product/usecase"
	"online-course.mifwar.com/pkg/fileupload/cloudinary"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *product.ProductHandler {
	productRepository := product2.NewProductRepository(db)
	fileUpload := fileupload.NewFileUpload()
	productUseCase := product3.NewProductUseCase(productRepository, fileUpload)
	productHandler := product.NewProductHandler(productUseCase)
	return productHandler
}
