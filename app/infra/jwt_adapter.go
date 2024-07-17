package infra

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAdapter struct {
	secret []byte
}

func NewJwtAdapter(secret string) JwtAdapter {
	return JwtAdapter{
		secret: []byte(secret),
	}
}

func (ja *JwtAdapter) Sign(username string) (string, error) {
	oneDayFromNow := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      oneDayFromNow,
		})

	tokenString, err := token.SignedString(ja.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ja *JwtAdapter) Verify(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ja.secret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
