package utility

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/entity"
	model "github.com/salamanderman234/peripheral-api/models"
)

func createJwtClaims(user model.User) entity.JWTClaims {
	claims := entity.JWTClaims{
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Hour)),
		},
	}
	return claims
}
func CreateJwtToken(user model.User) (string, error) {
	claims := createJwtClaims(user)
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte(config.GetAppSecretKey()))
	if err != nil {
		return "", err
	}
	return token, nil
}
