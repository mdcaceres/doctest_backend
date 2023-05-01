package services

import "github.com/mdcaceres/doctest/providers"

type Messaging struct {
	MessagingProvider providers.MessagingProvider
}

func NewMessaging() Messaging {
	return Messaging{
		MessagingProvider: providers.NewMessagingProvider(),
	}
}

func (m *Messaging) SendToToken(token string) {
	m.MessagingProvider.SendToToken(token)
}
