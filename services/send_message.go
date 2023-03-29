package services

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/mdcaceres/doctest/models"
	"log"
)

func SendToToken(app *firebase.App, token string) {
	registrationToken := token
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	message := models.GetNotificationMessage(registrationToken).Message

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}
