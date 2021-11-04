package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, msg *Message) (*Message, error) {
	log.Printf("Mensaje recibido del cliente %s", msg.Body)
	return &Message{Body: "Hola del servidor"}, nil
}
