package app

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
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
	sendToToken(app)
}

func sendToToken(app *firebase.App) {
	registrationToken := "YOUR_REGISTRATION_TOKEN"
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	//registrationToken := "d3adVu0tMDyBUTPkgc_l-0:APA91bHjb6-wWkT1ABGSasFqxrsOR3AdfcTjLc8b7f7yukWLt32GS4UA5XdIwZ8p98oOLp-CBcyuYaCYdEPRji_f2WSXO9JKb7XPjotm_3bdkk-7hJyxJS8JuUHt82xzGGJ6Aacy0QWb"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification Test",
			Body:  "Hello React!!",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}
