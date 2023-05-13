package user

import (
	"gorm.io/gorm"
	entity "online-course.mifwar.com/internal/user/entity"
)

type UserRepository interface {
	FindAll(offset, limit int) []entity.User
	FindById(id int) (*entity.User, error)
	Count() int
	FindByEmail(email string) (*entity.User, error)
	Create(entity entity.User) (*entity.User, error)
	Update(entity entity.User) (*entity.User, error)
	Delete(entity entity.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// Count implements UserRepository
func (repository *UserRepositoryImpl) Count() int {
	var user entity.User

	var totalUser int64

	repository.db.Model(&user).Count(&totalUser)

	return int(totalUser)
}

// FindByEmail implements UserRepository
func (ur *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll implements UserRepository
func (ur *UserRepositoryImpl) FindAll(offset int, limit int) []entity.User {
	var users []entity.User

	ur.db.Find(&users)

	return users
}

// FindById implements UserRepository
func (ur *UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	var user entity.User

	//can delete "where id" because it's the PK
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Save implements UserRepository
func (ur *UserRepositoryImpl) Create(entity entity.User) (*entity.User, error) {

	if err := ur.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Update implements UserRepository
func (ur *UserRepositoryImpl) Update(entity entity.User) (*entity.User, error) {

	if err := ur.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Delete implements UserRepository
func (ur *UserRepositoryImpl) Delete(entity entity.User) error {
	if err := ur.db.Save(&entity).Error; err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}
