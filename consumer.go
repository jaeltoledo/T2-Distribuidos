//pozo recibe informacion de los jugadores
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

var pozo int = 0

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func actualizarPozo(valor int) {
	var aux int
	aux = pozo
	pozo = aux + valor
}

func main() {

	//crear archivo
	file, err := os.Create("prueba.txt")
	if err != nil {
		log.Fatal("error", err)
	}
	defer file.Close()
	//fmt.Fprintf(file, "Escribir 1")
	// ----------------------------------------------------

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		var agregar int

		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			vbody := d.Body
			arreglo := strings.Split(string(vbody), "-")
			agregar = 100000000
			actualizarPozo(agregar)
			fmt.Fprintf(file, "Jugador_"+arreglo[0]+" Ronda_"+arreglo[1]+" "+strconv.Itoa(pozo))
			fmt.Printf("AAAAAAAAAAAAAA")
		}
	}()
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exit
		log.Printf("Pozo %d", pozo)

		os.Exit(1)
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
