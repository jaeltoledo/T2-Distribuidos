package main

import (
	"fmt"
	"math/rand"
)

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