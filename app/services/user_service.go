package service

import (
	"errors"
	"fmt"
	dto "go-api/app/domain/dtos"
	model "go-api/app/domain/models"
	"go-api/app/infra"
)

type UserService struct {
	repository   infra.UserRepository
	cryptography infra.BcryptAdapter
	token        infra.JwtAdapter
}

func NewUserService(repository infra.UserRepository, cryptography infra.BcryptAdapter, token infra.JwtAdapter) UserService {
	return UserService{
		repository,
		cryptography,
		token,
	}
}

func (us *UserService) SignUp(payload dto.SignUpRequest) error {
	if us.repository.FindByEmail(payload.Email) != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := us.cryptography.Hash(payload.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Phone:    payload.Phone,
		Password: hashedPassword,
	}

	err = us.repository.Create(&user)
	return err
}

func (us *UserService) SignIn(payload dto.SignInRequest) (dto.SignInResponse, error) {
	user := us.repository.FindByEmail(payload.Email)

	if user == nil {
		return dto.SignInResponse{}, errors.New("invalid credentials")
	}

	passwordMatch := us.cryptography.Compare(payload.Password, user.Password)
	if !passwordMatch {
		return dto.SignInResponse{}, fmt.Errorf("invalid credentials")
	}

	token, err := us.token.Sign(user.Email)
	if err != nil {
		return dto.SignInResponse{}, err
	}

	response := dto.SignInResponse{
		AccessToken: token,
		User:        *user,
	}

	return response, nil
}

func (us *UserService) Update(id uint, payload dto.UpdateUserRequest) error {
	user := us.repository.FindById(id)
	if user == nil {
		return errors.New("user not exists")
	}

	user.Name = payload.Name
	user.Phone = payload.Phone
	us.repository.Update(user)

	return nil
}
