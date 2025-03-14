package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Conn_Rabbit struct {
	Broker *amqp.Connection
	Channel *amqp.Channel
	Err string
}

func GetConnRabbit() *Conn_Rabbit {

	error := ""
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Obtener las variables
	rabbitUrl := os.Getenv("RABBIT_URL")

    // Conexión a RabbitMQ
    conn, err := amqp.Dial(rabbitUrl)
    if err != nil {
		log.Fatal("Error al abrir una conexión hacia rabbitmq")
    }

    // Abrimos un canal
    ch, err := conn.Channel()
    if err != nil {
        log.Fatal("Error al abrir un canal")
    }

    return &Conn_Rabbit{Broker: conn, Channel: ch, Err: error}

}

func (conn *Conn_Rabbit) FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
