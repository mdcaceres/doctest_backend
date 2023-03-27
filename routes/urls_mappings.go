package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/controllers"
	"github.com/mdcaceres/doctest/middleware"
)

func MapUrls(app *fiber.App) {
	micro := app.Group("/api")
	micro.Post("/auth/register", controllers.Register)
	micro.Post("/auth/login", controllers.Login)
	micro.Get("/auth/logout", middleware.DeserializeUser, controllers.Logout)
	micro.Get("/users/me", middleware.DeserializeUser, controllers.GetMe)
	micro.Get("/ping", controllers.Ping)
	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

}
