package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/CDonosoK/T2-Distribuidos/chat"
)

func main(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}
	defer conn.Close()

	c:= chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "Hola desde el Cliente!",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error al llamar a SayHello: %s", err)
	}

	log.Printf("Respuesta del servidor: %s", response.Body)
}