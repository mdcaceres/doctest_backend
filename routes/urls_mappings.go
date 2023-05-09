package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mdcaceres/doctest/handlers"
	"github.com/mdcaceres/doctest/middleware"
)

func MapUrls(app *fiber.App) {
	app.Get("/img/:id", handlers.Serve)
	micro := app.Group("/api")
	micro.Post("/auth/signup", handlers.Register)
	micro.Post("/auth/login", handlers.Login)
	micro.Get("/auth/logout", middleware.DeserializeUser, handlers.Logout)
	micro.Get("/users/me", middleware.DeserializeUser, handlers.GetMe)
	micro.Get("/user/:name", middleware.DeserializeUser, handlers.GetUserByName)
	micro.Get("/user/id/:id", middleware.DeserializeUser, handlers.GetUserById)
	micro.Get("/projects", middleware.DeserializeUser, handlers.GetProjects)
	micro.Get("/project/id/:id", middleware.DeserializeUser, handlers.GetProject)
	micro.Put("/user/:id", middleware.DeserializeUser, handlers.UpdateToken)
	micro.Post("/projects", middleware.DeserializeUser, handlers.CreateProject)
	micro.Put("/project/:id/img", middleware.DeserializeUser, handlers.UploadProjectImage)
	micro.Post("/project/invitation", middleware.DeserializeUser, handlers.CreateInvitation)
	micro.Put("/project/join", middleware.DeserializeUser, handlers.JoinProject)
	micro.Get("/ping", handlers.Ping)
	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

}
