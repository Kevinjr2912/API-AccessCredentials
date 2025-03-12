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
			var student models.StudentAttributesCreated

			err := json.Unmarshal(d.Body, &student)

			if err != nil {
				log.Printf("Error al decodificar el mensaje: %s", err)
				continue
			}

			log.Printf("[x] Estudiante recibido")

			var accessCredentials models.AccessCredentials

			user := fmt.Sprintf("%s-UP-%d", student.Name, student.Id)
			passWithouthash := fmt.Sprintf("%d%s", student.Id, student.Name)

			// Hasheamos la contraseña
			hash, err := bcrypt.GenerateFromPassword([]byte(passWithouthash), bcrypt.DefaultCost)
			if err != nil {
				log.Println("Error al hashear la contraseña")
				continue
			}

			accessCredentials.IdAccessCredentials = 1
			accessCredentials.IdStudent = student.Id
			accessCredentials.User = user
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