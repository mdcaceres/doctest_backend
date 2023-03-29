package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mdcaceres/doctest/config"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/routes"
	"github.com/mdcaceres/doctest/services"
	"github.com/mdcaceres/doctest/utils/logs"
	"log"
	"os"
)

func StartApplication() {
	logs.InfoLog.Println("space rocket has lift off")
	datasource.Connect()

	app := fiber.New()

	loggerConfig := logger.Config{
		Output: os.Stdout, // add file to save output
	}

	app.Use(
		cors.New(cors.Config{
			AllowCredentials: true}),
		/*
			AllowOrigins:     "http://localhost:3000",
			AllowHeaders:     "Origin, Content-Type, Accept",
			AllowMethods:     "GET, POST",
			AllowCredentials: true,
		*/

		logger.New(loggerConfig),
	)

	routes.MapUrls(app)

	log.Fatal(app.Listen(":8080"))

	logs.InfoLog.Println("space rocket in orbit")
}

func initFireBase() {
	app, _, _ := config.SetupFirebase()
	services.SendToToken(app, "token")
}
