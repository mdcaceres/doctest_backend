package models

import (
	"firebase.google.com/go/messaging"
)

func GetNotificationMessage(registrationToken string) *messaging.Message {
	return &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification Test",
			Body:  "msg.Body",
		},
		Token: registrationToken,
	}
}
