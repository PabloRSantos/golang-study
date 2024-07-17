package dto

type UpdateUserRequest struct {
	Name  string `form:"name" binding:"required"`
	Phone string `form:"phoneNumber"`
}
