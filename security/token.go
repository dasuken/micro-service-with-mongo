package security

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid jwt")
	jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

func NewToken(userId string)(string, error){
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		Issuer: userId,
		IssuedAt: time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}