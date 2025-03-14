package repositories

type IRabbit interface {
	SendMessageToBroker(email string)
}