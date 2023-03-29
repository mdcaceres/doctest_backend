package models

import (
	"firebase.google.com/go/messaging"
	"gorm.io/gorm"
)

type NotificationMessage struct {
	gorm.Model
	Message *messaging.Message
}

func GetNotificationMessage(registrationToken string) *NotificationMessage {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification Test",
			Body:  "Hello React!!",
		},
		Token: registrationToken,
	}

	return &NotificationMessage{
		Message: message,
	}
}


