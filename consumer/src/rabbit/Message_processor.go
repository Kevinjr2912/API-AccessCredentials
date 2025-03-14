package rabbit

import (
	"consumer/src/models"
	"consumer/src/request"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/crypto/bcrypt"
)

// Procesar el mensaje
func ProcessMessage(msgs <-chan amqp.Delivery) {
	forever := make(chan struct{})

	go func() {

		for d := range msgs {
			var s models.StudentCredentials

			err := json.Unmarshal(d.Body, &s)

			if err != nil {
				log.Printf("Error al decodificar el mensaje: %s", err)
				continue
			}

			log.Printf("[x] Estudiante recibido")

			var accessCredentials models.AccessCredentials

			passWithouthash := fmt.Sprintf("%d%s", s.Student.Id, s.Student.Name)

			// Hasheamos la contraseña
			hash, err := bcrypt.GenerateFromPassword([]byte(passWithouthash), bcrypt.DefaultCost)
			if err != nil {
				log.Println("Error al hashear la contraseña")
				continue
			}

			accessCredentials.IdAccessCredentials = 1
			accessCredentials.IdStudent = s.Student.Id
			accessCredentials.Email = s.Email
			accessCredentials.Password = string(hash)

			jsonPayload, err := json.Marshal(accessCredentials)

			if err != nil {
				log.Println("Error al serializar JSON")
				continue
			}

			request.Fetch(jsonPayload)
		}

	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}