package order

import (
	"database/sql"

	"gorm.io/gorm"
	discountEntity "online-course.mifwar.com/internal/discount/entity"
	userEntity "online-course.mifwar.com/internal/user/entity"
)

type Order struct {
	ID           int64                    `json:"id"`
	User         *userEntity.User         `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID       int64                    `json:"user_id"`
	Discount     *discountEntity.Discount `json:"discount" gorm:"foreignKey:DiscountID;references:ID"`
	DiscountID   *int64                   `json:"discount_id"`
	CheckoutLink string                   `json:"checkout_link"`
	ExternalID   string                   `json:"external_id"`
	Price        int64                    `json:"price"`
	TotalPrice   int64                    `json:"total_price"`
	Status       string                   `json:"status"`
	CreatedByID  int64                    `json:"created_by" gorm:"column:created_by"`
	CreatedBy    *userEntity.User         `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID  *int64                   `json:"updated_by"  gorm:"column:updated_by"`
	UpdatedBy    *userEntity.User         `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt    sql.NullTime             `json:"created_at"`
	UpdatedAt    sql.NullTime             `json:"updated_at"`
	DeletedAt    gorm.DeletedAt           `json:"deleted_at"`
}
