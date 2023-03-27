package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/utils/logs"
	"gorm.io/gorm"
	"time"
)

type Claims struct {
	gorm.Model
	//TeamId   string `json:"team_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Roles    []models.Role
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
