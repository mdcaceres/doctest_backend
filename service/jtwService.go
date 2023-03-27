package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mdcaceres/doctest/domains"
	"github.com/mdcaceres/doctest/domains/auth"
	"os"
	"strconv"
	"time"
)

var (
	jwtPrivateToken = os.Getenv("secret")
)

func GenerateToken(user *domains.User) (string, error) {
	claims := auth.Claims{}
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
