package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	dto "online-course.mifwar.com/internal/user/dto"
	entity "online-course.mifwar.com/internal/user/entity"
	repository "online-course.mifwar.com/internal/user/repository"
	utils "online-course.mifwar.com/pkg/utils"
)

type UserUseCase interface {
	FindAll(offset, limit int) []entity.User
	FindById(id int) (*entity.User, error)
	Create(userDto dto.UserRequestBody) (*entity.User, error)
	Update(userDto dto.UserRequestBody) (*entity.User, error)
	Delete(id int) error
}

type UserUseCaseImpl struct {
	repository repository.UserRepository
}

// Create implements UserUseCase
func (uu *UserUseCaseImpl) Create(userDto dto.UserRequestBody) (*entity.User, error) {

	checkUser, err := uu.repository.FindByEmail(*userDto.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if checkUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := entity.User{
		Name:         *userDto.Name,
		Email:        *userDto.Email,
		Password:     string(hashedPassword),
		CodeVerified: utils.GenerateRandomString(32),
	}

	dataUser, err := uu.repository.Create(user)

	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

// Delete implements UserUseCase
func (*UserUseCaseImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements UserUseCase
func (*UserUseCaseImpl) FindAll(offset int, limit int) []entity.User {
	panic("unimplemented")
}

// FindById implements UserUseCase
func (*UserUseCaseImpl) FindById(id int) (*entity.User, error) {
	panic("unimplemented")
}

// Update implements UserUseCase
func (*UserUseCaseImpl) Update(userDto dto.UserRequestBody) (*entity.User, error) {
	panic("unimplemented")
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{repository}
}
