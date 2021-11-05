package chat

import (
	"log"
	"math/rand"
	"strconv"
	"time"
	"golang.org/x/net/context"
	"fmt"
)

type Server struct {
}

var jugadoresTotales int = 1
//Variables que tendrán las jugadas de los jugadores
/*
var jugadasJugador1 string = ""
var jugadasJugador2 string = ""
var jugadasJugador3 string = ""
var jugadasJugador4 string = ""
var jugadasJugador5 string = ""
var jugadasJugador6 string = ""
var jugadasJugador7 string = ""
var jugadasJugador8 string = ""
var jugadasJugador9 string = ""
var jugadasJugador10 string = ""
var jugadasJugador11 string = ""
var jugadasJugador12 string = ""
var jugadasJugador13 string = ""
var jugadasJugador14 string = ""
var jugadasJugador15 string = ""
var jugadasJugador16 string = ""
*/

func numeroAleatorio(valorMin int, valorMax int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return valorMin+rand.Intn(valorMax-valorMin)
}

func (s *Server) Bienvenida(ctx context.Context, msg *MensajeBienvenida) (*MensajeBienvenida, error) {
	peticion := msg.Body
	value := ""
	log.Printf("El lider se ha conectado")
	switch peticion {
		case "1":
			
			value = strconv.Itoa(jugadoresTotales)
			jugadoresTotales += 1
			/*
			Debe verificar cuantos están conectados, 
			si son más de 16 no dejar entrar
			*/
	}
	return &MensajeBienvenida{Body: value}, nil
}

func (s *Server) EntreEtapas(ctx context.Context, msg *MensajeEntreEtapas) (*MensajeEntreEtapas, error) {
	peticion := msg.Body
	value := ""
	switch peticion {
		case "1":
			fmt.Printf("Ingrese 'comenzar' para iniciar la siguiente etapa: ")
			fmt.Scan(&value)
			value = "Se da por iniciada la etapa"
			/*
			El jugador puede pasar a la siguiente etapa
			*/
		case "2":
			log.Printf("Petición recibida: %s", msg.Body)
			value = "Puedes ver el pozo"
			/*
			Comunicación al pozo para recibir información y luego devolverla
			*/
	}
	return &MensajeEntreEtapas{Body: value}, nil
}

func (s *Server) Etapa1(ctx context.Context, msg *MensajeEtapa1) (*MensajeEtapa1, error) {
	numeroJugador := int(msg.Body)
	value := 0
	numeroLider := numeroAleatorio(6, 10)
	log.Printf("Numero elejido por el Lider: %d", numeroLider)

	if numeroJugador >= numeroLider {
		//El jugador es eliminado, por lo que se debe actualizar el pozo
		value = -1
	}
	return &MensajeEtapa1{Body: int32(value)}, nil
}

