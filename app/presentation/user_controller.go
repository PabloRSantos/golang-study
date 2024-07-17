package presentation

import (
	dto "go-api/app/domain/dtos"
	service "go-api/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) userController {
	return userController{
		userService,
	}
}

func (uc *userController) SignUp(ctx *gin.Context) {
	var request dto.SignUpRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	err = uc.userService.SignUp(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (uc *userController) SignIn(ctx *gin.Context) {
	var request dto.SignInRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	token, err := uc.userService.SignIn(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (uc *userController) Update(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		response := Response{
			Message: "user id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var request dto.UpdateUserRequest
	err = ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	err = uc.userService.Update(uint(userId), request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
