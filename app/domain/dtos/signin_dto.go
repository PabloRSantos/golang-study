package dto

import model "go-api/app/domain/models"

type SignInRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignInResponse struct {
	AccessToken string     `json:"accessToken"`
	User        model.User `json:"user"`
}
