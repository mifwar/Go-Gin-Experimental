//go:build wireinject
// +build wireinject

package order

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	cartRepository "online-course.mifwar.com/internal/cart/repository"
	cartUseCase "online-course.mifwar.com/internal/cart/usecase"
	discountRepository "online-course.mifwar.com/internal/discount/repository"
	discountUseCase "online-course.mifwar.com/internal/discount/usecase"
	handler "online-course.mifwar.com/internal/order/delivery/http"
	repository "online-course.mifwar.com/internal/order/repository"
	useCase "online-course.mifwar.com/internal/order/usecase"
	orderDetailRepository "online-course.mifwar.com/internal/order_detail/repository"
	orderDetailUseCase "online-course.mifwar.com/internal/order_detail/usecase"
	paymentUseCase "online-course.mifwar.com/internal/payment/usecase"
	productRepository "online-course.mifwar.com/internal/product/repository"
	productUseCase "online-course.mifwar.com/internal/product/usecase"
	fileUpload "online-course.mifwar.com/pkg/fileupload/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.OrderHandler {
	wire.Build(
		cartRepository.NewCartRepository,
		cartUseCase.NewCartUseCase,
		discountRepository.NewDiscountRepository,
		discountUseCase.NewDiscountUseCase,
		handler.NewOrderHandler,
		repository.NewOrderRepository,
		useCase.NewOrderUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUseCase.NewOrderDetailUseCase,
		paymentUseCase.NewPaymentUseCase,
		productRepository.NewProductRepository,
		productUseCase.NewProductUseCase,
		fileUpload.NewFileUpload,
	)

	return &handler.OrderHandler{}
}
