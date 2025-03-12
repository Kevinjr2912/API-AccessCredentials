package main

import "consumer/src/rabbit"

func main() {
	// Instanciamos un objeto de tipo Rabbit
	rabbitMQ := rabbit.NewRabbit()

	msg := rabbitMQ.ReceiveContent()
	
	rabbit.ProcessMessage(msg)
}