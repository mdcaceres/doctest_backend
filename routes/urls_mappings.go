package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
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
	micro.Get("/users/project/:id", middleware.DeserializeUser, handlers.GetAllByProject)
	micro.Get("/projects", middleware.DeserializeUser, handlers.GetProjects)
	micro.Get("/project/id/:id", middleware.DeserializeUser, handlers.GetProject)
	micro.Put("/user/:id", middleware.DeserializeUser, handlers.UpdateToken)
	micro.Post("/projects", middleware.DeserializeUser, handlers.CreateProject)
	micro.Put("/project/:id/img", middleware.DeserializeUser, handlers.UploadProjectImage)
	micro.Post("/test/:id/files", middleware.DeserializeUser, handlers.UploadFileToCase)
	micro.Get("/test/:id", middleware.DeserializeUser, handlers.GetCaseById)
	micro.Get("/test/user/:id", middleware.DeserializeUser, handlers.GetAllByUserId)
	micro.Post("bug/:id/files", middleware.DeserializeUser, handlers.UploadFileToBug)
	micro.Post("/test/:id/execution", middleware.DeserializeUser, handlers.ExecuteTest)
	micro.Post("/clients/:userId", middleware.DeserializeUser, handlers.CreateClient)
	micro.Get("/clients/:userId", middleware.DeserializeUser, handlers.GetClients)
	micro.Post("/project/invitation", middleware.DeserializeUser, handlers.CreateInvitation)
	micro.Get("/project/join", handlers.JoinProject)
	micro.Post("/project/:id/suite", middleware.DeserializeUser, handlers.CreateSuite)
	micro.Get("/project/:id/suites", middleware.DeserializeUser, handlers.GetSuites)
	micro.Get("/project/:id/tests", middleware.DeserializeUser, handlers.GetAllCasesByProjectId)
	micro.Get("/project/:id/bugs", middleware.DeserializeUser, handlers.GetAllBugsByProjectId)
	micro.Post("/project/:id/bug", middleware.DeserializeUser, handlers.CreateBug)
	micro.Put("/project/:id/bug/update", middleware.DeserializeUser, handlers.UpdateBug)
	micro.Post("/project/:id/test", middleware.DeserializeUser, handlers.CreateCase)
	micro.Get("/user/:id/bugs/status/:status", middleware.DeserializeUser, handlers.GetBugByUserId)
	micro.Post("/email", middleware.DeserializeUser, handlers.SendSimple)
	micro.Post("/bug/comment", middleware.DeserializeUser, handlers.AddBugComment)
	micro.Post("/post", middleware.DeserializeUser, handlers.CreatePost)
	micro.Get("/post/project/:id", middleware.DeserializeUser, handlers.GetAllPostsByProjectId)
	micro.Get("/ping", handlers.Ping)
	micro.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

}
