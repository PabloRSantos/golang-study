package presentation

import (
	"fmt"
	dto "go-api/app/domain/dtos"
	model "go-api/app/domain/models"
	"go-api/app/infra"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	token        infra.JwtAdapter
	requiredRole model.Role
}

func NewAuthMiddleware(token infra.JwtAdapter, requiredRole model.Role) AuthMiddleware {
	return AuthMiddleware{
		token,
		requiredRole,
	}
}

func (am *AuthMiddleware) Verify(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		err := fmt.Errorf("authorization header required")
		ctx.JSON(http.StatusUnauthorized, dto.NewErrorResponse(err))
		ctx.Abort()
		return
	}

	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		err := fmt.Errorf("invalid token format")
		ctx.JSON(http.StatusUnauthorized, dto.NewErrorResponse(err))
		ctx.Abort()
		return
	}

	token := parts[1]
	claims, err := am.token.Verify(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.NewErrorResponse(err))
		ctx.Abort()
		return
	}

	availableOnlyForAdmin := am.requiredRole == "ADMIN"
	userIsAdmin := claims.Role == "ADMIN"
	if availableOnlyForAdmin && !userIsAdmin {
		err := fmt.Errorf("resource not available for your user")
		ctx.JSON(http.StatusUnauthorized, dto.NewErrorResponse(err))
		ctx.Abort()
		return
	}

	ctx.Set("x-user-id", claims.ID)
	ctx.Set("x-user-role", claims.Role)
	ctx.Next()
}
