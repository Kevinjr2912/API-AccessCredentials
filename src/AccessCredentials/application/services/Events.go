package services

import "api_accesscredentials/src/AccessCredentials/application/repositories"

type Event struct {
	rabbit repositories.IRabbit
}

func NewEvent(rabbit repositories.IRabbit) *Event {
	return &Event{rabbit: rabbit}
}

func (e *Event) Run(email string) {
	e.rabbit.SendMessageToBroker(email)
}