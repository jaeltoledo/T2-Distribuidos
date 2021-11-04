package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {

}

func (s *Server) SendMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message: %s\n", message.body)
	return &Message{body: "Hello from the Server!"}, nil
}