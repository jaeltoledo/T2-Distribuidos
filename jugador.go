package main

import (
	"log"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/CDonosoK/T2-Distribuidos/chat"
)

func entreEtapa() string{
	decision := ""
	fmt.Println("-------------------------------------------------")
	fmt.Println(("¿Qué deseas hacer? [1] Seguir jugando \n [2] Ver el pozo"))
	fmt.Scanln(&decision)
	fmt.Println("-------------------------------------------------")

	return decision
}

/*
func jugarEtapa1(){
	fmt.Println("Bienvenido al primer juego: Luz Roja, Luz Verde")

	//Conexión con el lider
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}
	defer conn.Close()

	c:= chat.NewChatServiceClient(conn)

	numeroElegido := 0
	distanciaRecorrida := 0
	for ronda := 1; ronda <= 4; ronda++ {
		fmt.Println("Ronda: ", ronda, " - ELija un número entre 1 y 10: ")
		fmt.Scanln(&numeroElegido)

		message:= chat.Message{
			Body: numeroElegido,
		}
		response, err := c.etapa1(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error al llamar a etapa-1: %s", err)
		}

		switch response.Body {
			case "1":
				fmt.Println("Felicidades pasaste a la siguiente ronda")
				distanciaRecorrida += numeroElegido
			case "0":
				fmt.Println("Has sido eliminado")
				//Comunicar al lider que el jugador fue eliminado
		} 
		if distanciaRecorrida >= 21 {
			fmt.Println("Felicidades has pasado la etapa")
			break
		}
		
	}
	if distanciaRecorrida < 21 {
		fmt.Println("Has perdido")
		//Comunicar al lider que el jugador fue eliminado
	}
}

func etapa2(){}

func etapa3(){}
*/

func menuPrincipal() chat.Message{
	var opcion string
	//var estaVido := true
	log.Printf("Ingrese su petición: \n[1] Jugar a Squid-Game \n[2] Jugar Etapa \n[3] Ver el Pozo")
	fmt.Scanln(&opcion)

	message := chat.Message{
		Body: "-",
		}
	switch opcion {
		case "1":
			message.Body = "1"
		case "2":
			message.Body = "2"
		case "3":
			message.Body = "3"
	}

	return message

}


func main(){
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}
	defer conn.Close()

	c:= chat.NewChatServiceClient(conn)

	message := menuPrincipal()

	
	response, err := c.Bienvenida(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error al llamar a Bienvenida: %s", err)
	}

	log.Printf("Respuesta del servidor: %s", response.Body)
}