package order_detail

import (
	"gorm.io/gorm"
	entity "online-course.mifwar.com/internal/order_detail/entity"
)

type OrderDetailRepository interface {
	Create(entity entity.OrderDetail) (*entity.OrderDetail, error)
}

type OrderDetailRepositoryImpl struct {
	db *gorm.DB
}

// Create implements OrderDetailRepository
func (repository *OrderDetailRepositoryImpl) Create(entity entity.OrderDetail) (*entity.OrderDetail, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &OrderDetailRepositoryImpl{db}
}
