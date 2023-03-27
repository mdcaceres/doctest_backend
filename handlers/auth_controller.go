package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/models"
	"github.com/mdcaceres/doctest/models/auth"
	"github.com/mdcaceres/doctest/services"
	"github.com/mdcaceres/doctest/utils"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var (
	secret      = os.Getenv("secret")
	userService = services.NewUserService()
)

func Register(c *fiber.Ctx) error {
	var payload *auth.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := utils.ValidateStruct(payload)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	userResponse, serviceError := services.NewUserService().Create(c, payload)

	if serviceError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": serviceError.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": userResponse}})
}

func Login(c *fiber.Ctx) error {
	var credentials auth.Credentials

	if err := c.BodyParser(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := utils.ValidateStruct(credentials)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": errors})

	}

	var user models.User

	result := datasource.DB.Where("email = ?", credentials.Email).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": "incorrect user or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(credentials.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "incorrect user or password",
		})
	}

	token, err := services.GenerateToken(&user)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not obtain token",
		})
	}

	cookie := fiber.Cookie{
		Name:     "X-Tiger-Token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
		Domain:   "localhost",
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("X-Tiger-Token")

	token, err := jwt.ParseWithClaims(cookie, &auth.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*auth.Claims)

	var user models.User

	datasource.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "X-Tiger-Token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status":  "Success",
		"message": "log out",
	})
}
