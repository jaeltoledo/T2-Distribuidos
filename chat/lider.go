package chat

import (
	"log"
	"math/rand"
	"time"
	"golang.org/x/net/context"
	"fmt"
)

type Server struct {
}

var jugadoresTotales int = 16
var jugadorActual int = 0
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

var jugadores [16] int;
var grupo1 [8] int
var grupo2 [8] int

var arreglarGrupo bool = true;

func numeroAleatorio(valorMin int, valorMax int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return valorMin+rand.Intn(valorMax-valorMin)
}

func (s *Server) Bienvenida(ctx context.Context, msg *MensajeBienvenida) (*MensajeBienvenida, error) {
	peticion := msg.Body
	value := ""
	id := 1
	fmt.Println("Dejar entrar al jugador ID: ", jugadorActual)
	fmt.Scan(&value)
	switch peticion {
		case "1":
			
			value = "Usted ha sido aceptado para participar"
			id = jugadorActual
			jugadores[id] = 1
			jugadorActual += 1
			/*
			Debe verificar cuantos están conectados, 
			si son más de 16 no dejar entrar
			*/
	}
	return &MensajeBienvenida{Body: value, Id: int32(id)}, nil
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
	log.Printf("Numero elegido por el Lider: %d", numeroLider)

	if numeroJugador >= numeroLider {
		//El jugador es eliminado, por lo que se debe actualizar el pozo
		log.Printf("Se ha eliminado un jugador")
		value = -1
		jugadoresTotales = jugadoresTotales -1
		jugadores[int(msg.Id)] = 0
	}
	fmt.Println("Los jugadores que quedan son: ", jugadoresTotales)
	fmt.Println(jugadores)
	return &MensajeEtapa1{Body: int32(value)}, nil
}

func (s *Server) InicioEtapa2(ctx context.Context, msg *MensajeEtapa2) (*MensajeEtapa2, error) {
	Value := int32(0)
	if ((jugadoresTotales%2==0 || jugadoresTotales==1)&& arreglarGrupo) {
		arreglarGrupo = false;
		tocaG1 := true
		tocaG2 := false
		contador1 := 0
		contador2 := 0
		for i := 0; i <= 15; i++ {
			if (jugadores[i] == 1){
				if (tocaG1){
					grupo1[contador1] = i
					contador1++
					tocaG1 = false
					tocaG2 = true
				}
				if (tocaG2){
					grupo2[contador2] = i
					contador2++
					tocaG1 = true
					tocaG2 = false
				}
			}
		}
		fmt.Println("Los equipos han sido elegidos")

	}
	for i := 0; i <= 7; i++ {
		if (msg.Id == int32(grupo1[i])){
			Value = int32(1)
		}else{
			Value = int32(2)
		}
	}
	return &MensajeEtapa2{Body: int32(1), Id: msg.Id, Group: Value}, nil
}