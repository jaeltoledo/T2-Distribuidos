package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Bienvenida(ctx context.Context, msg *Message) (*Message, error) {
	peticion := msg.Body
	value := ""
	switch peticion {
		case "1":
			log.Printf("Petición recibida: %s", msg.Body)
			value = "No puedes jugar al Squid-Game"
		case "2":
			log.Printf("Petición recibida: %s", msg.Body)
			value = "No puedes la etapa"
		case "3":
			value = "No puedes ver el pozo"
			log.Printf("Petición recibida: %s", msg.Body)
	}
	return &Message{Body: value}, nil

	
}

