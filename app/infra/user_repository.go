package infra

import (
	"errors"
	"fmt"
	domain "go-api/app/domain/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) FindById(id uint) *domain.User {
	var user domain.User
	ur.connection.First(&user, id)

	return &user
}

func (ur *UserRepository) FindByEmail(email string) *domain.User {
	var user domain.User
	err := ur.connection.Where(&domain.User{Email: email}).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
	}

	return &user
}

func (ur *UserRepository) Create(user *domain.User) error {
	result := ur.connection.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}

func (ur *UserRepository) Update(user *domain.User) error {
	result := ur.connection.Model(&user).Updates(
		domain.User{
			Name:  user.Name,
			Phone: user.Phone,
		},
	)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}
