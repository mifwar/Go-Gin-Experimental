//go:build wireinject
// +build wireinject

package webhook

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	cartRepository "online-course.mifwar.com/internal/cart/repository"
	cartUseCase "online-course.mifwar.com/internal/cart/usecase"
	classRoomRepository "online-course.mifwar.com/internal/class_room/repository"
	classRoomUseCase "online-course.mifwar.com/internal/class_room/usecase"
	discountRepository "online-course.mifwar.com/internal/discount/repository"
	discountUseCase "online-course.mifwar.com/internal/discount/usecase"
	orderRepository "online-course.mifwar.com/internal/order/repository"
	orderUseCase "online-course.mifwar.com/internal/order/usecase"
	orderDetailRepository "online-course.mifwar.com/internal/order_detail/repository"
	orderDetailUseCase "online-course.mifwar.com/internal/order_detail/usecase"
	paymentUseCase "online-course.mifwar.com/internal/payment/usecase"
	productRepository "online-course.mifwar.com/internal/product/repository"
	productUseCase "online-course.mifwar.com/internal/product/usecase"
	handler "online-course.mifwar.com/internal/webhook/delivery/http"
	useCase "online-course.mifwar.com/internal/webhook/usecase"
	fileUpload "online-course.mifwar.com/pkg/fileupload/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.WebhookHandler {
	wire.Build(
		handler.NewWebHookHandler,
		useCase.NewWebhookUseCase,
		classRoomRepository.NewClassRoomRepository,
		classRoomUseCase.NewClassRoomUseCase,
		orderRepository.NewOrderRepository,
		orderUseCase.NewOrderUseCase,
		cartRepository.NewCartRepository,
		cartUseCase.NewCartUseCase,
		discountRepository.NewDiscountRepository,
		discountUseCase.NewDiscountUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUseCase.NewOrderDetailUseCase,
		paymentUseCase.NewPaymentUseCase,
		productRepository.NewProductRepository,
		productUseCase.NewProductUseCase,
		fileUpload.NewFileUpload,
	)

	return &handler.WebhookHandler{}
}
