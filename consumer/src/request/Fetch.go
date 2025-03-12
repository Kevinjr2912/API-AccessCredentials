package request

import (
	"bytes"
	"log"
	"net/http"
)

func Fetch(jsonPayload []byte) {

	// Ruta
	url := "http://localhost:8081/accessCredentials"

	// Hacemos la petición
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))

	if err != nil {
		log.Println("Error al hacer la petición:", err)
		return
	}

	if resp.StatusCode != http.StatusCreated {
		log.Printf("Error al mandar mensaje, código: %d", resp.StatusCode)
		return
	}

	log.Println("Solicitud procesada")

}