package order

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	cartUseCase "online-course.mifwar.com/internal/cart/usecase"
	discountEntity "online-course.mifwar.com/internal/discount/entity"
	discountUseCase "online-course.mifwar.com/internal/discount/usecase"
	dto "online-course.mifwar.com/internal/order/dto"
	entity "online-course.mifwar.com/internal/order/entity"
	repository "online-course.mifwar.com/internal/order/repository"
	orderDetailEntity "online-course.mifwar.com/internal/order_detail/entity"
	orderDetailUseCase "online-course.mifwar.com/internal/order_detail/usecase"
	paymentDto "online-course.mifwar.com/internal/payment/dto"
	paymentUseCase "online-course.mifwar.com/internal/payment/usecase"
	productEntity "online-course.mifwar.com/internal/product/entity"
	productUseCase "online-course.mifwar.com/internal/product/usecase"
)

type OrderUseCase interface {
	FindAll(offset, limit int) []entity.Order
	FindById(id int) (*entity.Order, error)
	Create(dto dto.OrderRequestBody) (*entity.Order, error)
}

type OrderUseCaseImpl struct {
	repository         repository.OrderRepository
	cartUseCase        cartUseCase.CartUseCase
	discountUseCase    discountUseCase.DiscountUseCase
	productUseCase     productUseCase.ProductUseCase
	orderDetailUseCase orderDetailUseCase.OrderDetailUseCase
	paymentUseCase     paymentUseCase.PaymentUseCase
}

// Create implements OrderUseCase
func (usecase *OrderUseCaseImpl) Create(dto dto.OrderRequestBody) (*entity.Order, error) {
	price := 0
	totalPrice := 0
	description := ""

	var products []productEntity.Product

	order := entity.Order{
		UserID: dto.UserID,
		Status: "pending",
	}

	var dataDiscount *discountEntity.Discount

	carts := usecase.cartUseCase.FindByUserId(int(dto.UserID), 0, 9999)

	if len(carts) == 0 {
		if dto.ProductID == nil {
			return nil, errors.New("cart anda kosong atau anda belum memasukkan product id")
		}
	}

	// Check data discount
	if dto.DiscountCode != nil {
		discount, err := usecase.discountUseCase.FindByCode(*dto.DiscountCode)

		if err != nil {
			return nil, errors.New("code discount sudah tidak berlaku")
		}

		if discount.RemainingQuantity == 0 {
			return nil, errors.New("code discount sudah tidak berlaku")
		}
		// Validasi lainnya misalnya check start date dan end date

		dataDiscount = discount
		fmt.Print(dataDiscount)
	}

	if len(carts) > 0 {
		for _, cart := range carts {
			product, err := usecase.productUseCase.FindById(int(cart.ProductID))

			if err != nil {
				return nil, err
			}

			products = append(products, *product)
		}
	} else if dto.ProductID != nil {
		// Jika user langsung melakukan checkout
		product, err := usecase.productUseCase.FindById(int(*dto.ProductID))

		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	for index, product := range products {
		price += int(product.Price)

		i := strconv.Itoa(index + 1)

		description = i + ". Product : " + product.Title + "<br/>"
	}

	totalPrice = price
	if dataDiscount != nil {
		if dataDiscount.Type == "rebate" {
			totalPrice = price - int(dataDiscount.Value)
		} else if dataDiscount.Type == "percent" {
			totalPrice = price - (price / 100 * int(dataDiscount.Value))
		}

		order.DiscountID = &dataDiscount.ID
	}

	order.Price = int64(price)           // Harga asli
	order.TotalPrice = int64(totalPrice) // Harga yang sudah dikurangi discount
	order.CreatedByID = dto.UserID

	externalId := uuid.New().String()

	order.ExternalID = externalId

	data, err := usecase.repository.Create(order)

	if err != nil {
		return nil, err
	}

	for _, product := range products {
		orderDetail := orderDetailEntity.OrderDetail{
			OrderID:     data.ID,
			ProductID:   product.ID,
			CreatedByID: order.UserID,
			Price:       product.Price,
		}

		usecase.orderDetailUseCase.Create(orderDetail)
	}

	// Hit payment xendit
	dataPayment := paymentDto.PaymentRequestBody{
		ExternalID:  externalId,
		Amount:      int64(totalPrice),
		PayerEmail:  dto.Email,
		Description: description,
	}

	payment, err := usecase.paymentUseCase.Create(dataPayment)

	if err != nil {
		return nil, err
	}

	data.CheckoutLink = payment.InvoiceURL

	usecase.repository.Update(*data)

	// Update remaining quantity discount
	if dto.DiscountCode != nil {
		_, err := usecase.discountUseCase.UpdateRemainingQuantity(int(dataDiscount.ID), 1, "-")

		if err != nil {
			return nil, err
		}
	}

	// Delete carts
	err = usecase.cartUseCase.DeleteByUserId(int(dto.UserID))

	if err != nil {
		return nil, err
	}

	return data, nil
}

// FindAll implements OrderUseCase
func (usecase *OrderUseCaseImpl) FindAll(offset int, limit int) []entity.Order {
	return usecase.repository.FindAll(offset, limit)
}

// FindById implements OrderUseCase
func (usecase *OrderUseCaseImpl) FindById(id int) (*entity.Order, error) {
	return usecase.repository.FindById(id)
}

func NewOrderUseCase(
	repository repository.OrderRepository,
	cartUseCase cartUseCase.CartUseCase,
	discountUseCase discountUseCase.DiscountUseCase,
	productUseCase productUseCase.ProductUseCase,
	orderDetailUseCase orderDetailUseCase.OrderDetailUseCase,
	paymentUseCase paymentUseCase.PaymentUseCase,
) OrderUseCase {
	return &OrderUseCaseImpl{
		repository,
		cartUseCase,
		discountUseCase,
		productUseCase,
		orderDetailUseCase,
		paymentUseCase,
	}
}
