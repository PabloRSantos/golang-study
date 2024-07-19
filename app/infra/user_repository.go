package infra

import (
	"errors"
	"fmt"
	model "go-api/app/domain/models"

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

func (ur *UserRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	err := ur.connection.Preload("Events").First(&user, id).Error

	return &user, err
}

func (ur *UserRepository) FindByEmail(email string) *model.User {
	var user model.User
	err := ur.connection.Where(&model.User{Email: email}).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
	}

	return &user
}

func (ur *UserRepository) Create(user *model.User) error {
	result := ur.connection.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}

func (ur *UserRepository) Update(user *model.User) error {
	result := ur.connection.Model(&user).Updates(
		model.User{
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
