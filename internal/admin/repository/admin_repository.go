package admin

import (
	"gorm.io/gorm"
	entity "online-course.mifwar.com/internal/admin/entity"
	"online-course.mifwar.com/pkg/utils"
)

type AdminRepository interface {
	FindAll(offset, limit int) []entity.Admin
	FindById(id int) (*entity.Admin, error)
	FindByEmail(email string) (*entity.Admin, error)
	Create(entity entity.Admin) (*entity.Admin, error)
	Update(entity entity.Admin) (*entity.Admin, error)
	Delete(entity entity.Admin) error
}

type AdminRepositoryImpl struct {
	db *gorm.DB
}

// Create implements AdminRepository
func (repository *AdminRepositoryImpl) Create(entity entity.Admin) (*entity.Admin, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Delete implements AdminRepository
func (repository *AdminRepositoryImpl) Delete(entity entity.Admin) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements AdminRepository
func (repository *AdminRepositoryImpl) FindAll(offset, limit int) []entity.Admin {
	var admins []entity.Admin

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&admins)

	return admins
}

// FindByEmail implements AdminRepository
func (repository *AdminRepositoryImpl) FindByEmail(email string) (*entity.Admin, error) {
	var admin entity.Admin

	if err := repository.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

// FindById implements AdminRepository
func (repository *AdminRepositoryImpl) FindById(id int) (*entity.Admin, error) {
	var admin entity.Admin

	if err := repository.db.First(&admin, id).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

// Update implements AdminRepository
func (repository *AdminRepositoryImpl) Update(entity entity.Admin) (*entity.Admin, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{db}
}
