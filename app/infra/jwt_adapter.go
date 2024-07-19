package infra

import (
	"fmt"
	model "go-api/app/domain/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAdapter struct {
	secret []byte
}

func NewJwtAdapter() JwtAdapter {
	return JwtAdapter{
		secret: []byte(os.Getenv("TOKEN_SECRET")),
	}
}

func (ja *JwtAdapter) Sign(payload model.TokenClaims) (string, error) {
	oneDayFromNow := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"ID":   payload.ID,
			"Role": payload.Role,
			"exp":  oneDayFromNow,
		})

	tokenString, err := token.SignedString(ja.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ja *JwtAdapter) Verify(tokenString string) (model.TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ja.secret, nil
	})

	if err != nil {
		return model.TokenClaims{}, err
	}

	if !token.Valid {
		return model.TokenClaims{}, fmt.Errorf("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)

	role, err := getRole(claims["Role"])
	if err != nil {
		return model.TokenClaims{}, err
	}

	parsedClaims := model.TokenClaims{
		ID:   uint(claims["ID"].(float64)),
		Role: role,
	}

	return parsedClaims, err
}

func getRole(role interface{}) (model.Role, error) {
	roleStr := role.(string)

	if string(model.ADMIN_ROLE) == roleStr {
		return model.ADMIN_ROLE, nil
	}

	if string(model.USER_ROLE) == roleStr {
		return model.USER_ROLE, nil
	}

	return "", fmt.Errorf("unknown role: %s", roleStr)
}
