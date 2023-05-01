package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/handlers"
	"github.com/mdcaceres/doctest/middleware"
)

func MapUrls(app *fiber.App) {
	micro := app.Group("/api")
	micro.Post("/auth/signup", handlers.Register)
	micro.Post("/auth/login", handlers.Login)
	micro.Get("/auth/logout", middleware.DeserializeUser, handlers.Logout)
	micro.Get("/users/me", middleware.DeserializeUser, handlers.GetMe)
	micro.Get("/user/:name", middleware.DeserializeUser, handlers.GetByName)
	micro.Post("/projects", middleware.DeserializeUser, handlers.CreateProject)
	micro.Post("/project/invitation", middleware.DeserializeUser, handlers.CreateInvitation)
	micro.Get("/ping", handlers.Ping)
	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

}
