package main

import (
	"fmt"
	"log"

	"github.com/CDonosoK/T2-Distribuidos/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func peticionUnion() chat.MensajeBienvenida{
	var opcion string
	log.Printf("Desea unirse al juego? \n[1] Si \n[2] No")
	fmt.Scanln(&opcion)

	message := chat.MensajeBienvenida{
		Body: opcion,
	}

	return message
}

func menuEntreEtapas() chat.MensajeEntreEtapas{
	var opcion string
	log.Printf("Ingrese su petición: \n[1] Jugar siguiente etapa \n[2] Ver el pozo")
	fmt.Scanln(&opcion)

	message := chat.MensajeEntreEtapas{
		Body: "-",
		}
	switch opcion {
		case "1":
			message.Body = "1"
		case "2":
			message.Body = "2"
	}

	return message

}


func main(){
	estaVivo := true
	idJugador := int32(0)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %s", err)
	}
	defer conn.Close()

	c:= chat.NewChatServiceClient(conn)

	fmt.Println("------------------------------------")
	fmt.Println("-     Bienvenido al Squid-Game     -")
	fmt.Println("------------------------------------")
	fmt.Println("")

	//Petición de unirse al juego
	message1 := peticionUnion()
	response1, err1 := c.Bienvenida(context.Background(), &message1)
	if err1 != nil {
		log.Fatalf("Error al llamar a Bienvenida: %s", err1)
	}
	log.Printf("%s", response1.Body)
	idJugador = response1.Id
	
	//Petición primera etapa o ver pozo
	message2 := menuEntreEtapas()
	response2, err2 := c.EntreEtapas(context.Background(), &message2)
	if err2 != nil {
		log.Fatalf("Error al llamar a EntreEtapas: %s", err2)
	}
	log.Printf("Respuesta del servidor: %s", response2.Body)

	//Jugar primera etapa
	fmt.Println("------------------------------------------")
	fmt.Println("Primera Etapa - Juego: Luz verde, Luz roja")
	fmt.Println("------------------------------------------")
	fmt.Println("")
	distanciaRecorrida := 0
	numeroElegido := 0
	for ronda := 1; ronda <= 4; ronda++ {
		fmt.Println("Ronda: ", ronda)
		fmt.Println("Elija un número entre 1 y 10")
		fmt.Scan(&numeroElegido)
		message3 := chat.MensajeEtapa1{
			Body: int32(numeroElegido),
			Id: idJugador,
		}
		response3, err3 := c.Etapa1(context.Background(), &message3)
		if err3 != nil {
			log.Fatalf("Error al llamar a Etapa1: %s", err3)
		}

		switch response3.Body {
			case int32(0):
				distanciaRecorrida += numeroElegido
				fmt.Println("Distancia recorrida: ", distanciaRecorrida)
			case int32(-1):
				fmt.Println("El jugador ha sido eliminado")
				estaVivo = false
		}
		if estaVivo == false {
			break
		}
		if distanciaRecorrida >= 21 {
			estaVivo = true
			fmt.Println("----------------------------------------------")
			fmt.Println("- Felicidades, has superado la priemra etapa -")
			fmt.Println("----------------------------------------------")
			fmt.Println("")
		}

	}
	if distanciaRecorrida < 21 && estaVivo == true {
		estaVivo = false
		fmt.Println("El jugador ha sido eliminado")
		message4 := chat.MensajeEtapa1{
			Body: int32(12),
			Id: idJugador,
		}
		response4, err4 := c.Etapa1(context.Background(), &message4)
		if err4 != nil {
			log.Fatalf("Error al llamar a Etapa1: %s", err4)
		}
		log.Printf("Respuesta del servidor: %s", response4.Body)
	}
	if estaVivo == true {
		//Petición segunda etapa o ver pozo

		message5 := menuEntreEtapas()
		response5, err5 := c.EntreEtapas(context.Background(), &message5)
		if err5 != nil {
			log.Fatalf("Error al llamar a EntreEtapas: %s", err5)
		}
		log.Printf("Respuesta del servidor: %s", response5.Body)
	}

	if estaVivo{
		//Iniciar la segunda etapa
		fmt.Println("--------------------------------------")
		fmt.Println("Segunda Etapa - Juego: Tirar la cuerda")
		fmt.Println("--------------------------------------")
		fmt.Println("")


	}

	if estaVivo{
		//Petición tercera etapa o ver pozo
		message6 := chat.MensajeEtapa2{
			Body: int32(1),
			Id: idJugador,

		}
		response6, err6 := c.InicioEtapa2(context.Background(), &message6)
		if err6 != nil {
			log.Fatalf("Error al llamar a EntreEtapas: %s", err6)
		}
		fmt.Println("EL grupo asignado es: ", response6.Group)

	}

	if estaVivo{
		//Iniciar la tercera etapa

	}

	if estaVivo{
		//Obtener las ganancias

	}
	
	
}