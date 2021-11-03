package main

import (
	"fmt"
  	"time"
	"log"
	"os"
	"time"
	"math/Rand"

	"github.com/Tarea1/Express/logistica"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func mostrarMenu() {
	fmt.Println("Bienvenido al Juego del Calamar")
	fmt.Println("Seleccione la acción que desea realizar:")
	fmt.Println("1. Enviar petición para unirse al juego")
	fmt.Println("2. Ver el monto acumulado")
	fmt.Println("3. Jugar Etapa")
  }

func enviarPeticion(){
}

func jugarEtapa1() {
	distanciaElegida := 0
	valorLider := 6 //Debe ir entre 6 y 10
	distanciaRecorrida:= 0
	eliminado := true
	fmt.Println("Bienvenido a la Etapa 1")
	for i:= 1; i<=4; i++ {
		fmt.Printf("Ronda %d - Ingrese un número entre 1 y 10: ", i)
		fmt.Scan(&distanciaElegida)
		/*
		Estaría faltando comunicarse con el lider para 
		obtener su número al azar
		*/

		//Comparación del valor con el del lider
		if (distanciaElegida >= valorLider){
			eliminado = true
			break
		}else{
			distanciaRecorrida += distanciaElegida
			if (distanciaRecorrida >=21) {
				eliminado = true
				break
			}
		}
	}
	if (eliminado) {
		fmt.Println("El jugador ha sido eliminado")
	}else{
		fmt.Println("El jugador ha ganado")
	}
	

}
func jugarEtapa2() {
	fmt.Println("Bienvenido a la Etapa 2")
	valorEntregado := rand.Intn(4)

	fmt.Println(valorEntregado)
}
func jugarEtapa3() {
	fmt.Println("Bienvenido a la Etapa 3")
}

func enviarJugada(){
}

func verPozo(){
}

func main() {
	jugarEtapa2()
	distanciaRecorrida := 0
	fmt.Scan(&distanciaRecorrida)
}