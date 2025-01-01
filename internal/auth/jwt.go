package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(userID int64) string {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(jwtKey)
	return signedToken
}
