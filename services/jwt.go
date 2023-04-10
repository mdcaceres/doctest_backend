package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/auth"
	"os"
	"strconv"
	"time"
)

var (
	jwtPrivateToken = os.Getenv("secret")
)

func GenerateToken(user *models.User) (string, error) {
	claims := auth.Claims{}
	claims.ID = user.ID
	claims.ExpiresAt = time.Now().Add(time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = strconv.Itoa(int(user.ID))
	claims.UserName = user.Name
	claims.Roles = user.Roles

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtPrivateToken))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
