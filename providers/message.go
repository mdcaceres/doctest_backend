package providers

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/mdcaceres/doctest/config"
	"github.com/mdcaceres/doctest/models"
	"log"
)

type MessagingProvider struct {
	FIRE *firebase.App
	CTX  context.Context
	MSG  *messaging.Client
}

func NewMessagingProvider() MessagingProvider {
	app, ctx, msg := config.SetupFirebase()
	return MessagingProvider{
		FIRE: app,
		CTX:  ctx,
		MSG:  msg,
	}
}

func (m *MessagingProvider) SendToToken(token string) {
	registrationToken := token

	message := models.GetNotificationMessage(registrationToken)

	response, err := m.MSG.Send(m.CTX, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}
