package dto

type SignUpRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Phone    string `form:"phoneNumber"`
	Password string `form:"password" binding:"required"`
}
