package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mdcaceres/doctest/config"
	"github.com/mdcaceres/doctest/datasource"
	"github.com/mdcaceres/doctest/routes"
	"github.com/mdcaceres/doctest/utils/logs"
	"log"
	"os"
)

func StartApplication() {
	logs.InfoLog.Println("space rocket has lift off")
	datasource.Connect()

	app := fiber.New()

	app.Static("/", "../uploads")
	app.Static("/", "./uploads") // Replace "./public" with the directory where your image files are located
	loggerConfig := logger.Config{
		Output: os.Stdout, // add file to save output
	}

	app.Use("/",
		cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins:     "http://localhost:4200",
			AllowMethods:     "POST, PUT, DELETE, GET, OPTIONS",
		}),

		logger.New(loggerConfig),
	)

	routes.MapUrls(app)

	initFireBase()

	logs.InfoLog.Println("space rocket in orbit")

	log.Fatal(app.Listen(":8080"))
}

func initFireBase() {
	_, _, _ = config.SetupFirebase()
	//services.SendToToken(app, "fcQ5n5Emz0aww8rlO4I4st:APA91bGVkJJcke2ERgxfpdZe0F2KmuiH4xmHm_bUjx1O76y3tjpkUmQ5zdxdbv44cQkqsdsXwlYHxZ53ercngm0P39rP5aL_JDJ3gi-tMrb7vwyU_dS3erAhAFgAHamq7IoWqioRFCGS")
}
