package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mdcaceres/doctest/domains"
	"github.com/mdcaceres/doctest/utils/logs"
	"gorm.io/gorm"
	"time"
)

type Claims struct {
	//TeamId   string `json:"team_id,omitempty"`
	gorm.Model
	UserName string `json:"user_name,omitempty"`
	Roles    []domains.Role
	jwt.StandardClaims
}

func (claims Claims) Valid() error {
	now := time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) {
		return nil
	}
	logs.ErrorLog.Println("invalid token")
	return fmt.Errorf("token is invalid")
}
